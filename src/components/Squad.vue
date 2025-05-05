<template>
  <v-col
    cols="3"
    lg="3"
    md="4"
    sm="6"
    xs="12"
  >
    <v-card
      class="squad"
      :title="squad.name"
    >
      <v-container fluid>
        <v-row no-gutters>
          <v-col cols="12">
            <Member
              v-for="member in squad.members"
              :key="member.id"
              :member="member"
              @click="onClick"
            />
            <v-card
              class="position ma-2"
              text="Add Member"
            >
              <div
                class="position drop-zone"
                :data-squad-id="squad.id"
                @dragover="onDragOver"
                @dragleave="onDragLeave"
                @drop="onDrop"
              />
            </v-card>
          </v-col>
        </v-row>
      </v-container>
      <v-card-actions>
        <v-btn
          variant="tonal"
          color="error"
          size="x-small"
          @click="remove"
        >
          remove
        </v-btn>
        <v-btn
          variant="tonal"
          size="x-small"
        >
          rename
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-col>
  <v-dialog
    v-model="renameDialog"
    width="500"
  >
    <v-card>
      <v-card-title>
        Rename Squad
      </v-card-title>
      <v-card-text>
        <v-text-field v-model="renameName" />
      </v-card-text>
      <v-card-actions>
        <v-spacer />
        <v-btn
          variant="text"
          @click="renameDialog = false"
        >
          Cancel
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script setup>
import { ref } from "vue";

import { useOperationStore } from "@/stores/operation";

const operationStore = useOperationStore();

const props = defineProps({
  squad: {
    type: Object,
    required: true,
    default: () => ({}),
  },
});

const renameDialog = ref(false);
const renameName = ref("");

function onDragOver(event) {
  event.preventDefault();
  event.target.style.backgroundColor = "#00000044";
}

function onDragLeave(event) {
  event.preventDefault();
  event.target.style.backgroundColor = "";
}

function onDrop(event) {
  event.preventDefault();
  event.target.style.backgroundColor = "";

  if (event.target.classList.contains("drop-zone")) {
    const memberId = event.dataTransfer.getData("text");
    operationStore.addMemberToSquad(event.target.dataset.squadId, memberId);
  }
}

function onClick(event) {
  event.preventDefault();
  const memberId = event.currentTarget.id;
  operationStore.removeMember(memberId);
}

function remove() {
  operationStore.removeSquad(props.squad.id);
}
</script>

<style lang="scss" scoped>
.squad {
  // darken background
  background-color: rgb(var(--v-theme-surface-darken-2));
}
</style>
