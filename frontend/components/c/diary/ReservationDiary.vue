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
            style="border-radius: 4px; position:relative"
            :alt="diary.title"
            cover
            height="270"
            width="100%"
          >
            <v-chip
              style="position:absolute; bottom: 10px; right: 10px; background: #fafafa !important;"
              class="elevation-4"
              text-color="accent"
              @click="triggerFileInput"
            >
              <v-icon size="24" color="accent">mdi-camera</v-icon>
              画像更新
            </v-chip>
          </v-img>
        </v-row>
      </v-container>

      <v-card-text class="text-body-1" style="overflow-y: auto;">
        <v-textarea
          :value="value"
          label="日記本文"
          auto-grow
          rows="7"
          outlined
          dense
          @input="updateContent"
        />
      </v-card-text>
      <v-card-actions>
        <v-spacer />
        <v-btn dense class="primary" @click="updateDiary">日記更新</v-btn>
      </v-card-actions>
    </v-card>
    <v-snackbar v-model="snackbar" :timeout="2000" color="accent" class="text-center">
      {{ snackbarMessage }}
    </v-snackbar>
    <!-- 非表示のファイル入力 -->
    <input
      ref="fileInput"
      type="file"
      style="display: none;"
      @change="handleFileChange"
    />
  </v-container>
</template>

<script>
import { ref } from 'vue'
import { useUserStore } from '~/store/user'

export default {
  props: {
    value: {
      type: String,
      required: true
    },
    diary: {
      type: Object,
      required: true
    },
    isHistory: {
      type: Boolean,
      default: false
    }
  },
  emits: ['input', 'camera-click', 'update-diary'],
  setup(props, { emit }) {

    const snackbar = ref(false)
    const snackbarMessage = ref('')
    const localContent = ref(props.diary.content)
    const fileInput = ref(null)

    const updateContent = (value) => {
      emit('input', value)
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

    const handleSelect = () => {
      emit('select', props.value.id)
    }

    const updateDiary = () => {
      emit('update-diary', props.value)
      showMessage('日記が更新されました')
    }

    const handleChipClick = () => {
      emit('camera-click')
    }

    const triggerFileInput = () => {
      fileInput.value.click()
    }

    // ファイルが選択されたときの処理
    const handleFileChange = async (event) => {
      const file = event.target.files[0]
      if (!file) return

      const userStore = useUserStore()
      userStore.initializeUser()
      const userInfo = userStore.userInfo

      const formData = new FormData()
      formData.append('photo', file)
      // formData.append('user_id', userInfo.id)
      try {
        await fetch(`${process.env.BASE_URL}/reservations/${props.diary.id}/diary_photo/users/${userInfo.id}`, {
          method: 'PATCH',
          body: formData
        })
        emit('camera-click')
      } catch (error) {
        console.error('Error uploading photo:', error)
        showMessage('写真のアップロードに失敗しました')
      }
    }

    return {
      snackbar,
      snackbarMessage,
      localContent,
      fileInput,
      updateContent,
      showDownloadMessage,
      showBookmarkMessage,
      handleSelect,
      updateDiary,
      handleChipClick,
      triggerFileInput,
      handleFileChange,
    }
  },
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
