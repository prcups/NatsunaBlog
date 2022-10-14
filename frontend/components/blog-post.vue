<template>
  <div id="post">
    <NuxtLink :to="'/post/' + id">{{ (ontop ? '⭐ ' : '') + title }}</NuxtLink>
    <div>
      分类：{{ classify }}
      <br>
      {{ time }}
    </div>
  </div>
</template>

<script setup>
const props = defineProps([
    'title', 'time', 'id', 'classify', 'ontop'
])
</script>

<style>
#post {
  margin-top: 1rem;
  border: black solid;
  padding: 1rem;
  text-align: center;
}

</style>