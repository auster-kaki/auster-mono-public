import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useExperienceStore = defineStore('experience', () => {
  const dates = ref([])
  const experiences = ref([])

  const fetchExperiences = async () => {
    // Simulating API call
    const response = await fetch('/api/experiences')
    experiences.value = await response.json()
  }

  const fetchDates = async () => {
    // Simulating API call
    const response = await fetch('/api/dates')
    dates.value = await response.json()
  }

  return {
    dates,
    experiences,
    fetchExperiences,
    fetchDates,
  }
})
