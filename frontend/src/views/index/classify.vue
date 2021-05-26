<template>
  <div role="tablist" style="min-height: 64vh">
    <h3 style="text-align: center;">所有分类</h3>
    <classify-comp v-for="item in allClassify" :key="item" :classify="item"></classify-comp>
  </div>
</template>

<script>
import classifyComp from "../../components/classify-comp"
import axios from "axios";

export default {
  data() {
    return {
      allClassify: []
    }
  },
  components: {
    'classify-comp': classifyComp
  },
  created() {
    axios({
      method: "GET",
      url: this.configVal.GetAllClassify
    })
      .then(res => {
        this.allClassify = res.data
      })
  }
}
</script>
