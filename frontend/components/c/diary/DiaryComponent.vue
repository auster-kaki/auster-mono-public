<script setup lang="ts">
import { defineProps } from 'vue'

interface Diary {
  image: string
  title: string
  date: string
  location: string
  content: string
}

const { diary } = defineProps<{
  diary: Diary | null
}>()
</script>

<template>
  <v-card v-if="diary">
    <v-img :src="diary.image" height="300"></v-img>
    <v-card-title>{{ diary.title }}</v-card-title>
    <v-card-subtitle
    >{{ diary.date }} -
      {{ diary.location }}
    </v-card-subtitle>
    <v-card-text>{{ diary.content }}</v-card-text>
  </v-card>
</template>

<style scoped></style>
