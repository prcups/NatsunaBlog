<template>
  <div>
    <p>登录</p>
    <input type="text" placeholder="用户名" v-model="username" />
    <input type="password" placeholder="密码" v-model="password" />
    <button @click="PostPassword">登录</button>
    <p>{{ alert }}</p>
  </div>
</template>

<script setup>
  import md5 from "js-md5";

  let username = ""
  let password = ""
  let alert = ""
  let config = useRuntimeConfig()

  async function PostPassword() {
    await useFetch(config.CheckUrl, {
      method: "post",
      params: {
        username: this.username,
        password: md5(this.password),
      }
    }).then(res => {
      if (res.data.isChecked === true) {
        useRouter.go(-1)
      } else {
        this.alert = "用户名或密码错误"
      }
    })
  }
</script>