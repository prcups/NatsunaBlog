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
  name: "index",
  data() {
    return {}
  }
}
</script>

<style scoped>

</style>
