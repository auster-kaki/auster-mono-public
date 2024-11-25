<script lang="ts">
import { defineComponent, ref } from 'vue'

interface Experience {
  id: number
  title: string
  description: string
  image: string
  isFurusatoNozei: boolean
}

export default defineComponent({
  name: '_id',
  layout: 'dokokani',
  setup() {
    const regionInfo = ref({
      name: '東京',
      celebrityContent: {
        videoUrl: 'https://example.com/video.mp4',
        title: '東京の魅力',
        description: '東京の隠れた名所を紹介します'
      }
    })

    const experiences = ref<Experience[]>([
      {
        id: 1,
        title: '浅草寺参拝体験',
        description: '浅草寺で伝統的な参拝方法を学びます',
        image: 'https://example.com/asakusa.jpg',
        isFurusatoNozei: true
      },
      {
        id: 2,
        title: '築地市場ツアー',
        description: '新鮮な魚介類と活気ある市場の雰囲気を体験',
        image: 'https://example.com/tsukiji.jpg',
        isFurusatoNozei: false
      },
      // Add more experiences as needed
    ])

    const fetchRegionInfo = async () => {
      // Commented out API call
      // const response = await fetch(`/api/region/${route.params.id}`)
      // regionInfo.value = await response.json()
    }

    const fetchExperiences = async () => {
      // Commented out API call
      // const response = await fetch(`/api/experiences/${route.params.id}`)
      // experiences.value = await response.json()
    }

    const goToDiaryGeneration = (experienceId: number) => {
      // Navigate to diary generation page
      // You might want to use vue-router here
      console.log(`Generating diary for experience ${experienceId}`)
    }

    return {
      regionInfo,
      experiences,
      goToDiaryGeneration
    }
  }
})
</script>

<template>
  <v-container>
    <v-row>
      <v-col cols="12">
        <h1>{{ regionInfo.name }}の体験</h1>
      </v-col>
    </v-row>

    <v-row>
      <v-col cols="12">
        <h2>ご当地有名人限定コンテンツ</h2>
        <v-card>
          <v-card-text>
            <video controls :src="regionInfo.celebrityContent.videoUrl" width="100%"></video>
            <h3>{{ regionInfo.celebrityContent.title }}</h3>
            <p>{{ regionInfo.celebrityContent.description }}</p>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>

    <v-row>
      <v-col cols="12">
        <h2>体験一覧</h2>
      </v-col>
    </v-row>

    <v-row>
      <v-col v-for="experience in experiences" :key="experience.id" cols="12" sm="6">
        <v-card>
          <v-img :src="experience.image" height="200px" cover></v-img>
          <v-card-title>{{ experience.title }}</v-card-title>
          <v-card-subtitle v-if="experience.isFurusatoNozei">ふるさと納税対応</v-card-subtitle>
          <v-card-text>{{ experience.description }}</v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="primary" @click="goToDiaryGeneration(experience.id)">
              選択
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<style scoped>
/* Add any scoped styles here if needed */
</style>
