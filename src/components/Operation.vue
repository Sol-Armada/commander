<template>
  <v-container
    fluid
    class="full-height-container pa-2"
    style="overflow-y: hidden;"
  >
    <v-row>
      <v-col cols="12">
        <h1 class="text-center">
          {{ operation.name }}
        </h1>
      </v-col>
    </v-row>
    <v-row no-gutters>
      <Members
        :members="filteredMembers"
        :loading="loading"
      />

      <v-col
        cols="10"
        class="operations-container"
      >
        <StandBy />

        <v-window>
          <v-window-item>
            <Board />
          </v-window-item>
        </v-window>
      </v-col>

      <v-dialog
        v-model="renameDialog"
        width="500"
      >
        <v-card>
          <v-card-title>
            Rename Operation
          </v-card-title>
          <v-card-text>
            <v-text-field v-model="newOpName" />
          </v-card-text>
          <v-card-actions>
            <v-spacer />
            <v-btn
              variant="text"
              @click="renameDialog = false"
            >
              Cancel
            </v-btn>
            <v-btn
              variant="text"
              @click="operation.rename(newOpName)"
            >
              Rename
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </v-row>
  </v-container>
</template>

<script setup>
import { ref, computed } from "vue";
import { useMembersStore } from "@/stores/members";
import { storeToRefs } from "pinia";
import { useOperationStore } from "@/stores/operation";

const operationStore = useOperationStore();
const membersStore = useMembersStore();
const renameDialog = ref(false);
const newOpName = ref("");

const { loading } = storeToRefs(membersStore);

defineProps({
  operation: {
    type: Object,
    required: true,
    default: () => ({}),
  },
});

const filteredMembers = computed(() => {
  return membersStore.members.filter((member) => {
    return !operationStore.hasMember(member.id);
  });
});
</script>

<style lang="scss">
.full-height-container {
  height: 100vh; /* Full viewport height */
  display: flex;
  flex-direction: column;
}

.operations-container {
  flex: 1; /* Ensures it takes the remaining space */
  overflow-y: auto; /* Allows scrolling if content overflows */
}
</style>
