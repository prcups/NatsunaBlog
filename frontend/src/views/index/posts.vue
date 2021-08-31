<template>
  <b-container>
    <div v-if="loading" align="center" style="height: 56vh;">
      <b-spinner style="margin-top: 20vh;" variant="primary" label="Spinning"></b-spinner>
      <p>少女祈祷中</p>
    </div>
    <div v-else-if="this.posts.length == 0" style="height: 56vh">
      <br>
      <p align="center">你还未发表过文章呢！</p>
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
  components: {
    'blog-post': BlogPost
  },
  data() {
    return {
      posts: [],
      pages: 1,
      loading: 1
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
  metaInfo() {
    return {
      title: this.configVal.Title,
      meta:[
        {
          name:'keywords',
          content:this.configVal.Tag
        },
        {
          name:"description",
          content:this.configVal.Description
        }
      ]
    }
  },
  created() {
    axios({
      method: 'get',
      url: this.configVal.GetPageNumUrl,
      params: {
        isAll: false
      }
    })
        .then(res => {
          this.pages = (res.data > 0 ? res.data : 1)
          if (this.curPage() > this.pages) {
            this.$router.push("/?page=" + this.pages)
          }
          axios({
            method: 'get',
            url: this.configVal.GetPostsUrl,
            params: {
              page: this.curPage(),
              isAll: false
            }
        })
          .then(res => {
            if (res.data) this.posts = res.data
            this.loading = 0
          })
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
