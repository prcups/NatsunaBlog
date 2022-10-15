<template>
  <div>
    <timeline-comp v-for="item in allTimeline" :timeline="item"></timeline-comp>
  </div>
</template>

<script setup>
let config = useRuntimeConfig()
let allTimeline = []

useHead({
  title: "归档" + " - " + config.Title
})

await $fetch(config.GetAllTimeline, {
  key: "allTimeline"
}).then(res => {
  allTimeline = res
})

</script>
