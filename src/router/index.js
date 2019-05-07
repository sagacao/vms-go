/* eslint-disable */
import Vue from 'vue'
import Router from 'vue-router'
import layout from '@/components/layout.vue'

const _import = require('./_import_' + process.env.NODE_ENV)

Vue.use(Router)

export const constantRouterMap = [
  {
    path: '/login',
    component: resolve => require(['../components/login/index.vue'], resolve),
    hidden: true
  },
  {
    path: '/404',
    component: resolve => require(['../components/error/404.vue'], resolve),
    meta: { title: '404' },
    hidden: true
  },
  {
    path: '/403',
    component: resolve => require(['../components/error/403.vue'], resolve),
    meta: { title: '403' },
    hidden: true
  }
]

export const asyncRouterMap = [
  {
    path: '/',
    redirect: '/introduction',
    name: 'home',
    icon: 'pie-graph',
    component: layout,
    meta: { title: '系统首页', icon: 'speedometer'},
    children: [ 
      {path: '/introduction', name: 'introduction', component: _import('introduction'), meta: { title: '系统首页', icon: 'speedometer' }}
    ]
  },
  {
    path: '/pages',
    name: 'stats',
    icon: 'ios-paper',
    component: layout,
    alwaysShow: true,
    meta: { title: '统计数据', icon: 'ios-paper'},
    children: [ 
      {path: '/logger', name: 'stats', component: _import('logger'), meta: { title: '数据管理', icon: 'ios-paper', role: ['admin', 'saga']} },
      {path: '/loggercut', name: 'stats', component: _import('loggercut'), meta: { title: '数据查询', icon: 'ios-paper', role: ['taptap']} }
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
