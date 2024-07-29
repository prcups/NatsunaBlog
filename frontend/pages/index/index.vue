<template>
  <main>
    <div v-if="posts.length === 0" style="text-align: center;">
      <p>你还未发表过文章呢！</p>
    </div>
    <div id="container" v-else>
      <blog-post v-for="item in posts" :key="item.id" v-bind="item"></blog-post>
    </div>
    <pager v-if="pages > 1" :current="curPage()" :total="pages" base="/"></pager>
  </main>
</template>

<script setup>
  let posts = []
  let pages = 1
  const config = useRuntimeConfig().public
  const route = useRoute()

  function curPage() {
    let queryPage = route.query.page
    return parseInt(queryPage ? queryPage : 1)
  }

  useHead({
    title: config.Title,
    meta: [
      {
        name: "description",
        content: config.Description
      }
    ]
  })

  await $fetch(config.GetPageNumUrl, {
    method: 'get',
    params: {
      isAll: false
    },
    key: "pageNum"
  }).then(res => {
    pages = (parseInt(res) > 0 ? parseInt(res) : 1)
    if (curPage() > pages) {
      useRouter().push("/?page=" + pages)
    }
    return $fetch(config.GetPostsUrl, {
      method: 'get',
      params: {
        page: curPage(),
        isAll: false
      },
      key: "page" + curPage()
    })
  }).then((res) => {
    if (res) posts = res
  })

</script>

<style>
  #container {
    display: flex;
    flex-direction: column;
  }
</style>
