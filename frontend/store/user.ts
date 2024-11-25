import { defineStore } from 'pinia'

export const useUserStore = defineStore('user', {
  state: () => ({
    userInfo: null as { id: string } | null
  }),
  actions: {
    initializeUser() {
      console.log(sessionStorage.getItem('userInfo'))
      // sessionStorage にデータがある場合は取得し、なければ初期データを設定
      const storedUserInfo = sessionStorage.getItem('userInfo')
      if (storedUserInfo) {
        this.userInfo = JSON.parse(storedUserInfo)
      } else {
        const testData = {
          id: 'csuk4ob6mh8s73e2f2ug',
        }
        this.userInfo = testData

        sessionStorage.setItem('userInfo', JSON.stringify(testData))
        console.log('初期データを sessionStorage に設定しました:', testData)
      }
    },
    updateUserInfo(newInfo: { id: string }) {
      if (!this.userInfo) {
        console.warn('ユーザー情報が初期化されていません')
        return
      }

      // 更新されたデータをマージ
      this.userInfo = { ...this.userInfo, ...newInfo }

      // sessionStorage にも更新を反映
      sessionStorage.setItem('userInfo', JSON.stringify(this.userInfo))
      console.log('ユーザー情報を更新しました:', this.userInfo)
    },
    clearUserData() {
      // ユーザー情報を削除
      this.userInfo = null
      sessionStorage.removeItem('userInfo')
      console.log('ユーザー情報を削除しました')
    }
  }
})
