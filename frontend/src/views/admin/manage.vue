<template>
  <div>
    <div v-if="isChecked">
      <h3>所有文章</h3>
      <br>
      <div style="min-height: 72vh">
        <b-row>
          <blog-manage v-for="item in posts" v-bind="item" :key="item.id"></blog-manage>
        </b-row>
      </div>
      <b-pagination-nav :link-gen="linkGen" :number-of-pages="pages" align="center" use-router></b-pagination-nav>
    </div>
    <div v-else>
      <p>请先登录</p>
    </div>
  </div>
</template>

<script>
import BlogManage from "../../components/blog-manage"
import {setIfChecked} from "@/assets/javascript/check"
import axios from "axios"

export default {
  name: "manage",
  components: {
    'blog-manage': BlogManage
  },
  data() {
    return {
      isChecked: undefined,
      pages: 1,
      posts: []
    }
  },
  created() {
    setIfChecked(this)
    axios({
      method: 'get',
      url: this.configVal.GetPageNumUrl,
      param: {
        isAll: true
      }
    })
        .then(res => {
          this.pages = res.data != 0 ? res.data : 1
        })
    axios({
      method: 'get',
      url: this.configVal.GetPostsUrl,
      params: {
        page: this.curPage(),
        isAll: true
      }
    })
        .then(res => {
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

<style scoped>
</style>