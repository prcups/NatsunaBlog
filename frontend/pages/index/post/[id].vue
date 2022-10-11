<template>
  <div>
    <div>
      <header>
        <br>
        <h1>{{ post.title }}</h1>
        <h6>分类：{{ post.classify }} | 标签：{{ post.tag != "" ? post.tag : "无" }}</h6>
        <h6>由 {{ post.author }} 于 {{ post.time }} 所写，被访问 {{ post.visit_times }} 次</h6>
        <hr>
      </header>
      <p>{{ post.content }}</p>
      <comment :post-id="this.$md5(this.post.id + this.post.title + this.post.content)"></comment>
    </div>
  </div>
</template>

<script>

export default {
  head() {
    return {
      title: this.post.title + " - " + this.configVal.Title,
      meta: [
        {
          name: "keywords",
          content: this.post.tag
        },
        {
          name: "description",
          content: this.desp
        }
      ]
    }
  },
  data() {
    return {
      loading: 1
    }
  },
  computed: {
    desp: function () {
      return this.post.content.substr(0, 20)
    }
  },
  created() {
    useFetch(this.$config.GetOnePostUrl, {
      method: "get",
      params: {
        id: parseInt(this.$route.params.id),
        visitOnly: true
      }
    }).then(res => {
      if (res.data === "") {
        this.$router.push("/page-not-found")
      }
      this.post = res.data
      this.loading = 0
    })
  },
  watch: {
    '$route'() {
      this.$router.go(0)
    }
  }
}
</script>