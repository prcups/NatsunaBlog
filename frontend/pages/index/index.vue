<template>
  <main>
    <div v-if="posts.length === 0">
      <br>
      <p>你还未发表过文章呢！</p>
    </div>
    <div v-else>
      <div>
        <blog-post v-for="item in posts" :key="item.id" v-bind="item"></blog-post>
      </div>
    </div>
    <br>
  </main>
</template>

<script setup>
  let posts = []
  let pages = 1
  const config = useRuntimeConfig()

  function linkGen(pageNum) {
    return pageNum === 1 ? '?' : `?page=${pageNum}`
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

  await useFetch(config.GetPageNumUrl, {
    method: 'get',
    params: {
      isAll: false
    }
  }).then(res => {
    pages = (res.data > 0 ? res.data : 1)
    if (curPage() > pages) {
      useRouter().push("/?page=" + pages)
    }
    return useFetch(config.GetPostsUrl, {
      method: 'get',
      params: {
        page: curPage(),
        isAll: false
      }
    })
  }).then(res => {
    if (res.data) posts = res.data
  })

  watch(useRoute(), () => {
      useRouter().go(0)
  })

</script>