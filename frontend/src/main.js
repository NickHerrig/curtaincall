import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

import App from './App.vue'
import Home from './Home.vue'

const router = new VueRouter({
  mode: 'history',
  routes: [
    { path: '/', component: Home },
  ]
})

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
