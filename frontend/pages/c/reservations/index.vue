<template>
  <v-container>
    <v-row justify="center" align="center">
      <v-col cols="12">
        <h2 class="text-center mb-4">予約確認</h2>
        <v-card v-if="reservations.length === 0" flat>
          <v-card-text class="text-center">予約がありません</v-card-text>
        </v-card>
        <v-card
          v-for="reservation in reservations"
          v-else
          :key="reservation.id"
          class="mb-4"
          flat
        >
          <v-row no-gutters>
            <v-col cols="4">
              <v-img
                :src="reservation.image"
                max-height="128px"
                max-width="158px"
                cover
              />
            </v-col>
            <v-col cols="8" class="px-2">
              <v-container>
                <v-row>
                  <strong>{{ reservation.title }}</strong>
                </v-row>
                <v-row>
                  {{ formatDate(reservation.fromDate) }} - {{ reservation.toDate }}
                </v-row>
                <v-row>
                  <div>{{ reservation.city }}</div>
                  <v-spacer />
                  <v-btn
                    class="mt-4"
                    color="primary"
                    outlined
                    @click="goToItinerary(reservation.id)"
                  >
                    旅程確認・日記更新
                  </v-btn>
                </v-row>
              </v-container>
            </v-col>
          </v-row>
        </v-card>
      </v-col>
    </v-row>
    <v-snackbar
      v-model="snackbar.show"
      :timeout="snackbar.timeout"
      :color="snackbar.color"
      dense
      :close-text="snackbar.closeText"
      :close-icon="snackbar.closeIcon"
    >
      {{ snackbar.text }}
      <template #action="{ attrs }">
        <v-btn
          text
          v-bind="attrs"
          @click="snackbar.show = false"
        >
          <v-icon>{{ snackbar.closeIcon }}</v-icon>
        </v-btn>
      </template>
    </v-snackbar>
  </v-container>
</template>

<script>
import { useUserStore } from '~/store/user'

export default {
  name: 'IndexPage',
  layout: 'mobile',
  data() {
    return {
      userInfo: {},
      reservations: [],
      snackbar: {
        show: false,
        text: '',
        timeout: 3000,
        color: 'success',
        closeText: '閉じる',
        closeIcon: 'mdi-close'
      }
    }
  },
  async mounted() {
    if (this.$route.query.reservation === 'success') {
      this.showSnackbar('予約が完了しました！')
    }
    const userStore = useUserStore()
    userStore.initializeUser()
    this.userInfo = userStore.userInfo
    await this.fetchReservations()
  },
  methods: {
    formatDate(date) {
      return new Date(date).toLocaleDateString('ja-JP')
    },
    goToItinerary(id) {
      this.$router.push(`/c/reservations/${id}/itinerary`)
    },
    async fetchReservations() {
      const params = new URLSearchParams({
        user_id: this.userInfo.id,
        filter: 'yet'
      })
      try {
        const response = await fetch(`${process.env.BASE_URL}/reservations?${params.toString()}`, {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json'
          }
        })
        const data = await response.json()

        // APIレスポンスを変換
        this.reservations = data.reservations.map(item => ({
          id: item.id,
          title: item.travel_spot_title,
          fromDate: item.from_date.split('T')[0],
          toDate: item.to_date.split('T')[0],
          city: '銚子',
          image: `${process.env.BASE_URL}/images/${item.diary_photo_path}`,
          isOffer: item.is_offer
        }))
      } catch (error) {
        console.error('予約の取得に失敗しました', error)
      }
    },
    showSnackbar(text) {
      this.snackbar.text = text
      this.snackbar.show = true
    }
  }
}
</script>
