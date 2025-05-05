<template>
  <v-container fluid>
    <v-row>
      <v-col cols="2">
        <v-row class="sticky">
          <v-col cols="12">
            <v-text-field
              v-if="operationStore.isCreator()"
              v-model="opName"
              placeholder="Rename operation"
              variant="plain"
            />
            <div v-else>
              {{ operation.name }}
            </div>
          </v-col>
          <Members
            :members="filteredMembers"
            :loading="loading"
          />
        </v-row>
      </v-col>

      <v-col cols="10">
        <v-row class="sticky">
          <v-col>
            <StandBy />
            <Controls />
          </v-col>
        </v-row>
        <v-row>
          <Board />
        </v-row>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import { computed, ref, watch } from "vue";
import { useMembersStore } from "@/stores/members";
import { storeToRefs } from "pinia";
import { useOperationStore } from "@/stores/operation";

const operationStore = useOperationStore();
const membersStore = useMembersStore();

const { currentOperation } = storeToRefs(operationStore);

const { loading } = storeToRefs(membersStore);
const opName = ref(currentOperation.value.name);
watch(opName, (newName) => {
  operationStore.renameOperation(newName);
});

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
.sticky {
  position: sticky;
  top: 0;
  z-index: 1;

  background-color: rgb(var(--v-theme-background));
}
</style>
