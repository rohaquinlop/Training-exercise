import Vue from 'vue'
import VueRouter from 'vue-router'
import SyncData from '../views/SyncData.vue'
import BuyerInfo from '../views/BuyerInfo.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'SyncData',
    component: SyncData
  },
  {
    path: '/buyers',
    name: 'Buyers',
    // route level code-splitting
    // this generates a separate chunk (Buyers.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import('../views/Buyers.vue')
  },
  {
    path: '/buyerinfo',
    name: 'BuyerInfo',
    component: BuyerInfo
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
