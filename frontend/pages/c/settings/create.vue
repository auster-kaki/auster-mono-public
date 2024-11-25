<script>
import { useUserStore } from '~/store/user'

export default {
  name: 'create',
  layout: 'mobile',
  data() {
    return {
      form: null,
      valid: false,
      newUser: {
        id: '',
        name: '',
        gender: '',
        age: null,
        hobbies: [],
        photo: null
      },
      hobbyOptions: [
        { id: 'cstkdiat6c3011a83so0', name: '釣り' },
        { id: 'cstkdiat6c3011a83sog', name: 'キャンプ' }
      ]
    }
  },
  methods: {
    addNewUser() {
      if (this.$refs.form.validate()) {
        const formData = new FormData()
        formData.append('id', 'dummy')
        formData.append('name', this.newUser.name)
        formData.append('gender', this.newUser.gender)
        formData.append('age', this.newUser.age)
        formData.append('hobbies', JSON.stringify(this.newUser.hobbies))
        formData.append('photo', this.newUser.photo)
        fetch(`${process.env.BASE_URL}/users`, {
          method: 'POST',
          body: formData
        })
          .then(response => {
            if (!response.ok) {
              throw new Error(`HTTP error! status: ${response.status}`)
            }
            return response.json()
          })
          .then(data => {
            useUserStore().updateUserInfo({ id: data.id })
            this.$router.push({ path: '/c/settings', query: { createUser: 'success' } })
          })
          .catch(error => {
            console.error('Error:', error)
            // エラーハンドリングをここに追加
          })
      }
    }
  }
}
</script>
<template>
  <v-container>
    <v-row justify="center" align="center">
      <v-col cols="12" sm="8" md="6">
        <h1 class="text-center mb-6">新規ユーザー作成</h1>

        <v-form ref="form" v-model="valid" @submit.prevent="addNewUser">
          <v-text-field
            v-model="newUser.name"
            label="ユーザー名"
            required
          ></v-text-field>

          <v-radio-group v-model="newUser.gender" row>
            <v-radio label="男性" value="male"></v-radio>
            <v-radio label="女性" value="female"></v-radio>
            <v-radio label="その他" value="other"></v-radio>
          </v-radio-group>

          <v-text-field
            v-model.number="newUser.age"
            label="年齢"
            type="number"
            required
          ></v-text-field>

          <v-combobox
            v-model="newUser.hobbies"
            :items="hobbyOptions"
            item-text="name"
            label="趣味"
            multiple
            chips
            small-chips
            deletable-chips
          ></v-combobox>

          <v-file-input
            v-model="newUser.photo"
            accept="image/*"
            label="顔写真"
            prepend-icon="mdi-camera"
          ></v-file-input>

          <v-btn
            color="primary"
            class="mt-4"
            block
            type="submit"
          >
            作成
          </v-btn>
        </v-form>
      </v-col>
    </v-row>
  </v-container>
</template>

<style scoped>
</style>
