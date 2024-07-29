<template>
  <div class="admin-blog-manage">
    <div class="admin-blog-info">
      <p>{{ title }}</p>
      <p>{{ time }}</p>
    </div>
    <button><a class="admin-edit-link" :href=editUrl>编辑</a></button>
    <button @click="deletePost">删除</button>
  </div>
</template>

<script setup>
const props = defineProps(['title', 'time', 'id'])
let config = useRuntimeConfig().public

const editUrl = computed(() => {
  return '/admin/modify/' + props.id
})

async function deletePost() {
  await $fetch(config.DeletePostUrl, {
    method: 'post',
    body: {
      id: props.id
    },
    credentials: 'include'
  }).then(res => {
    if (res == "OK") {
      useRouter().go(0)
    }
  })
}
</script>

<style>
.admin-blog-manage {
  display: flex;
  flex-flow: row;
  flex-wrap: wrap;
  justify-content: center;
  text-align: center;
  align-items: center;
}
.admin-blog-info {
  width: 100%;
}
.admin-edit-link {
  text-decoration: none;
  color: black;
}
</style>
