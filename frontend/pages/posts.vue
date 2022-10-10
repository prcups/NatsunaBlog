<template>
  <main>
    <div v-if="loading">
      <p>少女祈祷中</p>
    </div>
    <div v-else-if="this.posts.length === 0" style="height: 56vh">
      <br>
      <p>你还未发表过文章呢！</p>
    </div>
    <div v-else>
      <div>
        <blog-post v-for="item in posts" :key="item.id" v-bind="item"></blog-post>
      </div>
    </div>
    <br>
    <pager></pager>
  </main>
</template>

<script>
import BlogPost from "../components/blog-post"

export default {
  head: {
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
  },
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
