<template>
  <b-container>
    <header class="header">
      <br>
      <h1>{{ post.title }}</h1>
      <h6>由 {{ post.author }} 于 {{ post.time }} 所写，被访问 {{ post.visitTimes }} 次</h6>
      <hr>
    </header>
    <mavon-editor :value="post.content"
                  :subfield="false"
                  :defaultOpen="'preview'"
                  :toolbarsFlag="false"
                  :editable="false"
                  :scrollStyle="true"
                  :ishljs="true"
                  :boxShadow="false"
    ></mavon-editor>
  </b-container>
</template>

<script>
import axios from 'axios'
import mavonEditor from 'mavon-editor'

export default {
  name: "onepost",
  components: {
    'mavon-editor': mavonEditor.mavonEditor
  },
  data() {
    return {
      post: {
        title: "",
        time: "",
        content: "",
        visitTimes: "",
        author: ""
      }
    }
  },
  created() {
    axios({
      method: "get",
      url: this.configVal.GetOnePostUrl,
      params: {
        id: this.$route.params.id,
        visitOnly: true
      }
    })
        .then(res => {
          console.log(res.data)
          if (res.data == "") {
            this.$router.push("/page-not-found")
          }
          this.post = res.data
        })
  },
  watch: {
    '$route'() {
      this.$router.go(0)
    }
  }
}
</script>

<style scoped>
.header {
  text-align: center;
}
</style>