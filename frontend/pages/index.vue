<template>
  <div>
    <header>
      <NuxtLink to="/">
        <h1>{{ process.env.TITLE }}</h1>
      </NuxtLink>
      <h3>{{ process.env.DESC }}</h3>
      <div>
        <ul>
          <li>主页</li>
          <li>分类</li>
          <li>归档</li>
          <li v-for="item in process.env.NAV" :key="item.title">{{item.title}}</li>
        </ul>
      </div>
    </header>
    <NuxtPage></NuxtPage>
    <footer>
      <div>
        <hr>
        <p>2022 - 由 <a href="https://gitee.com /vanisignal/NatsunaBlog/">NatsunaBlog</a> 驱动</p>
      </div>
    </footer>
  </div>
</template>

<script>
export default {
  data() {
    return {}
  }
}
</script>

<style scoped>

</style>
