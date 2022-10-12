<template>
  <div>
    <header>
      <NuxtLink to="/">
        <h1>{{ config.Title }}</h1>
      </NuxtLink>
      <h3>{{ config.Description }}</h3>
      <nav>
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
    </header>
    <NuxtPage/>
    <footer>
      <div>
        <hr>
        <p>2022 - 由 <a href="https://github.com/prcups/NatsunaBlog/">NatsunaBlog</a> 驱动</p>
      </div>
    </footer>
  </div>
</template>

<script setup>
const config = useRuntimeConfig()

</script>