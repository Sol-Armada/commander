
import { defineStore } from 'pinia';
import { ref } from 'vue';
import {
    WebSocketChannels,
    WebSocketRequest,
    WebSocketResponse,
    WebSocketTypes,
} from "@/classes/ws";

export const useWebsocketStore = defineStore('websocket', () => {
    const socket = ref(null)
    const id = ref(null)
    const channels = ref({})

    async function connect() {
        if (socket.value && socket.value.readyState === WebSocket.OPEN) {
            console.debug("WebSocket already connected");
            return;
        }

        socket.value = new WebSocket(`${import.meta.env.VITE_WS_BASE_URL}`);
        socket.value.onclose = () => {
            console.debug("WebSocket connection closed");

            connected.value = false;
        };

        socket.value.onopen = () => {
            console.debug("WebSocket connection opened");

            let req = new WebSocketRequest(socket.value, "register", WebSocketTypes.REGISTER, {
                token: localStorage.getItem("token"),
            });
            req.send();

            connected.value = true;
        };

        socket.value.onmessage = (res) => {
            res = new WebSocketResponse(res);

            console.debug("WebSocket message received", res);

            if (res.type === WebSocketTypes.REGISTERED) {
                console.debug("WebSocket registered", res);
                id.value = res.data.id;
                return;
            }

            if (res.from === id.value) {
                console.debug("WebSocket message from self, ignoring", res);
                return;
            }

            if (res.channel === "keepalive") {
                console.debug("WebSocket message received for keep alive", res);
                let req = new WebSocketRequest(
                    socket.value,
                    WebSocketChannels.KEEP_ALIVE,
                    WebSocketTypes.PING,
                    {
                        token: localStorage.getItem("token"),
                    }
                );
                req.send();
                return;
            }

            console.debug("WebSocket message received", res);
            channels.value[res.channel]?.(res);
        };

        socket.value.onerror = (error) => {
            console.error("WebSocket error:", error);
        };
    }

    function connected() {
        return socket.value && socket.value.readyState === WebSocket.OPEN;
    }

    function addChannel(channel, callback) {
        if (channels.value[channel]) {
            console.debug("Channel already exists", channel);
            return;
        }
        channels.value[channel] = callback;
    }

    async function send(channel, type, data) {
        if (socket.value === null || socket.value.readyState !== WebSocket.OPEN) {
            await connect();
            // try again
            setTimeout(() => {
                send(channel, type, data);
            }, 250);
            return;
        }

        const req = new WebSocketRequest(socket.value, channel, type, data);
        req.send();
    }

    return {
        socket,
        connected,
        addChannel,
        connect,
        send,
    }
})
