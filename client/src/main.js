import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import router from './router'
import { createProvider } from './vue-apollo'
import store from './store'

Vue.config.productionTip = false

new Vue({
  vuetify,
  router,
  apolloProvider: createProvider(),
  store,
  render: h => h(App)
}).$mount('#app')
