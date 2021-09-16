import Vue from 'vue'
import VueRouter from 'vue-router'
import { landingRouter } from './landing.route'
import { authRouter } from './auth.route'
import { fallbackRouter } from './404.route'
import { mainRouter } from './main.route'

Vue.use(VueRouter)

const routes = [
  landingRouter,
  authRouter,
  fallbackRouter,
  mainRouter
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router