<template>
  <v-col cols="3">
    <v-card
      class="ship"
      :title="ship.name"
      :v-style="'background-image: url(\'' + ship.image + '\'); background-image-'"
    >
      <v-container fluid>
        <v-row no-gutters>
          <v-col
            v-for="position in ship.positions"
            :key="position.id"
            cols="12"
          >
            <!-- <v-card v-if="position.member !== null" class="position" @click="onMemberClick(ship.id, position.id, position.member)" :text="position.member.name"></v-card> -->
            <Member
              v-if="position.member !== null"
              :member="position.member"
              @click="onClick"
            />
            <v-card
              v-else
              :data-ship-id="ship.id"
              :data-id="position.id"
              class="position ma-2"
              :text="position.name"
            >
              <div
                class="position drop-zone"
                :data-ship-id="ship.id"
                :data-id="position.id"
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
      </v-card-actions>
    </v-card>
  </v-col>
</template>

<script setup>
import { ref } from "vue";

import { useMembersStore } from "@/stores/members";
import { useOperationStore } from "@/stores/operation";

const membersStore = useMembersStore();
const operationStore = useOperationStore();

const props = defineProps({
  ship: {
    type: Object,
    required: true,
    default: () => ({}),
  },
});

const shipRef = ref(props.ship);

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

  if (event.target.classList.contains("position")) {
    // replace with member
    const shipId = event.target.dataset.shipId;
    const positionId = parseInt(event.target.dataset.id);
    const memberId = event.dataTransfer.getData("text");

    console.log("shipId", shipId);

    operationStore.addToShipPosition(shipId, positionId, memberId);

    // membersStore.members.forEach((member) => {
    //   if (member.id == memberId) {
    //     const positionId = parseInt(
    //       getParent(event.target, "position").dataset.id
    //     );
    //     shipRef.value.positions[positionId].member = member;
    //     member.assigned = true;
    //     member.shipId = props.ship.id;

    //     shipRef.value.positions.forEach((position) => {
    //       if (
    //         position.member &&
    //         position.member.id == member.id &&
    //         position.id != positionId
    //       ) {
    //         position.member = null;
    //       }
    //     });

    //     return;
    //   }
    // });
  }
}

function getParent(ele, cls) {
  let e = ele;
  do {
    e = e.parentElement;
  } while (e && !e.classList.contains(cls));
  return e;
}

function onClick(event) {
  event.preventDefault();
  const memberId = event.currentTarget.id;
  operationStore.putOnStandby(memberId);
}

function remove() {
  operationStore.removeShip(props.ship.id);
}
</script>

<style lang="scss" scoped>
.ship {
  // darken background
  background-color: rgb(var(--v-theme-surface-darken-2));
}
</style>
