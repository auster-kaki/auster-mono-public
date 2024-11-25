<template>
  <v-container>
    <v-row justify="center" align="center">
      <v-col cols="12">
        <!-- TODO リスト -->
        <v-card v-if="todos.length > 0" class="mb-4">
          <v-card-title>準備リスト</v-card-title>
          <v-list>
            <v-list-item v-for="(todo, index) in todos" :key="index">
              <v-list-item-action>
                <v-checkbox v-model="todo.completed"></v-checkbox>
              </v-list-item-action>
              <v-list-item-content>
                <v-list-item-title>{{ todo.text }}</v-list-item-title>
              </v-list-item-content>
            </v-list-item>
          </v-list>
        </v-card>

        <!-- 申し込み中の予約 -->
        <v-card v-if="reservations.length > 0" class="mb-4">
          <v-card-title>申し込み中の予約</v-card-title>
          <v-list>
            <v-list-item
              v-for="reservation in reservations"
              :key="reservation.id"
            >
              <v-list-item-content>
                <v-list-item-title>{{ reservation.title }}</v-list-item-title>
                <v-list-item-subtitle
                >{{ reservation.date }}
                </v-list-item-subtitle>
              </v-list-item-content>
            </v-list-item>
          </v-list>
        </v-card>

        <!-- 新しい絵日記の通知 -->
        <v-card v-if="newDiaries.length > 0" class="mb-4">
          <v-card-title>新しい絵日記が届いています！</v-card-title>
          <v-list>
            <v-list-item v-for="newDiary in newDiaries" :key="newDiary.id">
              <v-list-item-content>
                <v-list-item-title>{{ newDiary.title }}</v-list-item-title>
              </v-list-item-content>
            </v-list-item>
          </v-list>
        </v-card>

        <!-- 絵日記の表示 -->
        <v-card v-if="currentDiary" class="mb-4">
          <v-card-title>{{ currentDiary.title }}</v-card-title>
          <v-card-subtitle>{{ currentDiary.date }}</v-card-subtitle>
          <v-img :src="currentDiary.image" height="200" contain></v-img>
          <v-card-text>{{ currentDiary.content }}</v-card-text>
          <v-card-actions>
            <v-btn color="primary" @click="reserveDiary">予約する</v-btn>
          </v-card-actions>
        </v-card>

        <!-- 新しい絵日記のモーダル -->
        <v-dialog v-model="showNewDiary" max-width="600px">
          <Diary :diary="newDiary" @close="closeNewDiary" />
        </v-dialog>
      </v-col>
    </v-row>
    <v-snackbar
      v-model="snackbar.show"
      :timeout="snackbar.timeout"
      :color="snackbar.color"
      dense
      :close-text="snackbar.closeText"
      :close-icon="snackbar.closeIcon"
    >
      {{ snackbar.text }}
      <template #action="{ attrs }">
        <v-btn
          text
          v-bind="attrs"
          @click="snackbar.show = false"
        >
          <v-icon>{{ snackbar.closeIcon }}</v-icon>
        </v-btn>
      </template>
    </v-snackbar>
  </v-container>
</template>

<script>
import Diary from '~/components/c/diary/DiaryComponent.vue'

export default {
  name: 'IndexPage',
  components: { Diary },
  layout: 'mobile',
  data() {
    return {
      todos: [],
      reservations: [],
      newDiaries: [],
      showNewDiary: false,
      currentDiary: null,
      snackbar: {
        show: false,
        text: '',
        timeout: 3000,
        color: 'success',
        closeText: '閉じる',
        closeIcon: 'mdi-close'
      }
    }
  },
  mounted() {
    this.fetchTodos()
    this.fetchreservations()
    this.fetchNewDiary()
    this.checkNewDiary()
    if (this.$route.query.reservation === 'success') {
      this.showSnackbar('予約が完了しました！')
    }
    if (this.$route.query.createUser === 'success') {
      this.showSnackbar('ユーザー作成が完了しました！')
    }
  },
  methods: {
    fetchTodos() {
      // API通信のコメントアウト
      // const response = await axios.get('/api/todos')
      // this.todos = response.data

      // ダミーデータ
      this.todos = [
        { text: '持ち物の準備', completed: false },
        { text: '乗車券の準備', completed: true }
      ]
    },
    fetchreservations() {
      // API通信のコメントアウト
      // const response = await axios.get('/api/pending-reservations')
      // this.reservations = response.data

      // ダミーデータ
      this.reservations = [
        { id: 1, title: '山登り体験', date: '2023-06-15' },
        { id: 2, title: '料理教室', date: '2023-06-20' }
      ]
    },
    fetchNewDiary() {
      // 新しい日記を取得してくる
      this.newDiary = {
        title: '',
        tags: [],
        date: '',
        img: ''
      }
    },
    checkNewDiary() {
      // API通信のコメントアウト
      // const response = await axios.get('/api/new-diary')
      // this.newDiary = response.data.hasNewDiary

      // ダミーデータ
      this.newDiary = true
    },
    showNewDiaryClick() {
      // API通信のコメントアウト
      // const response = await axios.get('/api/latest-diary')
      // this.currentDiary = response.data

      // ダミーデータ
      this.currentDiary = {
        date: '2023-06-10',
        title: '素晴らしい山登り体験',
        image: 'https://example.com/mountain.jpg',
        content: '今日は素晴らしい山登り体験をしました。景色が最高でした！'
      }
      this.showNewDiary = true
    },
    closeNewDiary() {
      this.showNewDiary = false
    },
    reserveDiary() {
      // 予約処理のロジックをここに実装
      this.showSnackbar('予約が完了しました！')
    },
    showSnackbar(text) {
      this.snackbar.text = text
      this.snackbar.show = true
    }
  }
}
</script>

<style>
</style>
