<template>
  <div>
    <div id="admin-edit-page" v-if="isChecked">
      <input id="admin-edit-title" v-model="post.title" placeholder="标题" />
      <div id="admin-edit-area">
        <MdEditor id="admin-editor"
          @onUploadImg="onUploadImg"
          @onInput="onInput"
          v-model="post.content"
          :toolbarsExclude="['save']"
        />
      </div>
      <div id="admin-edit-setting">
        <div>分类：
          <input v-model="post.classify" />
        </div>
        <div>标签：
          <input v-model="post.tag" />
        </div>
        <div>置顶：
          <input type="checkbox" v-model="post.ontop" />
          隐藏：
          <input type="checkbox" v-model="post.hid" />
        </div>
      </div>
      <button id="admin-edit-summit" @click="summit">发布</button>
    </div>
    <div v-else>
      <p>请先登录</p>
    </div>
  </div>
</template>

<script setup>
import { MdEditor } from "md-editor-v3"
import "md-editor-v3/lib/style.css"

import { onMounted, onUnmounted } from 'vue';

const handleBeforeUnload = (event) => {
  event.preventDefault();
};

let config = useRuntimeConfig().public
let user = ""
let isChecked = false
let edited = false
const route = useRoute()
const router = useRouter()

onUnmounted(() => {
  if (edited)
    window.removeEventListener('beforeunload', handleBeforeUnload)
});

let post = reactive({
  title: "",
  content: "",
  classify: "",
  tag: "",
  ontop: 0,
  hid: 0
})

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

const onInput = async () => {
  edited = true
  window.addEventListener('beforeunload', handleBeforeUnload)
}

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
  justify-items: center;
  max-width: 800px;
  margin: auto;
}

#admin-edit-title {
  font-size: 1.2rem;
  text-align: center;
  margin: 1rem;
}

#admin-edit-area {
  text-align: left;
  margin: 0 1rem 1rem 1rem;
}

#admin-editor {
  min-height: 60vh;
  height: auto;
}

#admin-edit-setting {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin: 0 1rem 1rem 1rem;
}

#admin-edit-summit {
  margin: 0 1rem 0 1rem;
}

</style>
