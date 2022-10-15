<template>
  <main>
    <div v-if="posts.length === 0">
      <br>
      <p>你还未发表过文章呢！</p>
    </div>
    <div id="container" v-else>
      <blog-post v-for="item in posts" :key="item.id" v-bind="item"></blog-post>
    </div>
    <br>
  </main>
</template>

<script setup>
  let posts = []
  let pages = 1
  const config = useRuntimeConfig()

  function linkGen(pageNum) {
    return pageNum === 1 ? '?' : '?page=' + pageNum
  }
  function curPage() {
    let queryPage = useRoute().query.page
    return queryPage ? queryPage : 1
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
    pages = (res > 0 ? res : 1)
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
  }).then(res => {
    posts = res
  })

</script>

<style>
  #container {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    grid-gap: 1rem;
  }
</style>