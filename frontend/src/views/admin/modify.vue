<template>
  <div>
    <div v-if="isChecked">
      <div style="margin-top: 1vh">
        <b-form-input v-model="post.title" placeholder="标题"></b-form-input>
        <mavon-editor v-model="post.content" ref="md" style="min-height: 70vh" @imgAdd="$imgAdd"
                      @imgDel="$imgDel"></mavon-editor>
      </div>

      <b-input-group style="margin-top: 1vh;" >
        <b-input-group-prepend>
          <b-input-group-text>文章设置</b-input-group-text>
        </b-input-group-prepend>
        <b-form-input placeholder="分类" v-model="post.classify"></b-form-input>
        <b-form-input placeholder="标签" v-model="post.tag"></b-form-input>
        <b-input-group-append>
          <b-form-checkbox button v-model="post.ontop" :checked="1" :unchecked-value="0">置顶</b-form-checkbox>
          <b-form-checkbox button v-model="post.hid" :checked="1" :unchecked-value="0">隐藏</b-form-checkbox>
        </b-input-group-append>
      </b-input-group>

      <b-button @click="summit" variant="primary" style="align-self: center; margin-top: 1vh;">发布</b-button>
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
        author: "",
        classify: "",
        tag: "",
        ontop: 0,
        hid: 0
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
          content: this.post.content,
          classify: this.post.classify,
          tag: this.post.tag,
          ontop: this.post.ontop,
          hid: this.post.hid
        }
      })
          .then(res => {
            this.$router.push("/post/" + parseInt(res.data))
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