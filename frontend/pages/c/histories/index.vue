<template>
  <v-container class="mb-16">
    <v-tabs v-model="activeTab">
      <v-tab>日記</v-tab>
      <v-tab>出会った人々</v-tab>
    </v-tabs>

    <v-tabs-items v-model="activeTab">
      <v-tab-item>
        <v-container>
          <v-row>
            <v-col
              v-for="(diary, index) in diaries"
              :key="index"
              cols="12"
              sm="6"
              md="4"
            >
              <diary-card
                :value="diary"
                @click="openDiaryModal(diary)"
              ></diary-card>
            </v-col>
          </v-row>
        </v-container>
      </v-tab-item>

      <v-tab-item>
        <v-list>
          <v-list-item
            v-for="(encounter, index) in encounters"
            :key="index"
            @click="openEncounterModal(encounter)"
          >
            <v-list-item-avatar>
              <v-img :src="encounter.avatar"></v-img>
            </v-list-item-avatar>
            <v-list-item-content>
              <v-list-item-title>{{ encounter.name }}</v-list-item-title>
              <v-list-item-subtitle>{{ encounter.place }}</v-list-item-subtitle>
            </v-list-item-content>
          </v-list-item>
        </v-list>
      </v-tab-item>
    </v-tabs-items>
    <v-dialog v-model="diaryModalOpen">
      <history-diary :diary="selectedDiary" :is-offer="selectedDiary?.isOffer" />
    </v-dialog>
    <v-dialog v-model="encounterModalOpen" max-width="600">
      <v-card v-if="selectedEncounter">
        <v-card-title>{{ selectedEncounter.name }}</v-card-title>
        <v-card-subtitle>{{ selectedEncounter.place }}</v-card-subtitle>
        <v-card-text>
          <v-list>
            <v-list-item v-for="(detail, index) in selectedEncounter.details" :key="index">
              <v-list-item-content>
                <v-list-item-subtitle class="text-subtitle-2">{{ detail.date }}</v-list-item-subtitle>
                <v-list-item-title class="text-h6">{{ detail.description }}</v-list-item-title>
              </v-list-item-content>
            </v-list-item>
          </v-list>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="primary" outlined @click="openChat(selectedEncounter)">チャットを開く</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <v-snackbar v-model="snackbar" color="accent" :timeout="3000">
      チャットが開始できる想定です
      <template #action="{ attrs }">
        <v-btn text v-bind="attrs" @click="snackbar = false">
          閉じる
        </v-btn>
      </template>
    </v-snackbar>
  </v-container>
</template>

<script>
import DiaryCard from '@/components/c/diary/DiaryCard.vue'
import { useUserStore } from '~/store/user'
import HistoryDiary from '~/components/c/diary/HistoryDiary.vue'

export default {
  components: {
    HistoryDiary,
    DiaryCard
  },
  layout: 'mobile',
  data() {
    return {
      userInfo: {},
      activeTab: null,
      diaries: [],
      encounters: [],
      diaryModalOpen: false,
      encounterModalOpen: false,
      snackbar: false,
      selectedDiary: null,
      selectedEncounter: null,
      diaryDetail: {
        id: '',
        date: '',
        image: '',
        title: '',
        content: '',
        isOffer: false,
      }
    }
  },
  async mounted() {
    const userStore = useUserStore()
    userStore.initializeUser()
    this.userInfo = userStore.userInfo
    await this.fetchHistories()
  },
  methods: {
    async fetchHistories() {
      try {
        const params = new URLSearchParams({
          user_id: this.userInfo.id,
          filter: 'done'
        })
        const response = await fetch(`${process.env.BASE_URL}/reservations?${params.toString()}`, {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json'
          }
        })
        const res = await response.json()
        this.diaries = res.reservations.map(item => ({
          image: `${process.env.BASE_URL}/images${item.diary_photo_path}`,
          id: item.id,
          title: item.travel_spot_title,
          date: item.from_date.split('T')[0],
          location: '銚子',
          content: item.travel_spot_description,
          isOffer: item.is_offer
        })).sort((a, b) => {
          if (b.isOffer !== a.isOffer) {
            return b.isOffer - a.isOffer
          }
          return new Date(b.date) - new Date(a.date)
        })

      } catch (error) {
        console.error('履歴の取得に失敗しました', error)
      }

      try {
        const params = new URLSearchParams({
          user_id: this.userInfo.id
        })
        const response = await fetch(`${process.env.BASE_URL}/encounters?${params.toString()}`, {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json'
          }
        })
        const data = await response.json()

        const groupedEncounters = data.reduce((acc, curr) => {
          if (!acc[curr.Name]) {
            acc[curr.Name] = {
              name: curr.Name,
              place: curr.Place,
              details: []
            }
          }
          acc[curr.Name].details.push({
            date: this.formatDate(curr.Date),
            description: curr.Description
          })
          return acc
        }, {})

        this.encounters = Object.values(groupedEncounters)
      } catch (error) {
        console.error('関係事業者の取得に失敗しました', error)
      }
    },
    async fetchHistoryDetail(id) {
      try {
        const response = await fetch(`${process.env.BASE_URL}/reservations/${id}`)
        const data = await response.json()
        this.diaryDetail = {
          id: data.id,
          date: data.from_date.split('T')[0],
          image: `${process.env.BASE_URL}/images${data.diary_photo_path}`,
          title: data.diary_title,
          content: data.diary_description,
          isOffer: data.is_offer
        }
      } catch (error) {
        console.error('Failed to fetch itinerary:', error)
      }
    },
    async openDiaryModal(diary) {
      await this.fetchHistoryDetail(diary.id)
      console.log(diary, this.diaries)
      this.selectedDiary = this.diaryDetail
      this.diaryModalOpen = true
    },
    openEncounterModal(encounter) {
      this.selectedEncounter = encounter
      this.encounterModalOpen = true
    },
    openChat() {
      this.snackbar = true
    },
    formatDate(dateString) {
      if (!dateString) {
        console.log('formatDate target not found', dateString)
        return dateString
      }
      return dateString.split('T')[0]
    }
  }
}
</script>
