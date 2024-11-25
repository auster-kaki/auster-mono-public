<template>
  <div>
    <v-container fluid class="pa-0 bg-image">
      <div>
        <!-- 中央寄せの文字 -->
        <div class="text-overlay text-center">
          <h1 class="display-2 white--text" style="font-size: 48px; color: white;text-shadow: 3px 4px 4px black;">
            どこかにマイル</h1>
        </div>
      </div>
    </v-container>
    <v-container class="form-container">
      <v-row>
        <v-col cols="12">
          <v-text-field
            label="発着地"
            outlined
            dense
          ></v-text-field>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" md="6">
          <v-text-field
            v-model="departureDate"
            label="行きの日付"
            outlined
            dense
            type="date"
          ></v-text-field>
        </v-col>
        <v-col cols="12" md="6">
          <v-text-field
            v-model="returnDate"
            label="帰りの日付"
            outlined
            dense
            type="date"
          ></v-text-field>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" md="6">
          <v-text-field
            v-model="departureTime"
            label="出発時間"
            outlined
            dense
            type="time"
          ></v-text-field>
        </v-col>
        <v-col cols="12" md="6">
          <v-text-field
            v-model="returnTime"
            label="到着時間"
            outlined
            dense
            type="time"
          ></v-text-field>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12">
          <v-text-field
            label="興味関心"
            outlined
            dense
          ></v-text-field>
        </v-col>
      </v-row>
      <v-row justify="end">
        <v-btn color="primary">検索する</v-btn>
      </v-row>
    </v-container>
    <div class="bottom-space"></div>
  </div>
</template>

<script>
export default {
  layout: 'dokokani',
  data() {
    return {
      departureDate: null,
      departureTime: null,
      returnDate: null,
      returnTime: null
    }
  }
}
</script>

<style scoped>
/* 背景画像の設定 */
.bg-image {
  background-image: url('/plane.jpg'); /* 画像のパスを設定 */
  background-size: cover; /* 画像をコンテナ全体にカバー */
  background-position: center; /* 中央寄せ */
  height: 100vh; /* コンテナの高さを画面全体に設定 */
  max-height: 500px; /* 高さの上限を500pxに設定 */
  position: relative; /* 絶対配置を可能にする */
}

/* 上下左右中央寄せのテキスト */
.text-overlay {
  position: absolute; /* 親要素に対して絶対配置 */
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%); /* 正確に中央寄せ */
}

/* 画面下部のスペース */
.bottom-space {
  height: 50px; /* スペースの高さを調整 */
}
</style>
