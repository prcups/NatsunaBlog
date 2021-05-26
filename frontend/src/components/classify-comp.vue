<template>
  <b-card no-body class="mb-1">
    <b-card-header header-tag="header" class="p-1" role="tab">
      <b-button block variant="info" @click="show">{{this.classify}}</b-button>
    </b-card-header>
    <b-collapse :visible="visiable" @hidden="hidden" accordion="my-accordion" role="tabpanel">
      <b-card-body>
        <b-list-group>
          <b-list-group-item v-for="item in posts" :key="item.id" :href="'/post/' + item.id">{{item.title}}</b-list-group-item>
        </b-list-group>
      </b-card-body>
    </b-collapse>
  </b-card>
</template>

<script>
import axios from "axios";

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
        axios({
          method: "GET",
          url: this.configVal.GetPostsOfClassify,
          params: {
            classify: this.classify
          }
        })
            .then(res => {
              this.posts = res.data
            })
      }
    },
    hidden() {
      this.visiable = false
    }
  }
}
</script>

<style scoped>

</style>
