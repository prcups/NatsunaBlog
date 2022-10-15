<template>
  <div>
    <div id="admin-manage" v-if="isChecked">
      <blog-manage v-for="item in posts" :key="item.id" v-bind="item"></blog-manage>
    </div>
    <div v-else>
      <p>请先登录</p>
    </div>
  </div>
</template>

<script setup>

let isChecked = false
let posts = []
let pages = 1
let config = useRuntimeConfig()
let route = useRoute()

function linkGen(pageNum) {
  return pageNum === 1 ? '?' : '?page=' + pageNum
}
function curPage() {
  let queryPage = route.query.page
  return queryPage ? queryPage : 1
}

if (process.client) {
  await $fetch(config.CheckUrl, {
    method: "post",
    key: "check",
    credentials: 'include'
  }).then(res => {
    isChecked = res.isChecked
    return $fetch(config.GetPageNumUrl, {
      method: 'get',
      params: {
        isAll: true
      },
      key: "adminPageNum"
    })
  }).then(res => {
    pages = (res > 0 ? res : 1)
    if (curPage() > pages) {
      useRouter().push("/?page=" + pages)
    }
    return $fetch(config.GetPostsUrl, {
      method: 'get',
      params: {
        page: curPage(),
        isAll: true
      },
      key: "adminPage" + curPage()
    })
  }).then(res => {
    posts = res
  })
}

</script>

<style>
  #admin-manage {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(256px, 1fr));
    grid-auto-rows: minmax(100px, auto);
    grid-gap: 20px;
  }
</style>