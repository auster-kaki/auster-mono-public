<template>
  <v-container class="pa-2">
    <v-card style="position: relative;" elevation-4>
      <v-container>
        <v-row class="pa-2 align-center">
          <div class="custom-button mr-2">
            {{ diary.date.slice(0, 4) }}<br>
            {{ diary.date.slice(5).replace('/', '') }}
          </div>
          <h3 class="text-center font-weight-bold">{{ diary.title }}</h3>
          <v-spacer />
          <v-icon size="24" @click="showBookmarkMessage">mdi-bookmark</v-icon>
        </v-row>
        <v-row class="px-2">
          <v-img
            :src="diary.image"
            style="border-radius: 4px"
            :alt="diary.title"
            cover
            height="270"
            width="100%"
          />
        </v-row>
      </v-container>

      <v-card-text class="text-body-1" style="overflow-y: auto;">
        <div v-if="!isHistory">{{ diary.content }}</div>
        <v-textarea
          v-else
          v-model="localContent"
          label="日記本文"
          auto-grow
          rows="7"
          outlined
          dense
          @input="updateContent"
        />
      </v-card-text>
      <v-card-actions v-if="!!isHistory">
        <v-spacer />
        <v-btn dense class="primary" @click="updateDiary">体験日記更新</v-btn>
      </v-card-actions>
      <v-sheet v-if="!isHistory" class="d-flex justify-center" style="position: relative;">
        <v-btn
          color="primary"
          class="px-4 py-2 mb-2"
          @click="handleSelect"
        >
          <span class="text-white font-weight-bold">この体験日記にする</span>
        </v-btn>
      </v-sheet>
    </v-card>
    <v-snackbar v-model="snackbar" :timeout="2000" color="accent" class="text-center">
      {{ snackbarMessage }}
    </v-snackbar>
  </v-container>
</template>

<script>
import { ref, watch } from 'vue'

export default {
  props: {
    diary: {
      type: [Object],
      required: true
    },
    isHistory: {
      type: Boolean,
      default: false
    }
  },
  setup(props, { emit }) {
    const snackbar = ref(false)
    const snackbarMessage = ref('')
    const localContent = ref(props.diary.content)

    watch(() => props.diary.content, (newContent) => {
      localContent.value = newContent
    })

    const updateContent = (value) => {
      emit('update:content', value)
    }

    const showMessage = (message) => {
      snackbarMessage.value = message
      snackbar.value = true
    }

    const showDownloadMessage = () => {
      showMessage('日記がダウンロードできる予定です')
    }

    const showBookmarkMessage = () => {
      showMessage('ブックマークして後から見返せるようになる予定です')
    }

    const showShareMessage = () => {
      showMessage('SNSでシェアできるようになる予定です')
    }

    const handleSelect = () => {
      emit('select', props.diary.id)
    }

    const updateDiary = () => {
      showMessage('日記が更新されました')
    }

    return {
      snackbar,
      snackbarMessage,
      localContent,
      updateContent,
      showDownloadMessage,
      showBookmarkMessage,
      showShareMessage,
      handleSelect,
      updateDiary
    }
  }
}
</script>

<style scoped>
:deep(.v-carousel__controls__item.v-btn) {
  /* アクティブなドットの色を変更する場合 */
  color: rgb(0, 0, 0, 0.54);
}

/* カードアクション全体のスタイリング */
.card-actions {
  display: flex;
  justify-content: space-between; /* ボタンとアイコンを左右に配置 */
  align-items: center; /* 垂直方向を中央揃え */
}

.custom-button {
  display: inline-block;
  padding: 4px 8px;
  background-color: #4E97E0; /* 本当はprimaryから取得してくるべき */
  color: white;
  text-align: center;
  border-radius: 2px;
  font-size: 16px;
  font-weight: bold;
}
</style>
