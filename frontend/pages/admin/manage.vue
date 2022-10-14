<template>
  <div>
    <div v-if="isChecked">
      <h3>所有文章</h3>
      <br>
      <div>
        <blog-manage v-for="item in posts" :key="item.id" v-bind="item"></blog-manage>
      </div>
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

function linkGen(pageNum) {
  return pageNum === 1 ? '?' : '?page=' + pageNum
}
function curPage() {
  let queryPage = useRoute().query.page
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
    if (res) posts = res
  })
}

</script>