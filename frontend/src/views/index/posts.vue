<template>
  <b-container>
    <div v-if="this.posts.length == 0" style="height: 56vh">
      <br>
      <p align="center">你还未发表过文章哦！</p>
    </div>
    <div v-else style="min-height: 56vh">
      <b-row>
        <blog-post v-for="item in posts" :key="item.id" v-bind="item"></blog-post>
      </b-row>
    </div>
    <br>
    <b-pagination-nav :link-gen="linkGen" :number-of-pages="pages" align="center" pills size="ml"
                      use-router></b-pagination-nav>
  </b-container>
</template>

<script>
import BlogPost from "../../components/blog-post"
import axios from 'axios'

export default {
  name: "index-posts",
  components: {
    'blog-post': BlogPost
  },
  data() {
    return {
      posts: [],
      pages: 1
    }
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
  created() {
    document.title = this.configVal.Title
    axios({
      method: 'get',
      url: this.configVal.GetPageNumUrl,
    })
        .then(res => {
          this.pages = res.data.pages != 0 ? res.data.pages : 1
        })
    axios({
      method: 'get',
      url: this.configVal.GetPostsUrl,
      params: {
        page: this.curPage()
      }
    })
        .then(res => {
          this.posts = res.data
        })
  },
  watch: {
    '$route': function () {
      this.$router.go(0)
    }
  }
}
</script>

<style scoped>
</style>
