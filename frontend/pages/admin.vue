<template>
  <div>
    <header>
      <nav>
        <NuxtLink href="/admin">NatsunaBlog控制台</NuxtLink>
        <div>
          <ul>
            <li>
              <NuxtLink href="/admin/manage">管理</NuxtLink>
            </li>
            <li>
              <NuxtLink href="/admin/modify/-1">新建</NuxtLink>
            </li>
            <li>
              <NuxtLink href="/">返回首页</NuxtLink>
            </li>
          </ul>
        </div>
        <div>
          <p v-if="user">{{ user }}</p>
          <button @click="buttonAct">{{ buttonWord }}</button>
        </div>
      </nav>
    </header>
    <NuxtPage/>
  </div>
</template>

<script setup>
  import {useFetch} from "nuxt/app";
  let config = useRuntimeConfig()
  let user = ""
  let isChecked = false

  async function logout() {
    await useFetch(config.LogOutUrl, {
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

  watch(() => {
    return useRoute().path
  }, () => {
    useRouter().go(0)
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