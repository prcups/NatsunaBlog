<template>
  <b-container>
    <header>
      <b-link href="/">
        <h1 style="color: black; margin-top: 1vh;">{{ this.configVal.Title }}</h1>
      </b-link>
      <h3 align="right">{{ this.configVal.Description }}</h3>
      <b-navbar type="light" variant="light">
        <b-navbar-nav>
          <b-nav-item exact exact-active-class="active" href="/">主页</b-nav-item>
          <b-nav-item exact exact-active-class="active" href="/classify">分类</b-nav-item>
          <b-nav-item exact exact-active-class="active" href="/timeline">归档</b-nav-item>
          <b-nav-item exact v-for="item in this.configVal.OnNav" :href="item.url" :key="item.title">{{item.title}}</b-nav-item>
        </b-navbar-nav>
      </b-navbar>
    </header>
    <router-view></router-view>
    <footer>
      <div style="text-align: center">
        <hr>
        <p>2021 - 由 <a href="https://gitee.com/vanisignal/NatsunaBlog/">NatsunaBlog</a> 驱动</p>
      </div>
    </footer>
  </b-container>
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
