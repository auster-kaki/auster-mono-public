<template>
  <v-container>
    <v-row justify="center" align="center">
      <v-col cols="12" sm="4" class="text-center">
        <v-btn class="mt-3" color="primary" block nuxt to="/dev">
          開発用ボタン
        </v-btn>
      </v-col>
      <v-col cols="12" sm="4" class="text-center">
        <v-btn class="mt-3" color="primary" block nuxt to="/c/home">
          旅行者向け画面へ
        </v-btn>
      </v-col>
      <v-col cols="12" sm="4" class="text-center">
        <v-btn class="mt-3" color="primary" block nuxt to="/b/home">
          事業者向け画面へ
        </v-btn>
      </v-col>
    </v-row>
    <v-row justify="center" align="center">
      <v-col cols="12" sm="4" class="text-center">
        <v-btn class="mt-3" color="error" block nuxt to="/fix">
          緊急用、ローカルに保存してあるものを吹き飛ばす
        </v-btn>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: 'IndexPage',
}
</script>
