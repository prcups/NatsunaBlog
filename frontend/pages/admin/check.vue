<template>
  <div>
    <p>登录</p>
    <input type="text" placeholder="用户名" v-model="username" />
    <input type="password" placeholder="密码" v-model="password" />
    <button @click="postPassword(username, password)">登录</button>
    <p>{{ alert }}</p>
  </div>
</template>

<script setup>
  import md5 from "js-md5";

  let username = ""
  let password = ""
  let alert = ""
  let config = useRuntimeConfig()

  async function postPassword(username, password) {
    await $fetch(config.CheckUrl, {
      method: "post",
      body: JSON.stringify({
        username: username,
        password: md5(password)
      }),
      credentials: 'include'
    }).then(res => {
      if (res.isChecked) {
        useRouter().go(-1)
      } else {
        alert = "用户名或密码错误"
      }
    })
  }
</script>