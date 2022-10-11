<template>
  <div>
    <div>
      <button block variant="info" @click="show">{{ this.classify }}</button>
    </div>
    <div>
      <div>
        <ul>
          <li v-for="item in posts" :key="item.id"><a :href="'/post/' + item.id">
            {{ item.title }}
          </a></li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>

import {useFetch} from "nuxt/app";

export default {
  data() {
    return {
      posts: [],
      visiable: false
    }
  },
  props: [
    'classify'
  ],
  methods: {
    show() {
      this.visiable = !this.visiable
      if (this.visiable) {
        useFetch(GetPostsOfClassify, {
          method: "GET",
          params: {
            classify: this.classify
          }
        }).then(res => {
          this.posts = res.data
        })
      }
    }
  }
}
</script>

<style scoped>

</style>
