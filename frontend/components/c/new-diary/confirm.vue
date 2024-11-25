<script>
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'NewDiaryConfirm',
  props: {
    itinerary: {
      type: Array,
      required: true
    }
  },
  emits: ['confirm'],
  data() {
    return {
      expressTickets: [
        {
          date: '2024/11/09',
          time: '11:00~',
          trainName: 'しおさい３号',
          price: 1500
        },
        {
          date: '2024/11/10',
          time: '11:00~',
          trainName: 'しおさい14号',
          price: 1500
        }
      ]
    }
  },
  computed: {
    totalAmount() {
      const expressTicketsTotal = this.expressTickets.reduce((sum, ticket) => sum + ticket.price, 0)
      const experienceTotal = this.itinerary
        .filter(item => item.kind === 'spot')
        .reduce((sum, item) => sum + item.price, 0)
      return expressTicketsTotal + experienceTotal
    }
  },
  methods: {
    confirmBooking() {
      this.$emit('confirm')
    }
  }
})
</script>

<template>
  <v-card elevation="0">
    <v-card-title>以下で予約しますか</v-card-title>
    <v-card-text>
      <h3>特急券</h3>
      <v-list>
        <v-list-item v-for="(ticket, index) in expressTickets" :key="index">
          <v-list-item-content>
            <v-list-item-title>{{ ticket.date }} {{ ticket.time }} {{ ticket.trainName }}</v-list-item-title>
            <v-list-item-subtitle>{{ ticket.price }}円</v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
      </v-list>

      <h3>体験</h3>
      <v-list>
        <v-list-item v-for="(item, index) in itinerary.filter(i => i.kind === 'spot')" :key="index">
          <v-list-item-content>
            <v-list-item-title>{{ item.title }}</v-list-item-title>
            <v-list-item-subtitle>{{ item.price }}円</v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
      </v-list>
      <v-divider />
      <h3 class="mt-8">合計</h3>
      <v-list>
        <v-list-item>
          <v-list-item-content>
            <v-list-item-title>{{ totalAmount }}円</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-card-text>
  </v-card>
</template>

<style scoped>
/* スタイルが必要な場合はここに追加 */
</style>
