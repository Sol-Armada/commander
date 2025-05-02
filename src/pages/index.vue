<template>
  <v-container class="bg-background">
    <v-row class="d-flex justify-center align-center">
      <v-col
        cols="12"
        md="8"
        lg="6"
      >
        <v-data-table-virtual
          :headers="headers"
          :items="vitrualOperations"
          :items-per-page="itemsPerPage"
          fixed-header=""
        >
          <template #top>
            <v-toolbar flat>
              <v-toolbar-title>Operations</v-toolbar-title>
              <v-spacer />
              <v-btn
                color="primary"
                variant="tonal"
                prepend-icon="mdi-plus"
                @click="newOperation"
              >
                New Operation
              </v-btn>
            </v-toolbar>
          </template>

          <template #item="{ item }">
            <tr @click="selectOperation(item.id)">
              <td>{{ item.name }}</td>
              <td>{{ item.createdDate.toLocaleDateString() }}</td>
            </tr>
          </template>
          <template #bottom>
            <v-pagination
              v-model="page"
              :length="Math.ceil(operations.length / itemsPerPage)"
              :total-visible="7"
            />
          </template>
          <template #no-data>
            No operations found.
          </template>
        </v-data-table-virtual>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import { useRouter } from "vue-router";
import { createOperation } from "../api";
import { useOperationStore } from "@/stores/operation";
import { storeToRefs } from "pinia";
import { computed, ref } from "vue";

const operationStore = useOperationStore();
const { operations } = storeToRefs(operationStore);
const vitrualOperations = computed(() => {
  return operations.value.map((operation) => ({
    ...operation,
    createdDate: new Date(Date.parse(operation.created)),
  }));
});

const itemsPerPage = 10;
const page = ref(1);

const headers = [
  { text: "Name", value: "name" },
  { text: "Status", value: "status" },
];

const router = useRouter();

async function newOperation() {
  let op = await createOperation();
  operationStore.addOperation(op);
  router.push({ path: `operations/${op.id}` });
}

function selectOperation(operationId) {
  operationStore.setOperation(operationId);
  router.push({ path: `operations/${operationId}` });
}
</script>

<style lang="scss" scoped>
.v-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}

table > tbody > tr {
  cursor: pointer;
}
table > tbody > tr:hover {
  background-color: rgba(0, 0, 0, 0.1);
}
</style>
