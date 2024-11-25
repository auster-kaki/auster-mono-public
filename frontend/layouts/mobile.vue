<template>
  <v-app>
    <v-main>
      <Nuxt />
    </v-main>
    <v-bottom-navigation v-model="value" fixed>
      <navigation-button
        v-for="item in navigationItems"
        :key="item.value"
        :value="item.value"
        :to="item.to"
        :icon="item.icon"
        :label="item.label"
      />
    </v-bottom-navigation>
  </v-app>
</template>

<script>
import NavigationButton from '~/components/c/common/NavicationButton.vue'
import { useUserStore } from '~/store/user'

export default {
  name: 'DefaultLayout',
  components: {
    NavigationButton
  },
  data() {
    return {
      value: 'create',
      username: '', // ユーザー名を保持するためのプロパティを追加
      navigationItems: [
        { value: 'create', to: '/c/new-diary', icon: 'mdi-airplane-takeoff', label: '予約' },
        { value: 'reservations', to: '/c/reservations', icon: 'mdi-calendar-check', label: '予約確認' },
        { value: 'history', to: '/c/histories', icon: 'mdi-book-open-page-variant', label: '日記' },
        { value: 'settings', to: '/c/settings', icon: 'mdi-account', label: 'ユーザー' }
      ]
    }
  },
  // ユーザー名を取得するためのメソッドを追加（実際の実装はアプリケーションの仕様に応じて変更してください）
  mounted() {
    const userStore = useUserStore()
    userStore.initializeUser()
  }
}
</script>
