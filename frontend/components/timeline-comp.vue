<template>
  <div>
    <div >
      <button @click="show">{{ this.timeline }}</button>
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
    'timeline'
  ],
  methods: {
    show() {
      this.visiable = !this.visiable
      if (this.visiable) {
        useFetch(GetPostsOfTimeline, {
          method: "GET",
          params: {
            timeline: this.timeline
          }
        }).then(res => {
          this.posts = res.data
        })
      }
    }
  }
}
</script>