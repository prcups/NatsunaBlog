<template>
  <div>
    <div>
      <header>
        <br>
        <h1>{{ post.title }}</h1>
        <h6>分类：{{ post.classify }} | 标签：{{ post.tag !== "" ? post.tag : "无" }}</h6>
        <h6>由 {{ post.author }} 于 {{ post.time }} 所写，被访问 {{ post.visit_times }} 次</h6>
        <hr>
      </header>
      <p>{{ post.content }}</p>
    </div>
  </div>
</template>

<script setup>
import md5 from "js-md5";

const config = useRuntimeConfig()

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
  if (res == "") {
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