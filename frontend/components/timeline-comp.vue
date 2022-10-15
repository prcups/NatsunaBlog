<template>
  <div>
    <p>{{ timeline }}</p>
    <ul>
      <li v-for="item in posts">
        <NuxtLink :to="'/post/' + item.id">
          {{ item.title }}
        </NuxtLink>
      </li>
    </ul>
  </div>
</template>

<script setup>
let config = useRuntimeConfig()
const props = defineProps(['timeline'])

let posts = []
await $fetch(config.GetPostsOfTimeline, {
  method: "get",
  params: {
    timeline: props.timeline
  },
  key: "timeline" + props.timeline
}).then(res => {
  posts = res
})
</script>
