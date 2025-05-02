<template>
  <div>
    <PageLoader v-if="loading" />
    <Operation
      v-else-if="currentOperation"
      :operation="currentOperation"
    />
    <div v-else>
      <h1>Operation not found</h1>
      <v-btn
        color="primary"
        variant="tonal"
        prepend-icon="mdi-home"
        @click="$router.push('/')"
      >
        Go to Home
      </v-btn>
    </div>
  </div>
</template>

<script setup>
import { useOperationStore } from "@/stores/operation";
import { onMounted, ref } from "vue";
import { storeToRefs } from "pinia";
import { useRouter } from "vue-router";

const router = useRouter();
const operationStore = useOperationStore();

const { currentOperation } = storeToRefs(operationStore);
const loading = ref(true);

onMounted(() => {
  var operationId = router.currentRoute.value.params.id;
  let ok = operationStore.setOperation(operationId);
  if (!ok) {
    router.push("/");
    return;
  }

  setTimeout(() => {
    loading.value = false;
  }, 1000);
});
</script>
