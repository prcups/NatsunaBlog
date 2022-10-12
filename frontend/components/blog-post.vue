<template>
  <div>
    <div>
      <NuxtLink :to="'/post/' + id">{{ (ontop ? '⭐ ' : '') + title }}</NuxtLink>
      <div>
        分类：{{ classify }}
        <br>
        {{ time }}
      </div>
    </div>
  </div>
</template>

<script setup>
const props = defineProps([
    'title', 'time', 'id', 'classify', 'ontop'
])
</script>