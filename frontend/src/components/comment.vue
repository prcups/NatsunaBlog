<template>
  <div class="comment-box">
    <div id="SOHUCS" :sid="'ddp-blog-'+postId"></div>
  </div>
</template>
<script>
export default {
  name: 'comment_changyan',
  props: {
    postId: {
      type: [String, Number],
      required: true
    }
  },
  mounted(){
    this.$nextTick(()=>{
      const appid = 'cyvtsKqVA';
      const conf = 'prod_78db8d3753a290d67513924c41ea5e4e';
      this.loadJs("https://changyan.sohu.com/upload/changyan.js", function () {
        window.changyan && window.changyan.api.config({ appid: appid, conf: conf })
      });
    })
  },
  methods: {
    loadJs(url, cb){
      try {
        const c = document.getElementsByTagName("head")[0] || document.head || document.documentElement;
        const b = document.createElement("script");
        b.setAttribute("type","text/javascript");
        b.setAttribute("charset","UTF-8");
        b.setAttribute("src", url);
        b.setAttribute("id", 'changyan_pc_js');
        if(window.attachEvent){
          b.onreadystatechange = function(){
            const e = b.readyState;
            if(e === "loaded" || e === "complete"){
              b.onreadystatechange = null;
              cb && cb();
            }
          }
        }else{
          if(cb){
            b.onload = cb
          }
        }
        c.appendChild(b)
      } catch (error) {
        cb && cb();
      }
    }
  },
  beforeDestroy() {
    const removeRep = /^http(s)?:\/\/changyan\./
    const $head = document.getElementsByTagName("head")[0] || document.head || document.documentElement;
    const $script = $head.querySelectorAll('script');
    $script.forEach((item, index) => {
      const src = item.getAttribute('src');
      if(src && removeRep.test(src)){
        $head.removeChild(item)
      }
    });
    for (const key in window) {
      if (/^(changyan(\d)?|cyan)/.test(key)){
        window[key] = undefined;
      }
    }
  }
}
</script>
<style scoped>
.comment-box{
  background: #fff;
  padding: 15px
}
</style>