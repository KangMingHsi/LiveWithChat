import Vue from 'vue'
import Axios from 'axios'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import router from './router'
import 'bulma/css/bulma.css'

Vue.config.productionTip = false
Vue.prototype.$http = Axios

new Vue({
  vuetify,
  router,
  render: h => h(App)
}).$mount('#app')
