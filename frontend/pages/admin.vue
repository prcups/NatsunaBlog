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
          <p>{{ user }}</p>
          <button @click="buttonAct">{{ buttonWord }}</button>
        </div>
      </nav>
    </header>
    <p v-if="useRouter().currentRoute.path === '/admin'">
      欢迎来到NatsunaBlog控制台！<br>
      请先点击右上角登录。<br>
      选择“管理“可编辑或删除您写过的文章。”新建“会创建新文章。<br>
    </p>
    <NuxtPage/>
  </div>
</template>

<script setup>
  import {useFetch} from "nuxt/app";
  let config = useRuntimeConfig()
  let user = ""
  let isChecked = false

  await useFetch(config.CheckUrl, {
    method: "post"
  }).then((res) => {
    user = res.data._value.user
    isChecked = res.data._value.isChecked
  })

  console.log(isChecked)

  async function logout() {
    await useFetch(config.LogOutUrl, {
      method: "get"
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
</script>