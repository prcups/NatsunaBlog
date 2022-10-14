<template>
  <div>
    <p>{{ title }}</p>
    <p>{{ time }}</p>
    <button><a :href=editUrl>编辑</a></button>
    <button @click="deletePost">删除</button>
  </div>
</template>

<script setup>
const props = defineProps(['title', 'time', 'id'])
let config = useRuntimeConfig()

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
    if (res.isDeleted) {
      useRouter().go(0)
    }
  })
}
</script>