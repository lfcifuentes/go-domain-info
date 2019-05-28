import Vue from 'vue'
import Router from 'vue-router'
import SearcInfo from './components/SearchInfo.vue'
import Info from './components/Info.vue'

Vue.use(Router)

const router = new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: SearcInfo
    },
    {
      path: '/data',
      name: 'data',
      component: Info
    }
  ]
})

export default router