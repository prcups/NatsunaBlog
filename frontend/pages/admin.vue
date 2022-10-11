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
    <p v-if="$router.currentRoute.path === '/admin'">
      欢迎来到NatsunaBlog控制台！<br>
      请先点击右上角登录。<br>
      选择“管理“可编辑或删除您写过的文章。”新建“会创建新文章。<br>
    </p>
    <NuxtPage/>
  </div>
</template>

<script>
import {logout, setIfChecked, toCheck} from "@/assets/javascript/check"

export default {
  name: "admin",
  data() {
    return {
      user: "",
      isChecked: undefined
    }
  },
  methods: {
    buttonAct() {
      if (this.isChecked) {
        logout(this)
      } else {
        toCheck(this)
      }
    }
  },
  created() {
    setIfChecked(this)
  },
  computed: {
    buttonWord: function () {
      return this.isChecked ? "退出" : "登录"
    }
  },
  watch: {
    '$route'(newVal, oldVal) {
      if (oldVal.name === "Check") {
        this.$router.go(0)
      }
    }
  }
}
</script>