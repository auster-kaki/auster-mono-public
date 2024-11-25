<script>
import { defineComponent } from 'vue'
import { useUserStore } from '~/store/user'

export default defineComponent({
  name: 'fix',
  mounted() {
    const userStore = useUserStore()
    userStore.clearUserData()
  }
})
</script>

<template>
  <div>なんか色々ダメになった時用、ローカルの保存領域をすべて吹き飛ばす</div>
</template>

<style scoped>

</style>
