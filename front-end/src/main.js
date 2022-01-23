import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import Routes from './routes'
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { faTrashAlt, faEdit } from '@fortawesome/free-solid-svg-icons'

library.add(
  faTrashAlt,
  faEdit
)

import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap/dist/js/bootstrap.min.js'
import './global/styles/variables.css';

Vue.use(VueRouter)
Vue.component('font-awesome-icon', FontAwesomeIcon)


const router = new VueRouter({
  routes: Routes
})

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
  router: router
}).$mount('#app')
