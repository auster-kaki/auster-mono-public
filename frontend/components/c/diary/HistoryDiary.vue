<template>
  <div>
    <v-card style="position: relative;" flat>
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
            style="border-radius: 4px;position:relative"
            :alt="diary.title"
            cover
            height="270"
            width="100%"
          >
            <v-chip
              v-if="isOffer"
              style="position:absolute; bottom: 15px; right: 10px;"
              class="elevation-4"
              color="accent"
              :class="['animate-bounce']"
            >
              <v-icon size="24">mdi-email</v-icon>
              <strong>スペシャルオファー！</strong>
            </v-chip>
          </v-img>
        </v-row>
      </v-container>

      <v-card-text class="text-body-1" style="overflow-y: auto;">
        <div>{{ diary.content }}</div>
      </v-card-text>
      <v-card-actions>
        <v-btn
          v-if="isOffer"
          color="primary"
          class="px-4 py-2 mb-2"
          @click="acceptTheOffer"
        >
          <span class="text-white font-weight-bold">オファーを受ける</span>
        </v-btn>
        <v-spacer />
        <v-btn icon @click="showShareMessage">
          <v-icon size="x-large">mdi-share-variant</v-icon>
        </v-btn>
        <v-btn icon class="mr-4" @click="showDownloadMessage">
          <v-icon size="x-large">mdi-download</v-icon>
        </v-btn>
      </v-card-actions>
      <v-dialog v-model="dialog" max-width="300">
        <v-card>
          <v-card-title class="text-center">
            {{ dialogMessage }}
          </v-card-title>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="primary" text @click="dialog = false">閉じる</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </v-card>
  </div>
</template>

<script>
import { ref, watch } from 'vue'

export default {
  props: {
    diary: {
      type: [Object],
      required: true
    },
    isOffer: {
      type: [Boolean],
      default: false
    }
  },
  setup(props, { emit }) {
    const dialog = ref(false)
    const dialogMessage = ref('')
    const localContent = ref(props.diary.content)

    watch(() => props.diary.content, (newContent) => {
      localContent.value = newContent
    })

    const updateContent = (value) => {
      emit('update:content', value)
    }

    const showMessage = (message) => {
      dialogMessage.value = message
      dialog.value = true
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

    const acceptTheOffer = () => {
      showMessage('オファーを受けて、旅程が組まれる予定です')
    }

    const handleSelect = () => {
      emit('select', props.diary.id)
    }

    const updateDiary = () => {
      showMessage('日記が更新されました')
    }

    return {
      dialog,
      dialogMessage,
      localContent,
      updateContent,
      showDownloadMessage,
      showBookmarkMessage,
      showShareMessage,
      acceptTheOffer,
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
