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
        <input type="checkbox" :checked="post.ontop" @input="post.ontop = $event" value="1" />
          置顶
        <input type="checkbox" :checked="post.hid" @input="post.hid = $event" value="1" />
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
let config = useRuntimeConfig()
let user = ""
let isChecked = false

let post = {
  title: "",
  content: "",
  classify: "",
  tag: "",
  ontop: false,
  hid: false
}

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
          }).then((res) => {
            callback(config.GetImageBaseUrl + "/" + res.name);
          })
            .catch((error) => rej(error))
        });
      })
  );
};

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
    if (res === "") {
      useRouter().push("/404")
    }
    post = res
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