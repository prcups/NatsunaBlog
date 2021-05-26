<template>
  <div role="tablist" style="min-height: 64vh">
    <h3 style="text-align: center;">所有归档</h3>
    <timeline-comp v-for="item in allTimeline" :key="item" :timeline="item"></timeline-comp>
  </div>
</template>

<script>
import timelineComp from "../../components/timeline-comp"
import axios from "axios";

export default {
  data() {
    return {
      allTimeline: []
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
        })
  }
}
</script>
