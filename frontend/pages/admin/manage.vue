<template>
  <div>
    <div v-if="isChecked">
      <h3>所有文章</h3>
      <br>
      <div style="min-height: 72vh">
        <div>
          <blog-manage v-for="item in posts" v-bind="item" :key="item.id"></blog-manage>
        </div>
      </div>
    </div>
    <div v-else>
      <p>请先登录</p>
    </div>
  </div>
</template>

<script>
import {useFetch} from "nuxt/app"

export default {
  data() {
    return {
      isChecked: undefined,
      pages: 1,
      posts: []
    }
  },
  created() {
    setIfChecked(this)
    useFetch(this.$config.GetPageNumUrl, {
      method: 'get',
      params: {
        isAll: true
      }
    }).then(res => {
      this.pages = res.data ? res.data : 1
      if (this.curPage() > this.pages) {
        this.$router.push("/admin/manage?page=" + this.pages)
      }
      return useFetch(this.$config.GetPostUrl, {
        method: 'get',
        params: {
          page: this.curPage(),
          isAll: true
        }
      })
    }).then(res => {
      this.posts = res.data
    })
  },
  methods: {
    linkGen(pageNum) {
      return pageNum === 1 ? '?' : `?page=${pageNum}`
    },
    curPage() {
      let queryPage = this.$route.query.page
      return queryPage ? queryPage : 1
    }
  },
  watch: {
    '$route'() {
      this.$router.go(0)
    }
  }
}
</script>