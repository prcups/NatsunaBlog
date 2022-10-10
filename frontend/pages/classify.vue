<template>
  <b-container>
    <div v-if="loading" align="center" style="height: 56vh;">
      <b-spinner style="margin-top: 20vh;" variant="primary" label="Spinning"></b-spinner>
      <p>少女祈祷中</p>
    </div>
    <div v-else role="tablist" style="min-height: 56vh">
        <h3 style="text-align: center;">所有分类</h3>
        <classify-comp v-for="item in allClassify" :key="item" :classify="item"></classify-comp>
    </div>
  </b-container>
</template>

<script>
import classifyComp from "../components/classify-comp"
import axios from "axios";

export default {
  data() {
    return {
      allClassify: [],
      loading: true
    }
  },
  metaInfo() {
    return {
      title: "分类" + " - " + this.configVal.Title
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
        this.loading = false
      })
  }
}
</script>
