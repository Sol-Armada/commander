// enum of websocket types
export const WebSocketTypes = {
    REGISTER: "register",
    REGISTERED: "registered",
    PING: "ping",
    LIST: "list",
    UPDATED: "updated",
    UPDATE: "update",
    DELETED: "deleted",
    DELETE: "delete",
    CREATE: "create",
    CREATED: "created",
}

export const WebSocketChannels = {
    KEEP_ALIVE: "keepalive",
    OPERATIONS: "operations",
    MEMBERS: "members",
}

export class WebSocketResponse {
    constructor(data) {
        let json = JSON.parse(data.data)
        this.data = json.data
        this.type = json.type
        this.channel = json.channel
        this.from = json.from
    }
}

export class WebSocketRequest {
    constructor(socket, channel, type, data) {
        this.socket = socket

        this.channel = channel
        this.type = type
        this.data = data
    }

    send() {
        const request = {
            channel: this.channel,
            type: this.type,
            data: this.data
        }
        if (this.socket && this.socket.readyState === WebSocket.OPEN) {
            this.socket.send(JSON.stringify(request))
        }
    }
}

