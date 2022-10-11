<template>
  <div>
    <p>{{ time }}</p>
    <button><a :href=editUrl>编辑</a></button>
    <button @click="deletePost">删除</button>
  </div>
</template>

<script>

import {useFetch} from "nuxt/app";

export default {
  name: "blog-manage",
  props: ['title', 'time', 'id'],
  methods: {
    deletePost() {
      useFetch(DeletePostUrl, {
        method: "post",
        params: {
          id: this.id
        }
      }).then(res => {
            this.$router.go(0)
          })
    }
  },
  computed: {
    editUrl: function () {
      return "/admin/modify/" + this.id
    }
  }
}
</script>