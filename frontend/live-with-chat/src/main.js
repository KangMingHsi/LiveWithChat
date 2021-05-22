import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import router from './router'
import store from './store'
import api from './apis'
import 'bulma/css/bulma.css'
// import VideoPlayer from 'vue-video-player'
// import 'videojs-contrib-hls'

// const hls = require('videojs-contrib-hls')
// Vue.use(hls)
// require('video.js/dist/video-js.css')
// require('vue-video-player/src/custom-theme.css')
// Vue.use(VideoPlayer)

Vue.config.productionTip = false
Vue.prototype.$api = api

new Vue({
  vuetify,
  router,
  store,
  render: h => h(App)
}).$mount('#app')
