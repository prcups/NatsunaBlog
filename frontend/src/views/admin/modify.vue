<template>
  <div>
    <div v-if="isChecked">
      <div style="min-height: 80vh">
        <br>
        <b-form-input v-model="post.title" placeholder="标题"></b-form-input>
        <mavon-editor v-model="post.content" ref="md" style="min-height: 70vh" @imgAdd="$imgAdd"
                      @imgDel="$imgDel"></mavon-editor>
      </div>
      <b-button @click="summit" variant="primary" style="align-self: center">发布</b-button>
    </div>
    <div v-else>
      <p>请先登录</p>
    </div>
  </div>
</template>

<script>
import axios from "axios"
import {setIfChecked} from "../../assets/javascript/check.js"
import mavonEditor from "mavon-editor"

export default {
  name: "modify",
  data() {
    return {
      user: "Not Login",
      isChecked: undefined,
      post: {
        title: "",
        time: "",
        content: "",
        visit_times: "",
        author: ""
      },
      formdata: undefined
    }
  },
  components: {
    'mavon-editor': mavonEditor.mavonEditor
  },
  methods: {
    $imgAdd(pos, $file) {
      this.formdata = new FormData();
      this.formdata.append("image", $file);
      axios({
        url: this.configVal.UploadImageUrl,
        method: 'post',
        data: this.formdata,
        headers: {'Content-Type': 'multipart/form-data'},
      }).then((res) => {
        let imageUrl = this.configVal.GetImageBaseUrl + "/" + res.data.name
        this.$refs.md.$img2Url(pos, imageUrl)
      })
    },
    $imgDel(pos) {
      axios({
        method: 'post',
        url: this.configVal.DeleteImageUrl,
        data: {
          url: pos[0]
        }
      })
    },
    summit() {
      axios({
        method: "post",
        url: this.configVal.UpdatePostUrl,
        data: {
          id: this.$route.params.id,
          title: this.post.title,
          content: this.post.content
        }
      })
          .then(res => {
            this.$router.push("/post/" + res.data)
          })
    }
  },
  created() {
    setIfChecked(this)
    if (this.$route.params.id != -1) {
      axios({
        method: "get",
        url: this.configVal.GetOnePostUrl,
        params: {
          id: this.$route.params.id,
          visitOnly: false
        }
      })
          .then(res => {
            if (res.data == "") {
              this.$router.push("/page-not-found")
            }
            this.post = res.data
          })
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