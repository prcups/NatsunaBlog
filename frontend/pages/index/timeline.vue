<template>
  <div>
    <div>
      <h3>所有归档</h3>
      <timeline-comp v-for="item in allTimeline" :key="item" :timeline="item"></timeline-comp>
    </div>
  </div>
</template>

<script>

import {useRuntimeConfig} from "nuxt/app"

export default {

  head() {
    return {
      title: "归档" + " - " + config.Title,
      config: useRuntimeConfig()
    }
  },
  data() {
    return {
      allTimeline: []
    }
  },
  created() {
    useFetch(config.GetAllTimeline, {
      method: "GET"
    }).then(res => {
      this.allTimeline = res.data
    })
  }
}
</script>
