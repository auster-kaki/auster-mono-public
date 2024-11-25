<template>
  <v-container>
    <v-row justify="center" align="center">
      <v-col cols="12" sm="8" md="6">
        <h1 class="text-center mb-6">(デモ用)ユーザー設定</h1>
        <v-select
          v-model="selectedUser"
          :items="users"
          item-text="name"
          item-value="id"
          label="ユーザー切り替え"
          outlined
          @change="changeUser"
        ></v-select>
        <v-row>
          <v-spacer />
          <v-col cols="6">
            <v-btn
              color="primary"
              class="mt-4"
              block
              to="/c/settings/create"
            >
              新規ユーザー作成
            </v-btn>
          </v-col>
        </v-row>
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
import { useUserStore } from '~/store/user'

export default {
  name: 'IndexPage',
  layout: 'mobile',
  data() {
    return {
      valid: false,
      selectedUser: 'csuk4ob6mh8s73e2f2ug',
      users: [],
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
  async mounted() {
    if (this.$route.query.createUser === 'success') {
      this.showSnackbar('ユーザー作成が完了しました！')
    }
    // await this.loadUserSettings()
    const userStore = useUserStore()
    userStore.initializeUser()
    console.log('mounted', userStore.userInfo.id)
    this.selectedUser = userStore.userInfo.id
    await this.fetchUsers()
  },
  methods: {
    fetchUsers() {

      return fetch(`${process.env.BASE_URL}/users`, {
        headers: {
          'Content-Type': 'application/json'
        }
      })
        .then(response => {
          if (!response.ok) {
            throw new Error('Network response was not ok')
          }
          return response.json()
        })
        .then(data => {
          this.users = data.Users.map(user => ({
            id: user.ID,
            name: user.Name,
            gender: user.Gender,
            age: user.Age,
            hobbies: [],
            photo: user.Photo
          }))
        })
        .catch(error => {
          console.error('There was a problem with the fetch operation:', error)
          useUserStore().clearUserData() // とりあえず初期化しておく
          throw error
        })
    },
    changeUser(_id) {
      const userStore = useUserStore()
      userStore.updateUserInfo({ id: this.selectedUser })
    },
    showSnackbar(text) {
      this.snackbar.text = text
      this.snackbar.show = true
    }
  }
}
</script>
