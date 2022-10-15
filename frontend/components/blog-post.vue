<template>
  <div class="post">

    <div class="img-container">
      <img :src="'../assets/image/anime/' + (Math.floor(Math.random() * 8) + 1) + '.avif'"  alt=""/>
    </div>

    <div class="post-info">
      <NuxtLink class="post-title" :to="'/post/' + id">{{ (ontop ? '⭐ ' : '') + title }}</NuxtLink>
      <p>{{ classify }}</p>
      <p>{{ time }}</p>
    </div>
  </div>
</template>

<script setup>
const props = defineProps([
    'title', 'time', 'id', 'classify', 'ontop'
])
</script>

<style>
.post {
  margin-top: 1rem;
  padding: 1rem;
  display: flex;
  flex-flow: row;
  flex-wrap: wrap;
  justify-content: center;
  background-color: #c7b370;
}

.img-container {
  width: 24rem;
  margin-left: 1rem;
  margin-right: 1rem;
}

img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.post-info {
  display: flex;
  flex-direction: column;
  flex: 1 15rem;
  text-align: center;
  justify-content: center;
  margin-left: 1rem;
  margin-right: 1rem;
}

.post-title{
  color: white;
  text-decoration: none;
  font-size: 1.5rem;
}

</style>