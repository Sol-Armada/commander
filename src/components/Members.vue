<template>
  <v-row>
    <v-col cols="12">
      <!-- <div class="members-list-container text-center"> -->
      <div class="search-container">
        <v-text-field
          id="search"
          v-model="searchVal"
          hide-details
          name="search"
          label="Search for a member"
          variant="solo-filled"
          class="elevation-3"
          style="height: auto"
        />
        <v-virtual-scroll
          :items="filteredMembers"
          class="elevation-3 mr-2 pl-2"
          style="background: rgb(var(--v-theme-surface)); max-height: 400px;"
        >
          <template #default="{item}">
            <Member :member="item" />
          </template>
        </v-virtual-scroll>
      </div>

      <!-- <Member
      v-for="member in filteredMembers"
      :key="member.id"
      :member="member"
    /> -->
      <!-- </div> -->
    </v-col>
  </v-row>
</template>

<script setup>
import { ref, computed, onMounted, onBeforeUnmount } from "vue";

const props = defineProps({
  onClick: {
    type: Function,
    default: () => {},
  },
  members: {
    type: Array,
    required: true,
    default: () => [],
  },
});

const searchVal = ref("");

const handleClickOutside = (event) => {
  const virtualScroll = document.querySelector(".v-virtual-scroll");
  if (virtualScroll && !virtualScroll.contains(event.target)) {
    searchVal.value = "";
  }
};

onMounted(() => {
  document.addEventListener("click", handleClickOutside);
});

onBeforeUnmount(() => {
  document.removeEventListener("click", handleClickOutside);
});
const filteredMembers = computed(() => {
  if (searchVal.value == "") {
    return [];
  }
  return props.members.filter((member) => {
    return member.name.toLowerCase().includes(searchVal.value.toLowerCase());
  });
});
</script>

<style lang="scss" scoped>
.member-list {
  background: rgb(var(--v-theme-surface));
}
</style>
