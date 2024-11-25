<script setup lang="ts">
import { defineProps, defineEmits } from 'vue'

interface Diary {
  id: number
  image: string
  title: string
  date: string
  location: string
  content: string
  isOffer?: boolean
}

const props = defineProps<{
  value: Diary
}>()

const emit = defineEmits(['click'])

const openDiaryModal = () => {
  emit('click', props.value)
}

</script>

<template>
  <v-card @click="openDiaryModal">
    <v-img :src="value.image" height="300" style="position:relative; object-fit: cover; object-position: center top;">
      <v-chip
        v-if="value.isOffer"
        style="position:absolute; bottom: 15px; right: 10px;"
        class="elevation-4"
        color="accent"
        :class="['animate-bounce']"
      >
        <v-icon size="24">mdi-email</v-icon>
        <strong>スペシャルオファー！</strong>
      </v-chip>
    </v-img>
    <v-card-title>{{ value.title }}</v-card-title>
    <v-card-subtitle>{{ value.date }} - {{ value.location }}</v-card-subtitle>
    <v-card-text>{{ value.content ? value.content.substring(0, 100) + '...' : '' }}</v-card-text>
  </v-card>
</template>

<style scoped>
@keyframes bounce {
  0%, 100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-10px);
  }
}

.animate-bounce {
  animation: bounce 2s infinite;
}
</style>
