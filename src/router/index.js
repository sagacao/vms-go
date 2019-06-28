import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'             the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRouterMap
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRouterMap = [
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },

  {
    path: '/404',
    component: () => import('@/views/404'),
    hidden: true
  },

  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [{
      path: 'dashboard',
      name: 'Dashboard',
      component: () => import('@/views/dashboard/index'),
      meta: { title: '系统首页', icon: 'dashboard' }
    }]
  }]

export const asyncRouterMap = [
  {
    path: '/example',
    component: Layout,
    redirect: '/example/table',
    name: 'Example',
    hidden: true,
    meta: { title: 'Example', icon: 'example' },
    children: [
      {
        path: 'table',
        name: 'Table',
        component: () => import('@/views/table/index'),
        meta: { title: 'Table', icon: 'table' }
      },
      {
        path: 'tree',
        name: 'Tree',
        component: () => import('@/views/tree/index'),
        meta: { title: 'Tree', icon: 'tree' }
      }
    ]
  },

  {
    path: '/form',
    component: Layout,
    hidden: true,
    children: [
      {
        path: 'index',
        name: 'Form',
        component: () => import('@/views/form/index'),
        meta: { title: 'Form', icon: 'form' }
      }
    ]
  },

  {
    path: '/stats',
    component: Layout,
    name: 'stats',
    meta: {
      title: '游戏数据',
      icon: 'nested'
    },
    children: [
      {
        path: 'stats',
        component: () => import('@/views/stats/stats'), // Parent router-view
        name: 'stats',
        meta: { title: '数据管理', roles: ['admin', 'saga'] }
      },
      {
        path: 'statscut',
        component: () => import('@/views/stats/statscut'),
        name: 'statscut',
        meta: { title: '数据查询', roles: ['guest'] }
      }
    ]
  },

  {
    path: '/svrmgr',
    component: Layout,
    name: 'svrmgr',
    meta: {
      title: '游戏配置',
      icon: 'nested',
      roles: ['admin', 'saga'] 
    },
    children: [
      {
        path: 'share',
        component: () => import('@/views/svrmgr/share'), // Parent router-view
        name: 'share',
        meta: { title: '分享设置', roles: ['admin', 'saga'] }
      },
      {
        path: 'switch',
        component: () => import('@/views/svrmgr/switch'), // Parent router-view
        name: 'switch',
        meta: { title: '功能开关', roles: ['admin', 'saga'] }
      },
      {
        path: 'egg',
        component: () => import('@/views/svrmgr/egg'), // Parent router-view
        name: 'egg',
        meta: { title: '砸蛋控制', roles: ['admin', 'saga'] }
      },
      {
        path: 'gates',
        component: () => import('@/views/svrmgr/gates'), // Parent router-view
        name: 'gates',
        meta: { title: '网关配置', roles: ['admin', 'saga'] }
      },
      {
        path: 'appctor',
        component: () => import('@/views/svrmgr/appctor'), // Parent router-view
        name: 'appctor',
        meta: { title: 'APP配置', roles: ['admin', 'saga'] }
      },
      {
        path: 'cstatus',
        component: () => import('@/views/svrmgr/cstatus'), // Parent router-view
        name: 'cstatus',
        meta: { title: '渠道开关', roles: ['admin', 'saga'] }
      },
      {
        path: 'gifts',
        component: () => import('@/views/svrmgr/gifts'), // Parent router-view
        name: 'gifts',
        meta: { title: '礼包码', roles: ['admin', 'saga'] }
      }
    ]
  },

  {
    path: '/payment',
    component: Layout,
    name: 'payment',
    meta: {
      title: '付费数据',
      icon: 'nested',
      roles: ['admin', 'saga'] 
    },
    children: [
      {
        path: 'payment',
        component: () => import('@/views/payment/payment'), // Parent router-view
        name: 'payment',
        meta: { title: '付费数据', roles: ['admin', 'saga'] }
      },
      {
        path: 'paysum',
        component: () => import('@/views/payment/paysum'), // Parent router-view
        name: 'paysum',
        meta: { title: '付费统计', roles: ['admin', 'saga'] }
      },
      {
        path: 'payplayersum',
        component: () => import('@/views/payment/payplayersum'), // Parent router-view
        name: 'payplayersum',
        meta: { title: '付费人数', roles: ['admin', 'saga'] }
      }
    ]
  },
  {
    path: '/game',
    component: Layout,
    name: 'game',
    meta: {
      title: '游戏数据',
      icon: 'nested',
      roles: ['admin', 'saga'] 
    },
    children: [
      {
        path: 'user',
        component: () => import('@/views/game/user'), // Parent router-view
        name: 'user',
        meta: { title: '游戏数据', roles: ['admin', 'saga'] }
      }
    ]
  },
  // 404 page must be placed at the end !!!
  { path: '*', redirect: '/404', hidden: true }
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRouterMap
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
