<template>
  <v-row>
    <v-col
      cols="12"
      style="flex-grow: 1;"
    >
      <v-sheet class="bg-transparent">
        <v-row dense>
          <Ship
            v-for="ship in currentOperation.ships"
            :key="ship.id"
            :ship="ship"
          />

          <Squad
            v-for="squad in currentOperation.squads"
            :key="squad.id"
            :squad="squad"
          />
        </v-row>
      </v-sheet>
    </v-col>
  </v-row>
</template>

<script setup>
import { useOperationStore } from "@/stores/operation";
import { storeToRefs } from "pinia";

const operationStore = useOperationStore();

const { currentOperation } = storeToRefs(operationStore);
</script>

<style lang="scss" scoped>
</style>
