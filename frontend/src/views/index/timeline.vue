<template>
  <b-container>
    <div v-if="loading" align="center" style="height: 56vh;">
      <b-spinner style="margin-top: 20vh;" variant="primary" label="Spinning"></b-spinner>
      <p>少女祈祷中</p>
    </div>
    <div v-else role="tablist" style="min-height: 64vh">
      <h3 style="text-align: center;">所有归档</h3>
      <timeline-comp v-for="item in allTimeline" :key="item" :timeline="item"></timeline-comp>
    </div>
  </b-container>
</template>

<script>
import timelineComp from "../../components/timeline-comp"
import axios from "axios";

export default {
  data() {
    return {
      allTimeline: [],
      loading: true
    }
  },
  metaInfo() {
    return {
      title: "归档" + " - " + this.configVal.Title
    }
  },
  components: {
    'timeline-comp': timelineComp
  },
  created() {
    axios({
      method: "GET",
      url: this.configVal.GetAllTimeline
    })
        .then(res => {
          this.allTimeline = res.data
          this.loading = false
        })
  }
}
</script>
