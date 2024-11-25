<script lang="ts">
import { defineComponent, ref } from 'vue'

export default defineComponent({
  name: 'ExperienceSelection',
  props: {
    experiences: {
      type: Array,
      required: true
    }
  },
  setup() {
    const selectedLanguage = ref('ja')
    const videoSrc = ref('/auster-mono/new-diary/choshi.mp4')

    const changeLanguage = (lang: string) => {
      selectedLanguage.value = lang
          switch (lang) {
            case 'ja':
              videoSrc.value = '/auster-mono/new-diary/choshi.mp4'
              break
            case 'en':
              videoSrc.value = '/auster-mono/new-diary/eigo.mp4'
              break
            case 'cn':
              videoSrc.value = '/auster-mono/new-diary/chinese.mp4'
              break
            default:
              videoSrc.value = '/auster-mono/new-diary/choshi.mp4'
          }    }

    return {
      selectedLanguage,
      videoSrc,
      changeLanguage
    }
  },
  methods: {
    selectExperience(id: string) {
      this.$emit('click', id)
    }
  }
})
</script>

<template>
  <v-container>
    <v-row>
      <v-col cols="12">
        <h3 class="text-h5 mb-4 text-grey">銚子について</h3>
        <v-card>
          <v-card-text>
            <v-container>
              <v-row class="d-flex justify-center">
                <video
                  controls
                  class="responsive-video" :src="videoSrc" />
              </v-row>
              <v-row>
                <v-spacer />
                <v-select
                  v-model="selectedLanguage"
                  :items="[
                  { text: '日本語', value: 'ja' },
                  { text: 'English', value: 'en' },
                  { text: 'Chinese', value: 'cn' }
                ]"
                  class="ml-4"
                  style="max-width: 150px;"
                  @change="changeLanguage"
                ></v-select>
              </v-row>
            </v-container>


            <v-card-title class="text-h5">海と歴史が織りなす町</v-card-title>
            <v-card-subtitle>千葉県の東端に位置する銚子市。豊かな自然と歴史的な街並みが魅力の町をご紹介します。
            </v-card-subtitle>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12">
        <h3 class="text-h5 mb-4">体験</h3>
        <v-row>
          <v-col v-for="experience in experiences" :key="experience.id" cols="12" sm="6" md="6" lg="6">
            <v-card>
              <v-img :src="experience.image" :alt="experience.title" height="200" style="position:relative" cover>
                <v-chip
                  v-if="experience.hasFurusatoNozei" color="accent" outlined
                  style="position:absolute; top: 10px; right: 10px; background: #fafafa !important;">
                  <strong>ふるさと納税対象</strong>
                </v-chip>
              </v-img>
              <v-card-title>{{ experience.title }}</v-card-title>
              <v-card-text>
                <p>{{ experience.description }}</p>
              </v-card-text>
              <v-card-actions>
                <v-spacer />
                <v-btn class="primary" @click="selectExperience(experience.id)">この体験日記を見る</v-btn>
              </v-card-actions>
            </v-card>
          </v-col>
        </v-row>
      </v-col>
    </v-row>
  </v-container>
</template>

<style scoped>
/* レスポンシブ動画スタイル */
.responsive-video {
  width: 100%; /* 画面幅に合わせる */
  max-width: 800px; /* 最大幅を設定 */
  height: auto; /* 縦横比を維持 */
}
</style>
