<template>
  <div>
    <div id="admin-edit-page" v-if="isChecked">
      <input id="admin-edit-title" v-model="post.title" placeholder="标题" />
      <div id="admin-edit-container">
        <md-editor
            id="admin-edit-area"
            @on-upload-img="onUploadImg"
            v-model="post.content"
            :toolbarsExclude="['save']"
        />
      </div>
      <p>文章设置</p>
      <div>
        分类
        <input v-model="post.classify" />
      </div>
      <div>
        标签
        <input v-model="post.tag" />
      </div>
      <div>
        <input type="checkbox" v-model="post.ontop" />
          置顶
        <input type="checkbox" v-model="post.hid" />
          隐藏
      </div>
      <button @click="summit">发布</button>
    </div>
    <div v-else>
      <p>请先登录</p>
    </div>
  </div>
</template>

<script setup>
let config = useRuntimeConfig().public
let user = ""
let isChecked = false
const route = useRoute()
const router = useRouter()

let post = reactive({
  title: "",
  content: "",
  classify: "",
  tag: "",
  ontop: 0,
  hid: 0
})

import MdEditor from "md-editor-v3"
import "md-editor-v3/lib/style.css"

const onUploadImg = async (files, callback) => {
  const res = await Promise.all(
      files.map((file) => {
        return new Promise((rev, rej) => {
          const form = new FormData();
          form.append('image', file);
          $fetch(config.UploadImageUrl, {
            method: 'POST',
            body: form,
            credentials: 'include'
          }).then((res) => rev(res))
            .catch((error) => rej(error));
        });
      })
  );
  callback(res.map((item) => config.GetImageBaseUrl + "/" + item.name));
};

async function summit() {
  await $fetch(config.UpdatePostUrl, {
    method: "post",
    body: {
      id: route.params.id,
      title: post.title,
      content: post.content,
      classify: post.classify,
      tag: post.tag,
      ontop: post.ontop,
      hid: post.hid
    },
    credentials: 'include'
  }).then(res => {
    router.push("/post/" + parseInt(res))
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
    if (route.params.id === -1) {
      return
    }
    return $fetch(config.GetOnePostUrl, {
      method: "get",
      params: {
        id: route.params.id,
        visitOnly: false
      },
      key: "modifyPost" + route.params.id
    })
  }).then(res => {
    if (route.params.id === "-1") return
    if (res === "") {
      useRouter().push("/404")
    }
    if (res.ontop == 1) res.ontop = true
    else res.ontop = false
    if (res.hid == 1) res.hid = true
    else res.hid = false
    post = reactive(res)
  })
}
</script>

<style>

#admin-edit-page {
  display: flex;
  flex-direction: column;
  text-align: center;
  align-items: center;
}

#admin-edit-title {
  font-size: 1.2rem;
  text-align: center;
}

#admin-edit-container {
  margin: 1rem 10rem 1rem 10rem;
}

#admin-edit-area {
  min-height: 64vh;
  text-align: left;
}

#admin-edit-page > button {
  margin: 10px;
}
</style>
