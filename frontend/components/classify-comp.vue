<template>
  <div class="classification-root">
    <div class="classification-box">
      <h2>{{ classify }}</h2>
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


<style>
.classification-root {
  display: flex;
  flex-direction: row;
  justify-content: center;
}

.classification-box {
  width: 100%;
  text-align: center;
  margin-top: 1rem;
  padding: 1rem 25% 2rem 25%;
  background-color: #c7b370;
}

</style>