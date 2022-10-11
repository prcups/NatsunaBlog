<template>
  <div>
    <div :title="(ontop ? '⭐ ' : '') + title">
      <div>
        分类：{{ classify }}
        <br>
        {{ time }}
      </div>
      <button onclick="window.location.href='post/'+id">点击阅读</button>
    </div>
  </div>
</template>

<script>
export default {
  name: "blog-post",
  props: ['title', 'time', 'id', 'classify', 'ontop']
}
</script>