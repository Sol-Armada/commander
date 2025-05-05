<template>
  <v-col
    cols="3"
    lg="3"
    md="4"
    sm="6"
    xs="12"
  >
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
import { useOperationStore } from "@/stores/operation";

const operationStore = useOperationStore();

const props = defineProps({
  ship: {
    type: Object,
    required: true,
    default: () => ({}),
  },
});

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
    const shipId = event.target.dataset.shipId;
    const positionId = parseInt(event.target.dataset.id);
    const memberId = event.dataTransfer.getData("text");

    operationStore.addToShipPosition(shipId, positionId, memberId);
  }
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
  background-color: rgb(var(--v-theme-surface-darken-2));
}
</style>
