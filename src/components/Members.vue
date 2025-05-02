<template>
  <v-col
    cols="2"
    class="h-screen"
  >
    <div class="members-list-container text-center">
      <div class="search-container">
        <v-text-field
          id="search"
          v-model="searchVal"
          hide-details
          name="search"
          label="Search"
          variant="solo-filled"
          class="elevation-3 mr-2 pa-0"
        />
        <v-progress-linear
          v-if="loading"
          indeterminate
          color="cyan"
        />
      </div>
      <!-- <v-virtual-scroll class="members-list" :items="filteredMembers" height="90dvh">
                <template v-slot:default="{ item }">
                    <Member :member="item" :onDragStart="onDragStart" :onDragEnd="onDragEnd" />
                </template>
            </v-virtual-scroll> -->
      <Member
        v-for="member in filteredMembers"
        :key="member.id"
        :member="member"
        @click="onClick"
      />
    </div>
  </v-col>
</template>

<script setup>
import { ref, computed } from "vue";

const props = defineProps({
  onClick: {
    type: Function,
    default: () => {},
  },
  loading: {
    type: Boolean,
    default: true,
  },
  members: {
    type: Array,
    required: true,
    default: () => [],
  },
});

const searchVal = ref("");
const filteredMembers = computed(() => {
  return props.members.filter((member) => {
    return member.name.toLowerCase().includes(searchVal.value.toLowerCase());
  });
});
</script>

<style lang="scss" scoped>
.search-container {
  position: sticky;
  top: 0;

  z-index: 1;
}

.members-list-container {
  position: sticky;
  top: 0;

  max-height: 100%;
  overflow-y: scroll;
}
</style>
