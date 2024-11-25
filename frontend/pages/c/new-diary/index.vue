<template>
  <div>
    <v-stepper v-model="currentStep" class="mb-16" flat max-width="1200px" style="margin: 0 auto;">
      <v-stepper-header style="box-shadow: unset !important;">
        <v-stepper-step :complete="currentStep > 1" step="1" />
        <v-divider />
        <v-stepper-step :complete="currentStep > 2" step="2" />
        <v-divider />
        <v-stepper-step :complete="currentStep > 3" step="3" />
        <v-divider />
        <v-stepper-step :complete="currentStep > 4" step="4" />
        <v-divider />
        <v-stepper-step :complete="currentStep > 5" step="5" />
        <v-divider />
        <v-stepper-step :complete="currentStep > 6" step="6" />
      </v-stepper-header>

      <v-stepper-items>
        <v-stepper-content  style="padding: 8px" step="1">
          <departure-selection
            :departure-place="departureForm.departurePlace"
            :departure-date="departureForm.departureDate"
            :return-date="departureForm.returnDate"
            :departure-time="departureForm.departureTime"
            :return-time="departureForm.returnTime"
            :interests="departureForm.interests"
            @submit="handleDepatureSelectionSubmit"
          />
        </v-stepper-content>

        <v-stepper-content style="padding: 8px" step="2">
          <destination-selection
            @destination-selected="handleDestinationSelected"
          />
          <v-btn text @click="currentStep -= 1">
            戻る
          </v-btn>
        </v-stepper-content>

        <v-stepper-content style="padding: 8px" step="3">
          <experience-selection
            :experiences="experienceForm.experiences"
            @click="handleSelectExperience"
          />
          <v-container>
            <v-row class="mt-4 pb-4">
              <v-btn text @click="currentStep -= 1">戻る</v-btn>
            </v-row>
          </v-container>
        </v-stepper-content>
        <v-stepper-content style="padding: 8px" step="4">
          <diary-carousel :diary="createdDiary" @select="handleSelectDiary" />
          <v-btn text @click="currentStep -= 1">
            戻る
          </v-btn>
        </v-stepper-content>
        <v-stepper-content style="padding: 8px" step="5">
          <h2 class="text-center">旅程</h2>
          <NewDiaryItinerary
            :bring="bring"
            :itinerary="itinerary"
          />
          <v-container>
            <v-row class="mt-4 pb-4">
              <v-btn text @click="currentStep -= 1">戻る</v-btn>
              <v-spacer />
              <v-btn color="primary" @click="onGoToConfirm">申し込み確認画面へ</v-btn>
            </v-row>
          </v-container>
        </v-stepper-content>
        <v-stepper-content  style="padding: 8px" step="6">
          <NewDiaryConfirm
            :itinerary="itinerary"
            @confirm="handleConfirm"
          />
          <v-container>
            <v-row class="mt-4 pb-4">
              <v-btn text @click="currentStep -= 1">戻る</v-btn>
              <v-spacer />
              <v-btn color="primary" @click="handleConfirm">予約する</v-btn>
            </v-row>
          </v-container>
        </v-stepper-content>
      </v-stepper-items>
    </v-stepper>
    <v-overlay :value="loading">
      <v-progress-circular indeterminate size="64"></v-progress-circular>
    </v-overlay>
  </div>
</template>

<script>
import DepartureSelection from '~/components/c/new-diary/departure-selection.vue'
import DestinationSelection from '~/components/c/new-diary/destination-selection.vue'
import ExperienceSelection from '~/components/c/new-diary/experience-selection.vue'
import NewDiaryItinerary from '~/components/c/new-diary/itinerary.vue'
import NewDiaryConfirm from '~/components/c/new-diary/confirm.vue'
import { useUserStore } from '~/store/user'
import DiaryCarousel from '~/components/c/diary/DiaryCarousel.vue'

export default {
  components: {
    DiaryCarousel,
    NewDiaryConfirm,
    NewDiaryItinerary,
    ExperienceSelection,
    DestinationSelection,
    DepartureSelection
  },
  layout: 'mobile',
  data() {
    return {
      currentStep: 1,
      departureForm: {
        departurePlace: '',
        departureDate: null,
        returnDate: null,
        departureTime: null,
        returnTime: null,
        interests: []
      },
      destinationForm: {
        id: ''
      },
      experienceForm: {
        experiences: []
      },
      selectedTravelSpotId: '',
      createdDiary: {
        id: 1,
        date: '2024/11/23',
        image: 'http://localhost:3000/auster-mono/_nuxt/static/destination/choshi.jpg',
        title: 'ダミーデータ！！！',
        content: '今日は早朝から漁船に乗り、期待に胸を膨らませて出航しました。風は少し冷たかったけれど、海の静けさが心地よかったです。そして、ついに大物のヒラマサがヒット！かなりの引きで、腕がパンパンになりましたが、無事に釣り上げることができました。この魚の力強さと美しさには感動しました。次回もこのサイズを狙いたいと思います！'
      },
      bring: [],
      itinerary: [],
      userInfo: {
        id: ''
      },
      vendorMaster: [
        {
          'id': 'can1',
          'name': 'CampsiteTORAMI',
          'address': '千葉県長生郡一宮町東浪見1611'
        },
        {
          'id': 'can2',
          'name': 'Beach Camp 九十九里',
          'address': '千葉県大網白里市四天木2761-40'
        },
        {
          'id': 'can3',
          'name': '銚子電気鉄道',
          'address': '千葉県銚子市新生町2丁目297番地'
        },
        {
          'id': 'gyo1',
          'name': '銚子市漁業協同組合',
          'address': '千葉県銚子市川口町 2丁目6528番地'
        },
        {
          'id': 'gyo2',
          'name': '銚子市生活環境課 清掃美化班',
          'address': '千葉県銚子市若宮町1-1 （銚子市役所本庁舎4階）'
        },
        {
          'id': 'gyo3',
          'name': '銚子市観光課',
          'address': '千葉県銚子市若宮町1-1 （銚子市役所本庁舎4階）'
        }
      ],
      loading: false
    }
  },
  mounted() {
    const userStore = useUserStore()
    userStore.initializeUser()
    this.userInfo = userStore.userInfo
  },
  methods: {
    handleDepatureSelectionSubmit(model) {
      this.departureForm = model
      this.currentStep += 1
    },
    async handleDestinationSelected(id) {
      this.loading = true
      this.destinationForm.id = id
      const params = new URLSearchParams({
        user_id: this.userInfo.id,
        hobby_id: this.departureForm.interests
      })

      try {
        const response = await fetch(`${process.env.BASE_URL}/travel_spots?${params.toString()}`, {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json'
          }
        })

        if (!response.ok) {
          throw new Error('Network response was not ok')
        }
        const data = await response.json()

        let experiences = []
        if (data.length >= 5) {
          experiences = data.slice(0, 4).map((spot, _i) => ({
            id: spot.ID,
            vendorId: spot.VendorID,
            image: spot.PhotoPath ? `${process.env.BASE_URL}/images${spot.PhotoPath}` : 'https://placehold.jp/300x200.png',
            title: spot.Name,
            description: spot.Description,
            address: spot.Address,
            hasFurusatoNozei: Math.random() < 0.5 // ランダムにtrueかfalseを設定
          }))
        } else {
          // そのまま入れる
          experiences = data.map((spot, _i) => ({
            id: spot.ID,
            vendorId: spot.VendorID,
            image: spot.PhotoPath ? `${process.env.BASE_URL}/images${spot.PhotoPath}` : 'https://placehold.jp/300x200.png',
            title: spot.Name,
            description: spot.Description,
            address: spot.Address,
            hasFurusatoNozei: Math.random() < 0.5 // ランダムにtrueかfalseを設定
          }))
        }
        console.log(this.experienceForm.experiences)

        this.experienceForm = {
          experiences
        }
        this.currentStep += 1
      } catch (error) {
        console.error('There was a problem with the fetch operation:', error)
      } finally {
        this.loading = false
      }
    },
    async handleSelectExperience(id) {
      this.loading = true
      try {
        const response = await fetch(`${process.env.BASE_URL}/travel_spots/${id}/diary`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({
            user_id: this.userInfo.id,
            hobby_id: this.departureForm.interests,
            date: this.departureForm.departureDate
          })
        })
        this.selectedTravelSpotId = id
        if (!response.ok) {
          throw new Error('Network response was not ok')
        }
        this.currentStep += 1
        // JSONレスポンスを直接パース
        const data = await response.json()

        // 画像パスを使用して画像URLを構築
        const imageUrl = data.PhotoPath
          ? `${process.env.BASE_URL}/images/${data.PhotoPath}`
          : 'https://placehold.jp/300x200.png'

        this.createdDiary = {
          id: data.ID,
          date: this.departureForm.departureDate,
          image: imageUrl,
          title: data.Title,
          content: data.Body
        }

      } catch (error) {
        console.error('There was a problem with the fetch operation:', error)
      } finally {
        this.loading = false
      }
    },
    async handleSelectDiary(_id) {
      this.loading = true
      const params = new URLSearchParams({
        user_id: this.userInfo.id
      })
      try {
        const response = await fetch(`${process.env.BASE_URL}/travel_spots/${this.selectedTravelSpotId}/itineraries?${params.toString()}`, {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json'
          }
        })
        const data = await response.json()
        this.bring = data.Items.map((item) => item.Name)
        this.itinerary = data.TravelSpotItineraries.map((itinerary) => ({
          id: itinerary.ID,
          travelSpotId: itinerary.TravelSpotID,
          kind: itinerary.Kind,
          title: itinerary.Title,
          description: itinerary.Description,
          takeTime: itinerary.TakeTime,
          price: itinerary.Price,
          order: itinerary.Order
        }))
        this.currentStep += 1
      } catch (error) {
        console.error('There was a problem with the fetch operation:', error)
      } finally {
        this.loading = false
      }
    },
    onGoToConfirm() {
      this.currentStep += 1
    },
    async handleConfirm() {
      this.loading = true
      try {
        await fetch(`${process.env.BASE_URL}/reservations`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({
            user_id: this.userInfo.id,
            travel_spot_id: this.selectedTravelSpotId,
            travel_spot_diary_id: this.createdDiary.id.toString(),
            from_date: this.departureForm.departureDate,
            to_date: this.departureForm.returnDate
          })
        })

        const selectedExperience = this.experienceForm.experiences.find((experience) => experience.id === this.selectedTravelSpotId)
        const vendor = this.vendorMaster.find((vendor) => vendor.id === selectedExperience.vendorId)

        await fetch(`${process.env.BASE_URL}/encounters`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({
            user_id: this.userInfo.id,
            name: vendor.name,
            place: vendor.address,
            date: this.departureForm.departureDate,
            description: this.createdDiary.title
          })
        })
        this.$router.push({ path: '/c/reservations', query: { reservation: 'success' } })
      } catch (error) {
        console.error('There was a problem with the fetch operation:', error)
      } finally {
        this.loading = false
      }
    }
  }
}
</script>
