<template>
  <div class="post-root">
    <div class="post-box">
      <h1>{{ post.title }}</h1>
      <p>分类：{{ post.classify }} | 标签：{{ post.tag !== "" ? post.tag : "无" }}</p>
      <p>由 {{ post.author }} 于 {{ post.time }} 所写，被访问 {{ post.visit_times }} 次</p>
      <markdown-it class="post-content" :source="post.content"></markdown-it>
    </div>
  </div>
</template>

<script setup>
import markdownIt from 'vue3-markdown-it'
import 'highlight.js/styles/monokai.css'

defineNuxtComponent(markdownIt)

const config = useRuntimeConfig().public

let post = {
  title: "",
  time: "",
  id: "",
  classify: "",
  tag: "",
  content: ""
}

let desp = computed(() => {
  return String(post.content).substring(0, 20)
})

await $fetch(config.GetOnePostUrl, {
  method: "get",
  params: {
    id: useRoute().params.id,
    visitOnly: true
  },
  key: "post" + useRoute().params.id
}).then(res => {
  if (res === "") {
   useRouter().push("/404")
  }
  post = res

})

useHead({
  title: post.title + " - " + config.Title,
  meta: [
    {
      name: "description",
      content: desp
    }
  ]
})

</script>

<style>
.post-root {
  display: flex;
  flex-direction: row;
  justify-content: center;
}

.post-box {
  width: 100%;
  text-align: center;
  margin-top: 1rem;
  padding: 1rem 25% 2rem 25%;
  background-color: #c7b370;
}

@media screen and (max-width: 1200px) {
  .post-box {
    padding: 1rem 10% 2rem 10%;
  }
}

.post-content {
  text-align: left;
}

.post-content h1 {
  font-size: 1.5rem;
}

.post-content h1 {
  font-size: 1.6rem;
}

.post-content h2 {
  font-size: 1.5rem;
}

.post-content h3 {
  font-size: 1.4rem;
}

.post-content h4 {
  font-size: 1.3rem;
}

.post-content h5 {
  font-size: 1.2rem;
}

.post-content h6 {
  font-size: 1.1rem;
}
</style>
