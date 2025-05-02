<template>
  <v-row
    class="stand-by-container elevation-2 rounded"
    no-gutters
  >
    <v-col>
      <div class="pa-2 rounded">
        <v-row>
          <v-col>
            <h2>Stand By</h2>
            <v-divider />
          </v-col>
        </v-row>
        <v-row
          no-gutters
          class="stand-by-list-container"
        >
          <div
            class="stand-by drop-zone"
            @dragover="onDragOver"
            @dragleave="onDragLeave"
            @drop="onDrop"
            @click="onClick"
          />
          <v-col
            v-for="member in currentOperation.onStandBy"
            :key="member.id"
            cols="2"
            class="stand-by-list"
          >
            <Member
              :member="member"
              :on-click="onClick"
            />
          </v-col>
        </v-row>
      </div>
    </v-col>
  </v-row>
</template>

<script setup>
import { useOperationStore } from "@/stores/operation";
import { storeToRefs } from "pinia";

const operationStore = useOperationStore();
const { currentOperation } = storeToRefs(operationStore);

function onDragLeave(event) {
  event.preventDefault();
  event.target.style.backgroundColor = "";
}

function onDragOver(event) {
  event.preventDefault();
  event.target.style.backgroundColor = "#00000044";
}

function onClick(event) {
  event.preventDefault();
  const memberId = event.currentTarget.id;
  operationStore.removeMember(memberId);
}

function onDrop(event) {
  event.preventDefault();
  event.target.style.backgroundColor = "";

  // check if the target has the class 'stand-by-list'
  if (event.target.classList.contains("stand-by")) {
    const memberId = event.dataTransfer.getData("text");
    operationStore.putOnStandby(memberId);
  }
}
</script>

<style lang="scss" scoped>
.stand-by-container {
  min-height: 7rem;

  position: sticky;
  top: 0;
  z-index: 1;

  background: rgb(var(--v-theme-surface));
  color: rgb(var(--v-theme-on-surface));
}

.stand-by.drop-zone {
  min-height: 65px;
}
</style>
