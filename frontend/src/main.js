import Vue from 'vue'
import App from './App.vue'
import router from './router'
import BootstrapVue from 'bootstrap-vue'
import Axios from 'axios'
import VueAxios from 'vue-axios'
import md5 from 'js-md5'
import Global from "../global"
import Meta from 'vue-meta'

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import 'mavon-editor/dist/css/index.css'

Vue.use(BootstrapVue)
Vue.use(VueAxios, Axios)
Vue.use(Meta)

Vue.config.productionTip = false
Vue.prototype.$md5 = md5
Vue.prototype.configVal = Global

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
