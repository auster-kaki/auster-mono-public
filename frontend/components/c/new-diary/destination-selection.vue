<template>
  <v-container>
    <h1 class="text-h4 mb-4">行き先を選択してください</h1>
    <v-row>
      <v-col v-for="destination in destinations" :key="destination.name" cols="12" sm="6">
        <v-card
          @click="handleDestinationClick(destination)">
          <v-img
            :src="require(`@/static/destination/${destination.image}`)" :aspect-ratio="5/3"
            style="position:relative" cover>
            <v-chip
              v-if="!destination.available"
              color="error" outlined style="position:absolute; top: 10px; right: 10px; background: #fafafa !important;">
              <strong>Not Available</strong>
            </v-chip>
          </v-img>
          <v-card-title>{{ destination.name }}</v-card-title>
          <v-card-subtitle v-if="destination.subName">{{ destination.subName }}</v-card-subtitle>
          <v-card-text style="white-space: pre-line;">
            <p>{{ destination.description }}</p>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
    <v-snackbar v-model="showSnackbar" :timeout="3000" color="error">
      デモでは銚子市のみ対応しています
    </v-snackbar>
  </v-container>
</template>

<script>
export default {
  name: 'DestinationSelection',
  data() {
    return {
      destinations: [
        {
          id: 1,
          name: '千葉県銚子市',
          image: 'choshi.jpg',
          description: '銚子市は、銚子港の新鮮な海産物と歴史的な銚子電鉄から見える美しい景観が魅力の地域です。' + '\n' + 'NRI Hackathon 2024では、地元の方と交流するための場所として選ばれました。',
          available: true
        },
        {
          id: 2,
          name: '茨城県鉾田市',
          image: 'hokota.jpg',
          description: '鉾田市は、日本有数の農業地域として知られ、良質な米やメロン、野菜が特産品です。' + '\n' + 'NRI Hackathon 2024では、地元の方と交流するための場所として選ばれました。',
          available: false
        },
        {
          id: 3,
          name: '静岡県掛川市',
          image: 'kakegawa.jpg',
          description: '掛川市は、静岡県を代表するお茶の産地として知られ、特に「掛川茶」はその深い味わいで全国にファンを持っています。' + '\n' + 'NRI Hackathon 2024では、地元の方と交流するための場所として選ばれました。',
          available: false
        },
        {
          id: 4,
          name: '羽田PiO PARK',
          subName: '〜未来都市「ハネダ・オーバードライブ」〜',
          image: 'haneda.jpg',
          description: '羽田PiO PARKは、サイバーパンクの未来都市に迷い込んだかのような、煌びやかな光の海が広がります。天空を突き刺すほど高い高層ビルは、星々と同化しながらまるで宇宙に手を伸ばしているかのよう。ビルの間を縫うように飛び交う空飛ぶ車や無数のドローンが、上空をカラフルな流線で埋め尽くします。\n' +
            'そしてその上空では、巨大な宇宙船が低軌道を漂い、ビルの合間にポートインする姿も珍しくありません。グラフィティの光が踊る壁、未来的なモノレールが音もなく滑る軌道、そしてどこからともなく響くサイバーパンクな電子音楽が、この空間をさらなる異次元へと誘います。\n' +
            '羽田PiO PARKは、地球なのか宇宙なのかもわからなくなるこの景色を前に、訪れた者の脳内CPUを100%フル稼働させる、究極の「迷い込んだ感」体験ができます。' + '\n\n' + 'NRI Hackathon 2024では、最終発表会場として選ばれたとかそうではないとか・・',
          available: false
        }
      ],
      showSnackbar: false
    }
  },
  methods: {
    handleDestinationClick(destination) {
      if (destination.available) {
        this.selectDestination(destination.id)
      } else {
        this.showSnackbar = true
      }
    },
    selectDestination(id) {
      this.$emit('destination-selected', id)
    }
  }
}
</script>

<style scoped>
.selected {
  border: 2px solid #1976D2;
}

.unavailable {
  opacity: 0.6;
  cursor: not-allowed;
}
</style>
