<template>
  <div>
    <p>{{ classify }}</p>
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
  let posts = []
  let config = useRuntimeConfig()
  const props = defineProps([
    'classify'
  ])

  await $fetch(config.GetPostsOfClassify, {
    method: "get",
    params: {
      classify: props.classify
    },
    key: "classify" + props.classify
  }).then(res => {
    posts = res
  })
</script>
