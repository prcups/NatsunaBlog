<template>
  <div>
    <div v-if="isChecked">
      <input v-model="post.title" placeholder="标题" />
      <input type="text" v-model="post.content" />
      <p>文章设置</p>
      <input placeholder="分类" v-model="post.classify" />
      <input placeholder="标签" v-model="post.tag" />

      <input type="checkbox" :checked="post.ontop" @input="post.ontop = $event" value="1" />
        置顶
      <input type="checkbox" :checked="post.hid" @input="post.hid = $event" value="1" />
        隐藏
      <button @click="summit">发布</button>
    </div>
    <div v-else>
      <p>请先登录</p>
    </div>
  </div>
</template>

<script setup>
let post = {
  title: "",
  content: "",
  classify: "",
  tag: "",
  ontop: false,
  hid: false
}
let config = useRuntimeConfig()
let user = ""
let isChecked = false

async function summit() {
  await $fetch(config.UpdatePostUrl, {
    method: "post",
    body: {
      id: useRoute().params.id,
      title: post.title,
      content: post.content,
      classify: post.classify,
      tag: post.tag,
      ontop: post.ontop,
      hid: post.hid
    },
    credentials: 'include'
  }).then(res => {
    useRouter().push("/post/" + parseInt(res))
  })
}

if (process.client) {
  await $fetch(config.CheckUrl, {
    method: "post",
    server: false,
    key: "check",
    credentials: 'include'
  }).then(res => {
    isChecked = res.isChecked
    return $fetch(config.GetOnePostUrl, {
      method: "get",
      params: {
        id: useRoute().params.id,
        visitOnly: false
      },
      key: "modifyPost" + useRoute().params.id
    })
  }).then(res => {
    if (res == "") {
      useRouter().push("/404")
    }
    post = res
  })
}
</script>