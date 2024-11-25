<script lang="ts">
import { defineComponent, ref } from 'vue'

interface Destination {
  id: number
  name: string
  image: string
}

export default defineComponent({
  name: 'tabisaki',
  layout: 'dokokani',
  setup() {
    const destinations = ref<Destination[]>([])

    const fetchDestinations = async () => {
      // API通信のコード（現在はコメントアウト）
      // const response = await fetch('/api/destinations')
      // destinations.value = await response.json()

      // ダミーデータを返却
      destinations.value = [
        { id: 1, name: '地域名1', image: '/images/destination-1.jpg' },
        { id: 2, name: '地域名2', image: '/images/destination-2.jpg' },
        { id: 3, name: '地域名3', image: '/images/destination-3.jpg' },
        { id: 4, name: '地域名4', image: '/images/destination-4.jpg' },
      ]
    }

    fetchDestinations()

    return {
      destinations
    }
  }
})
</script>

<template>
  <v-container class="tabisaki-container">
    <v-row>
      <v-col cols="12">
        <h1 class="text-h4 text-center mb-6">あなたにおすすめの旅先はこちら！</h1>
      </v-col>
    </v-row>
    <v-row>
      <v-col v-for="(destination, index) in destinations" :key="index" cols="12" sm="6">
        <v-card class="tabisaki-item">
          <v-card-title class="tabisaki-badge text-caption">
            ご当地有名人限定コンテンツ公開中！
          </v-card-title>
          <v-img
            :src="destination.image"
            :alt="`旅先${index + 1}の画像`"
            height="200"
            cover
          ></v-img>
          <v-card-text>
            <h2 class="text-h6 mb-2">{{ destination.name }}</h2>
            <v-btn
              color="primary"
              text
              :to="destination.id"
            >
              地域情報を見る
            </v-btn>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<style scoped>
.tabisaki-badge {
  position: absolute;
  top: 10px;
  right: 10px;
  background-color: #ff6b6b !important;
  color: white;
  padding: 5px 10px;
  border-radius: 20px;
  font-size: 0.8em;
  font-weight: bold;
  z-index: 1;
}
</style>
