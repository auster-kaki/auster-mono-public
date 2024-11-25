<script setup lang="ts">
import { ref } from 'vue'
import axios from 'axios'
import DiaryCarousel from '@/components/c/diary/DiaryCarousel.vue'

const baseUrl = ref('https://api.example.com')
const endpoint = ref('/users')
const method = ref('GET')
const pathParams = ref('')
const queryParams = ref('')
const requestBody = ref('')
const response = ref('')

const diaryData = ref({
  id: 1,
  title: '美しい富士山',
  date: '2023-05-01',
  location: '山梨県',
  content:
    '富士山の麓でキャンプを楽しみました。澄んだ空気と雄大な景色に感動しました。',
  image: 'https://example.com/fuji.jpg'
})

const sendRequest = async () => {
  try {
    const url = `${baseUrl.value}${endpoint.value}${pathParams.value}`
    const config = {
      method: method.value,
      url,
      params: queryParams.value ? JSON.parse(queryParams.value) : {},
      data: requestBody.value ? JSON.parse(requestBody.value) : {}
    }
    const result = await axios(config)
    response.value = JSON.stringify(result.data, null, 2)
  } catch (error: any) {
    response.value = `Error: ${error.message}`
  }
}
</script>

<template>
  <div>
    <h1>Diaryテスト画面</h1>
    <diary-carousel
      :diary="{
          id: 1,
          date: '2024/11/23',
          image: 'choshi.jpg',
          title: '大物ヒラマサとの出会い',
          content: '今日は早朝から漁船に乗り、期待に胸を膨らませて出航しました。風は少し冷たかったけれど、海の静けさが心地よかったです。そして、ついに大物のヒラマサがヒット！かなりの引きで、腕がパンパンになりましたが、無事に釣り上げることができました。この魚の力強さと美しさには感動しました。次回もこのサイズを狙いたいと思います！'
        }" />
    <div>
      <h1>DiaryCarouselテスト画面</h1>
    </div>
  </div>
</template>

<style scoped>
div {
  margin-bottom: 10px;
}

label {
  display: block;
}

input,
select,
textarea {
  width: 100%;
  max-width: 400px;
}

textarea {
  height: 100px;
}

pre {
  background-color: #f0f0f0;
  padding: 10px;
  white-space: pre-wrap;
  word-wrap: break-word;
}
</style>
