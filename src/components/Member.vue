<template>
  <v-card
    v-if="member !== undefined"
    :id="member.id"
    :class="'member my-2 mr-2 ' + membersStore.getRankName(member.rank)"
    :draggable="true"
    @dragstart="onDragStart"
    @dragover="onDragEnd"
    @click="onClick"
  >
    <v-card :text="membersStore.getNameWithRank(member)" />
  </v-card>
</template>

<script setup>
import { useMembersStore } from "@/stores/members";

const membersStore = useMembersStore();

defineProps({
  onClick: {
    type: Function,
    default: () => {},
  },
  member: Object,
});

function onDragStart(event) {
  event.dataTransfer.dropEffect = "move";
  event.dataTransfer.effectAllowed = "move";
  event.dataTransfer.setData("text/plain", event.target.id);
}

function onDragEnd(event) {
  event.preventDefault();
}
</script>

<style lang="scss" scoped>
.member {
  border-left: 3px solid rgb(var(--v-theme-primary));

  &.Admiral {
    border-color: rgb(var(--v-theme-admiral));
  }

  &.Commander {
    border-color: rgb(var(--v-theme-commander));
  }

  &.Lieutenant {
    border-color: rgb(var(--v-theme-lieutenant));
  }

  &.Specialist {
    border-color: rgb(var(--v-theme-specialist));
  }

  &.Technician {
    border-color: rgb(var(--v-theme-technician));
  }

  &.Member {
    border-color: rgb(var(--v-theme-member));
  }

  &.Recruit {
    border-color: rgb(var(--v-theme-recruit));
  }

  &.Guest {
    border-color: rgb(var(--v-theme-guest));
  }

  &.Ally {
    border-color: rgb(var(--v-theme-ally));
  }

  :hover {
    cursor: grab;
  }

  :active {
    cursor: grabbing;
  }
}
</style>
