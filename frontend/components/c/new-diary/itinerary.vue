<script>
export default {
  name: 'NewDiaryItinerary',
  props: {
    bring: {
      type: Array,
      required: true
    },
    itinerary: {
      type: Array,
      required: true
    }
  },
  methods: {}
}
</script>

<template>
  <v-timeline align-top dense>
    <v-timeline-item
      color="secondary"
      small
    >
      <v-card style="max-width: 250px">
        <v-card-title class="text-h6">
          持ち物
        </v-card-title>
        <v-card-text>
          <ul>
            <li v-for="(item, index) in bring" :key="index">
              {{ item }}
            </li>
          </ul>
        </v-card-text>
      </v-card>
    </v-timeline-item>
    <v-timeline-item
      v-for="item in itinerary"
      :key="item.id"
      :color="item.kind === 'move' ? 'grey' : 'primary'"
      small
    >
      <v-card  style="max-width: 250px">
        <v-card-title class="text-h6">
          {{ item.title }}
        </v-card-title>
        <v-card-subtitle>
          {{ item.description }}
        </v-card-subtitle>
        <v-card-text>
          <div class="d-flex justify-space-between align-center">
            <div>所要時間: {{ item.takeTime }}分</div>
            <div v-if="item.kind !== 'move'">{{ item.price }}円</div>
          </div>
        </v-card-text>
      </v-card>
    </v-timeline-item>
  </v-timeline>
</template>

<style scoped>
</style>
