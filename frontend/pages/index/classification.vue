<template>
  <div>
    <div>
      <h3 style="text-align: center;">所有分类</h3>
      <classify-comp v-for="item in allClassify" :key="item" :classify="item"></classify-comp>
    </div>
  </div>
</template>

<script>
export default {
  head() {
    return {
      title: "分类"
    }
  },
  data() {
    return {
      allClassify: []
    }
  },
  created() {
    useFetch(this.$config.GetAllClassifyUrl, {
      method: "GET",
    }).then(res => {
      this.allClassify = res.data
    })
  }
}
</script>
