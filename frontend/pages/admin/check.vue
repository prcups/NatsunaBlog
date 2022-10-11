<template>
  <div>
    <form>
      <p>登录</p>
      <input type="text" placeholder="用户名" v-model="username" />
      <input type="password" placeholder="密码" v-model="password" />
        <button @click="PostPassword">登录</button>
        <p>{{ alert }}</p>
    </form>
  </div>
</template>

<script>

import {useFetch} from "nuxt/app";

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
      useFetch(this.$config.CheckUrl, {
        method: "post",
        params: {
          username: this.username,
          password: this.$md5(this.password),
        }
      }).then(res => {
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