<template>
  <div>
    <div>
      <h3>所有分类</h3>
      <classify-comp v-for="item in allClassify" :classify="item"></classify-comp>
    </div>
  </div>
</template>

<script setup>
let allClassify = []
const config = useRuntimeConfig()

useHead({
  title: "分类 - " + config.Title
})

await useFetch(config.GetAllClassify, {
  key: "allClassify"
}).then(res => {
  allClassify = res.data
})

</script>