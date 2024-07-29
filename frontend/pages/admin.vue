<template>
  <div>
    <client-only class="admin">
      <div class="admin-nav">
        <NuxtLink id="admin-title" to="/admin">NatsunaBlog控制台</NuxtLink>
        <ul>
          <li>
            <NuxtLink to="/admin/manage">管理</NuxtLink>
          </li>
          <li>
            <NuxtLink to="/admin/modify/-1">新建</NuxtLink>
          </li>
          <li>
            <NuxtLink to="/">返回首页</NuxtLink>
          </li>
        </ul>
        <button id="toLogin" @click="buttonAct">{{ buttonWord + " " + user }}</button>
      </div>
      <NuxtPage :key="$route.fullPath" />
    </client-only>
  </div>
</template>

<script setup>
  let config = useRuntimeConfig().public
  let user = ""
  let isChecked = false

  async function logout() {
    await $fetch(config.LogOutUrl, {
      method: "get",
      server: false,
      credentials: 'include'
    }).then(res => {
      useRouter().go(0)
    })
  }

  async function toCheck() {
    useRouter().push('/admin/check')
  }

  async function buttonAct() {
    if (isChecked) {
      await logout()
    } else {
      await toCheck()
    }
  }

  let buttonWord = computed(() => {
    return isChecked ? "退出" : "登录"
  })

  if (process.client) {
    await $fetch(config.CheckUrl, {
      method: "post",
      server: false,
      key: "check",
      credentials: 'include'
    }).then(res => {
      isChecked = res.isChecked
      user = res.user
    })
  }

</script>

<style>
  .admin-nav {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .admin-nav ul li {
    display: inline-block;
    margin: 0 1rem;
  }

  .admin-nav ul li a {
    text-decoration: none;
    color: white;
  }

  button {
    border: none;
    padding: 5px;
    text-align: center;
    text-decoration: none;
    display: inline-block;
    cursor: pointer;
  }

  #toLogin {
    margin-right: 1rem;
  }

  #admin-title {
    text-decoration: none;
    color: white;
    margin-left: 1rem;
  }
</style>
