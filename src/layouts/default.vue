<template>
  <v-main>
    <router-view />
  </v-main>
</template>

<script setup>
import { useAppStore } from "@/stores/app";
import { useWebsocketStore } from "@/stores/websocket";
import { onMounted } from "vue";

const appStore = useAppStore();
const wsStore = useWebsocketStore();

onMounted(() => {
  wsStore.connect();
  appStore.getSelf();
});
</script>

<style lang="scss">
html,
body {
  background: rgb(var(--v-theme-background));

  ::-webkit-scrollbar {
    display: none;
  }
}

#app {
  padding: 0;
}

.drop-zone {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}
</style>
