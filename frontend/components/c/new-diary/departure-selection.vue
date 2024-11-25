<template>
  <div>
    <v-container class="form-container">
      <v-row>
        <v-col cols="12">
          <v-autocomplete
            v-model="departurePlace"
            :items="departurePlaceOptions"
            label="発着地"
            outlined
            dense
          ></v-autocomplete>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="6" md="6">
          <v-text-field
            v-model="departureDate"
            label="行きの日付"
            outlined
            dense
            type="date"
          ></v-text-field>
        </v-col>
        <v-col cols="6" md="6">
          <v-text-field
            v-model="departureTime"
            label="出発時間"
            outlined
            dense
            type="time"
          ></v-text-field>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="6" md="6">
          <v-text-field
            v-model="returnDate"
            label="帰りの日付"
            outlined
            dense
            type="date"
          ></v-text-field>
        </v-col>
        <v-col cols="6" md="6">
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
          <v-select
            v-model="interests"
            :items="interestOptions"
            item-text="name"
            item-value="id"
            label="趣味"
            outlined
            dense
            :error-messages="interestError"
          ></v-select>
        </v-col>
      </v-row>
      <v-row justify="end">
        <v-btn color="primary" @click="onSubmit">検索する</v-btn>
      </v-row>
    </v-container>
    <div class="bottom-space"></div>
  </div>
</template>

<script>
export default {
  name: 'DepartureSelection',
  props: {
    initialDeparturePlace: {
      type: String,
      default: ''
    },
    initialDepartureDate: {
      type: String,
      default: null
    },
    initialReturnDate: {
      type: String,
      default: null
    },
    initialDepartureTime: {
      type: String,
      default: null
    },
    initialReturnTime: {
      type: String,
      default: null
    },
    initialInterests: {
      type: Array,
      default: () => []
    }
  },
  data() {
    const today = new Date()
    const tomorrow = new Date(today)
    tomorrow.setDate(tomorrow.getDate() + 1)

    return {
      departurePlace: this.initialDeparturePlace || '東京',
      departureDate: this.initialDepartureDate || today.toISOString().substr(0, 10),
      departureTime: this.initialDepartureTime || '12:00',
      returnDate: this.initialReturnDate || tomorrow.toISOString().substr(0, 10),
      returnTime: this.initialReturnTime || '18:00',
      departurePlaceOptions: ['東京', '新宿', '渋谷', '池袋', '上野', '品川', '秋葉原', '銀座'],
      interests: this.initialInterests,
      interestOptions: [
        { id: 'cstkdiat6c3011a83so0', name: '釣り' },
        { id: 'cstkdiat6c3011a83sog', name: 'キャンプ' },
      ],
      interestError: ''
    }
  },
  watch: {
    interests() {
      this.validateInterests()
    }
  },
  methods: {
    onSubmit() {
      if (this.validateInterests()) {
        this.$emit('submit', {
          departurePlace: this.departurePlace,
          departureDate: this.departureDate,
          departureTime: this.departureTime,
          returnDate: this.returnDate,
          returnTime: this.returnTime,
          interests: this.interests
        })
      }
    },
    validateInterests() {
      if (this.interests.length > 0) {
        this.interestError = ''
        return true
      } else {
        this.interestError = '必須です'
        return false
      }
    }
  }
}
</script>

<style scoped>
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
