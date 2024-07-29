<template>
  <div>
    <classify-comp v-for="item in allClassify" :classify="item"></classify-comp>
  </div>
</template>

<script setup>
let allClassify = []
const config = useRuntimeConfig().public

useHead({
  title: "分类 - " + config.Title
})

await $fetch(config.GetAllClassify, {
  key: "allClassify"
}).then(res => {
  allClassify = res
})

</script>
