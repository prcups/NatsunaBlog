<template>
  <div>
    <div v-if="isChecked">
      <form>
        <input v-model="post.title" placeholder="标题" />
        <input type="text" v-model="post.content" />
        <p>文章设置</p>
        <input placeholder="分类" v-model="post.classify" />
        <input placeholder="标签" v-model="post.tag" />

        <input type="checkbox" :checked="post.ontop" @input="post.ontop = $event" value="1" />
        置顶
        <input type="checkbox" :checked="post.hid" @input="post.hid = $event" value="1" />
        隐藏
        <button @click="summit">发布</button>
      </form>
    </div>
    <div v-else>
      <p>请先登录</p>
    </div>
  </div>
</template>

<script>
import {setIfChecked} from "../../../assets/javascript/check.js"
import {useFetch} from "nuxt/app";

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
  methods: {
    $imgAdd(pos, $file) {
      this.formdata = new FormData();
      this.formdata.append("image", $file);
      useFetch(this.$config.UploadImageUrl, {
        method: 'post',
        data: this.formdata,
        headers: {'Content-Type': 'multipart/form-data'},
      }).then((res) => {
        let imageUrl = GetImageBaseUrl + "/" + res.data.name
        this.$refs.md.$img2Url(pos, imageUrl)
      })
    },
    $imgDel(pos) {
      useFetch(this.$config.DeleteImageUrl, {
        method: 'post',
        data: {
          url: pos[0]
        }
      })
    },
    summit() {
      useFetch(this.$config.UpdatePostUrl, {
        method: "post",
        data: {
          id: this.$route.params.id,
          title: this.post.title,
          content: this.post.content,
          classify: this.post.classify,
          tag: this.post.tag,
          ontop: this.post.ontop,
          hid: this.post.hid
        }
      }).then(res => {
        this.$router.push("/post/" + parseInt(res.data))
      })
    }
  },
  created() {
    setIfChecked(this)
    if (this.$route.params.id != -1) {
      useFetch(this.$config.GetOnePostUrl, {
        method: "get",
        params: {
          id: this.$route.params.id,
          visitOnly: false
        }
      }).then(res => {
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