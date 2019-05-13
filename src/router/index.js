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
    component: layout,
    meta: { title: '系统首页', icon: 'el-icon-s-home'},
    children: [ 
      {path: '/introduction', name: 'introduction', component: _import('pages/introduction'), meta: { title: '系统首页', icon: 'el-icon-s-home' }}
    ]
  },
  {
    path: '/pages',
    name: 'stats',
    component: layout,
    alwaysShow: true,
    meta: { title: '统计数据', icon: 'el-icon-edit'},
    children: [ 
      {path: '/logger', name: 'statsad', component: _import('pages/logger'), meta: { title: '数据管理', icon: 'el-icon-edit', role: ['admin', 'saga']} },
      {path: '/loggercut', name: 'statscut', component: _import('pages/loggercut'), meta: { title: '数据查询', icon: 'el-icon-view', role: ['taptap', 'dychannel']} }
    ]
  },
  {
    path: '/settings',
    name: 'settings',
    component: layout,
    meta: { title: '游戏管理', icon: 'el-icon-setting', role: ['admin', 'saga']},
    children: [
      {path: '/share', name: 'share', component: _import('settings/share'), meta: { title: '分享设置', icon: 'el-icon-share', role: ['admin', 'saga']} },
      {path: '/switch', name: 'switch', component: _import('settings/switch'), meta: { title: '功能开关', icon: 'el-icon-share', role: ['admin', 'saga']} },
      {path: '/egg', name: 'egg', component: _import('settings/egg'), meta: { title: '每日砸蛋控制', icon: 'el-icon-share', role: ['admin', 'saga']} }
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
