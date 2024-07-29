<template>
  <div>
    <header>
      <nav>
        <NuxtLink id="title" to="/">
          <p>{{ config.Title }}</p>
        </NuxtLink>
        <ul>
          <li><NuxtLink to="/">主页</NuxtLink></li>
          <li><NuxtLink to="/classification">分类</NuxtLink></li>
          <li><NuxtLink to="/timeline">归档</NuxtLink></li>
          <li v-for="item in config.Nav">
            <NuxtLink :to="item.url">
              {{ item.title }}
            </NuxtLink>
          </li>
        </ul>
      </nav>
      <p id="description">{{ config.Description }}</p>
    </header>
    <NuxtPage :key="$route.fullPath" />
    <footer>
      <p>2022 - 由 <a href="https://github.com/prcups/NatsunaBlog/">NatsunaBlog</a> 驱动</p>
    </footer>
  </div>
</template>

<script setup>
const config = useRuntimeConfig().public

</script>

<style>

body {
  background: #5a544b;
  color: white;
  margin: 0;
}

header {
  background-image: url(@/assets/image/bg.avif);
  background-size: cover;
  background-position: 50% 60%;
  display: flex;
  flex-wrap: wrap;
  height: 100vh;
  justify-content: center;
}

nav {
  width: 100%;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  position: fixed;
}

#title {
  text-decoration-line: none;
  font-size: 1.2rem;
  margin-left: 5%;
  color: inherit;
}

nav ul {
  margin-right: 5%;
  font-size: 1.2rem;
  color: inherit;
}

nav ul li{
  display: inline-block;
  margin-left: 2rem;
  color: inherit;
}

nav ul li a {
  text-decoration-line: none;
  color: inherit;
}

#description {
  text-align: center;
  font-size: 4rem;
  color: white;
  align-self: center;
}

@media screen and (max-aspect-ratio: 3/4) {
  header {
    background-image: url(@/assets/image/bg-phone.avif);
  }
  #description {
    font-size: 2.5rem;
  }
}

footer {
  margin-top: 1rem;
  text-align: center;
  padding-top: 1rem;
  padding-bottom: 1rem;
  background-color: #1f3134;
  color: white;
  font-size: 1rem;
}

footer a {
  color: white;
}

</style>
