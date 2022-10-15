<template>
  <div class="timeline-root">
    <div class="timeline-box">
      <h1>{{ timeline }}</h1>
      <ul>
        <li v-for="item in posts">
          <NuxtLink :to="'/post/' + item.id">
            {{ item.title }}
          </NuxtLink>
        </li>
      </ul>
    </div>
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

<style>
.timeline-root {
  display: flex;
  flex-direction: row;
  justify-content: center;
}

.timeline-box {
  width: 100%;
  text-align: center;
  margin-top: 1rem;
  padding: 1rem 25% 2rem 25%;
  background-color: #c7b370;
}

</style>