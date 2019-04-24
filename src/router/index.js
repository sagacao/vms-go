import Vue from 'vue'
import Router from 'vue-router'
import home from '@/components/home'

const _import = require('./_import_' + process.env.NODE_ENV)

Vue.use(Router)

export const constantRouterMap = [
  {
    path: '/login',
    component: resolve => require(['../components/login/index.vue'], resolve),
    hidden: true
  },
  {
    path: '/pages',
    redirect: '/404',
    component: {
      render (c) { return c('router-view') }
    },
    meta: { title: 'pages' },
    children: [
      {
        path: '/404',
        component: resolve => require(['../components/error/404.vue'], resolve),
        meta: { title: '404' }
      },
      {
        path: '/403',
        component: resolve => require(['../components/error/403.vue'], resolve),
        meta: { title: '403' }
      }
    ]
  }
]

export const asyncRouterMap = [
  {
    path: '/',
    redirect: '/dashboard',
    name: '首页',
    component: home,
    hidden: false,
    children: [
      {
        path: '/vms/dashboard',
        name: 'dashboard',
        icon: 'speedometer',
        component: _import('dashboard')
      }
    ]
  },

  { path: '*', redirect: '/404', hidden: true }

]

export default new Router({
  mode: 'hash',
  linkActiveClass: 'open active',
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRouterMap
})
