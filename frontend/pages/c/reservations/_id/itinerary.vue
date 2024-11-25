<script>
import NewDiaryItinerary from '~/components/c/new-diary/itinerary.vue'
import ReservationDiary from '~/components/c/diary/ReservationDiary.vue'
import { useUserStore } from '~/store/user'

export default {
  name: 'Itinerary',
  components: { ReservationDiary, NewDiaryItinerary },
  layout: 'mobile',
  data() {
    return {
      diary: {
        id: '',
        date: '',
        image: '',
        title: '',
        content: ''
      },
      diaryContent: '',
      bring: [],
      itinerary: [],
      userInfo: {},
    }
  },
  async mounted() {
    const userStore = useUserStore()
    userStore.initializeUser()
    this.userInfo = userStore.userInfo
    await this.fetchItinerary()
    // API通信の代わりにダミーデータを返す
    // this.bring = [
    //   'ダミー軍手', 'ダミー手袋', 'ダミージャケット', 'ダミーハンカチ'
    // ]

    // this.itinerary = [
    //   { type: '移動', duration: 30, description: '目安: 30分' },
    //   { type: '観光', duration: 30, description: 'しあわせ三蔵記念撮影(30分)' },
    //   { type: '移動', duration: 45, description: '目安：45分' },
    //   { type: 'アクティビティ', duration: 180, description: '14:00~ 一本釣り体験( 3時間 )' },
    //   { type: '移動', duration: 120, description: '目安: 2時間' }
    // ]
    // this.diary = {
    //   id: 1,
    //   date: '2024/11/23',
    //   image: 'http://localhost:3000/auster-mono/destination/choshi.jpg',
    //   title: 'ダミーデータ！！！',
    //   content: '今日は早朝から漁船に乗り、期待に胸を膨らませて出航しました。風は少し冷たかったけれど、海の静けさが心地よかったです。そして、ついに大物のヒラマサがヒット！かなりの引きで、腕がパンパンになりましたが、無事に釣り上げることができました。この魚の力強さと美しさには感動しました。次回もこのサイズを狙いたいと思います！'
    // }
    // this.diaryContent = '今日は早朝から漁船に乗り、期待に胸を膨らませて出航しました。風は少し冷たかったけれど、海の静けさが心地よかったです。そして、ついに大物のヒラマサがヒット！かなりの引きで、腕がパンパンになりましたが、無事に釣り上げることができました。この魚の力強さと美しさには感動しました。次回もこのサイズを狙いたいと思います！'
  },
  methods: {
    async fetchItinerary() {
      try {
        const response = await fetch(`${process.env.BASE_URL}/reservations/${this.$route.params.id}`)
        const data = await response.json()
        this.diary = {
          id: data.id,
          date: data.from_date.split('T')[0],
          image: `${process.env.BASE_URL}/images${data.diary_photo_path}`,
          title: data.diary_title,
          content: data.diary_description
        }
        this.diaryContent = data.diary_description
        this.itinerary = data.travelSpotItineraries.map(item => ({
          type: item.type,
          takeTime: item.take_time,
          price: item.price,
          title: item.title,
          description: item.description
        }))
        data.travelSpotItineraries.map(item => item.items.map(item => this.bring.push(item.name)))
        console.log(data)
      } catch (error) {
        console.error('Failed to fetch itinerary:', error)
      }
    },
    onDiaryClick() {
      this.$router.go()
    },
    async updateDiary(updatedDescription) {
      console.log('updateDiary', updatedDescription)
      try {
        const response = await fetch(`${process.env.BASE_URL}/reservations/${this.$route.params.id}/diary_description`, {
          method: 'PATCH',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({
            Description: updatedDescription,
            user_id: this.userInfo.id
          })
        })

        if (response.ok) {
          console.log('Diary updated successfully')
          this.diary.content = updatedDescription
          this.diaryContent = updatedDescription
        } else {
          console.error('Failed to update diary')
        }
      } catch (error) {
        console.error('Error updating diary:', error)
      }
    }
  }
}
</script>

<template>
  <v-container class="mb-8">
    <h2 class="text-center mb-4">日記</h2>
    <ReservationDiary
      v-model="diaryContent"
      :diary="diary"
      class="mb-8"
      :is-history="true"
      @camera-click="onDiaryClick"
      @update-diary="updateDiary"
    />
    <h2 class="text-center">旅程</h2>
    <NewDiaryItinerary
      :bring="bring"
      :itinerary="itinerary"
    />
  </v-container>
</template>

<style scoped>

</style>
