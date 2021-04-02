import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import router from './router'
import store from './store'
import api from './apis'
import 'bulma/css/bulma.css'

Vue.config.productionTip = false
Vue.prototype.$api = api

new Vue({
  vuetify,
  router,
  store,
  render: h => h(App)
}).$mount('#app')
