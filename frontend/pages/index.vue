<template>
  <div>
    <header>
      <NuxtLink to="/">
        <h1>{{ config.Title }}</h1>
      </NuxtLink>
      <h3>{{ config.Description }}</h3>
      <div>
        <ul>
          <li>主页</li>
          <li>分类</li>
          <li>归档</li>
          <li v-for="item in config.Nav" :key="item.title">{{ item.title }}</li>
        </ul>
      </div>
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
import {useRuntimeConfig} from "nuxt/app"

const config = useRuntimeConfig()

</script>