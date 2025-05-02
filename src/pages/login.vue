<template>
  <PageLoader v-if="authenticating" />
  <div
    v-else
    class="login-container bg-background"
  >
    <v-card>
      <v-btn
        size="x-large"
        :href="loginUrl"
      >
        Login with Discord
      </v-btn>
    </v-card>
  </div>
</template>

<script setup>
import { onMounted, ref } from "vue";

import { authenticated as checkAuthentication, authenticate } from "@/api";

const authenticating = ref(false);
const loginUrl = ref(`${import.meta.env.VITE_DISCORD_AUTH_URL}`);

onMounted(async () => {
  if (localStorage.getItem("authenticated") === "true") {
    try {
      authenticating.value = true;
      let authed = await checkAuthentication();
      if (authed) {
        authenticating.value = false;
        // redirect to home page
        window.location.href = "/";
        return;
      }

      authenticating.value = false;
    } catch (err) {
      console.error("Error checking authentication:", err);
      authenticating.value = false;
    }
  }

  // if the url has the code in the params
  const urlParams = new URLSearchParams(window.location.search);
  const code = urlParams.get("code");
  if (code) {
    authenticating.value = true;
    try {
      let authenticated = await authenticate(code);

      if (
        authenticated.token !== null &&
        authenticated.token !== undefined &&
        authenticated.token !== ""
      ) {
        authenticating.value = false;
        window.location.href = "/";
        localStorage.setItem("authenticated", "true");
        localStorage.setItem("token", authenticated.token);
      }
    } catch (err) {
      authenticating.value = false;
      console.error("Error authenticating:", err);
    }

    // remove the code from the url
    window.history.replaceState({}, document.title, window.location.pathname);
  }
});
</script>
<route lang="yaml">
meta:
  layout: login
</route>
<style lang="scss" scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}
</style>
