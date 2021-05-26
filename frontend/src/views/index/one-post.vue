<template>
  <b-container>
    <header class="header">
      <br>
      <h1>{{ post.title }}</h1>
      <h6>分类：{{post.classify}} | 标签：{{post.tag != "" ? post.tag : "无"}}</h6>
      <h6>由 {{ post.author }} 于 {{ post.time }} 所写，被访问 {{ post.visit_times }} 次</h6>
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
        content: ""
      }
    }
  },
  computed: {
    desp: function() {
      return this.post.content.substr(0, 20)
    }
  },
  metaInfo() {
    return {
      title: this.post.title + " - " + this.configVal.Title,
      meta:[
        {
          name: "keywords",
          content:this.post.tag
        },
        {
          name:"description",
          content:this.desp
        }
      ]
    }
  },
  created() {
    axios({
      method: "get",
      url: this.configVal.GetOnePostUrl,
      params: {
        id: parseInt(this.$route.params.id),
        visitOnly: true
      }
    })
        .then(res => {
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