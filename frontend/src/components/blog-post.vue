<template>
  <b-col cols="4" style="min-width: 300px; margin-left: auto; margin-right: auto">
    <b-card :title="(ontop ? '[置顶]' : '') + title">
      <b-card-text>
        分类：{{ classify }}
        <br>
        {{ time }}
      </b-card-text>
      <b-button variant="light" :to="'post/' + id">点击阅读</b-button>
    </b-card>
  </b-col>
</template>

<script>
export default {
  name: "blog-post",
  props: ['title', 'time', 'id', 'classify', 'ontop']
}
</script>

<style scoped>

</style>
