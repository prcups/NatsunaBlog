<template>
  <div>
    <div>
      <h3>所有归档</h3>
      <timeline-comp v-for="item in allTimeline" :timeline="item"></timeline-comp>
    </div>
  </div>
</template>

<script setup>
let config = useRuntimeConfig()
let allTimeline = []

useHead({
  title: "归档" + " - " + config.Title
})

await useFetch(config.GetAllTimeline, {
  key: "allTimeline"
}).then(res => {
  allTimeline = res.data
})

</script>
