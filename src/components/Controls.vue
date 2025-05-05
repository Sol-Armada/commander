<template>
  <v-row class="py-2">
    <v-col cols="2">
      <v-autocomplete
        v-model="selectedShip"
        hide-details
        density="comfortable"
        variant="solo-filled"
        item-props
        :label="loading ? 'Loading...' : 'Add Ship'"
        :disabled="loading"
        :items="shipNames"
      />
    </v-col>
    <v-col
      cols="2"
      align-self="center"
    >
      <v-btn @click="addSquad">
        Add Squad
      </v-btn>
    </v-col>
  </v-row>
</template>

<script setup>
import { useShipsStore } from "@/stores/ships";
import { useOperationStore } from "@/stores/operation";
import { ref, computed, watch } from "vue";
import { storeToRefs } from "pinia";

const shipsStore = useShipsStore();
const operationStore = useOperationStore();

const { loading, ships } = storeToRefs(shipsStore);

const shipNames = computed(() => {
  const s = [];
  ships.value.forEach((ship) => {
    let found = s.find((s) => s.title === ship.name || s.value === ship.id);
    if (!found) {
      s.push({ title: ship.name, value: ship.id });
    }
  });
  return s;
});

const selectedShip = ref(null);
watch(selectedShip, (val) => {
  if (val == null) {
    return;
  }
  let ship = shipsStore.getShip(val);
  operationStore.addShip(ship);
  selectedShip.value = null;
});

function addSquad() {
  operationStore.addSquad();
}
</script>
