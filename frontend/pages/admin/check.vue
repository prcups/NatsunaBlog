<template>
  <b-row @submit.stop.prevent>
    <b-col cols="3" offset="4">
      <br>
      <b-form style="margin-top: 20vh">
        <label style="text-align: center">登录</label>
        <b-form-input placeholder="用户名" v-model="username"></b-form-input>
        <b-form-input type="password" placeholder="密码" v-model="password"></b-form-input>
        <b-button variant="primary" style="align-self: center; margin-top: 1vh;" @click="PostPassword">登录</b-button>
        <b-form-text style="margin-top: 1vh">{{ alert }}</b-form-text>
      </b-form>
    </b-col>
  </b-row>
</template>

<script>
import axios from "axios"

export default {
  name: "check",
  data() {
    return {
      username: "",
      password: "",
      alert: ""
    }
  },
  methods: {
    PostPassword() {
      console.log("username: " + this.username)
      console.log("password(after encryption): " + this.$md5(this.password))
      console.log("You can add it to SQL to login.")
      axios({
        method: "post",
        url: this.configVal.CheckUrl,
        data: {
          username: this.username,
          password: this.$md5(this.password),
        }
      })
          .then(res => {
            if (res.data.isChecked == true) {
              this.$router.go(-1)
            } else {
              this.alert = "用户名或密码错误"
            }
          })
    }
  }
}
</script>

<style scoped>

</style>