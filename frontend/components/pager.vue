current<template>
  <ul>
    <li v-if="current - 1 > 1"><NuxtLink :to="base + '?page=1'">1</NuxtLink></li>
    <li v-if="current - 1 > 2">...</li>
    <li v-if="current - 1 >= 1"><NuxtLink :to="base + '?page=' + (current - 1)">{{ current - 1 }}</NuxtLink></li>
    <li class="crtsty"><NuxtLink :to="base + '?page=' + current">{{ current }}</NuxtLink></li>
    <li v-if="current + 1 <= total"><NuxtLink :to="base + '?page=' + (current + 1)">{{ current + 1 }}</NuxtLink></li>
    <li v-if="current + 1 < total - 1">...</li>
    <li v-if="current + 1 < total"><NuxtLink :to="base + '?page=' + total">{{ total }}</NuxtLink></li>
  </ul>
</template>

<script setup>
  const props = defineProps({
    'current': Number,
    'total': Number,
    'base': String
  })

</script>

<style scoped>

ul {
  text-align: center;
}

li {
  display: inline-flex;
  margin: 0 0.5rem;
  background-color: #c7b370;
  height: 3rem;
  width: 3rem;
  align-items: center;
  justify-content: center;
}

li a {
  text-decoration: none;
  color: white;
}

li.crtsty {
  border: white 1px solid;
}

</style>