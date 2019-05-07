// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import store from './store'
import ElementUI from 'element-ui'
import i18n from './utils/lang' // internationalization
import 'element-ui/lib/theme-chalk/index.css' // 默认主题
// import '../static/css/theme-green/index.css';       // 浅绿色主题
import './assets/css/icon.css'
import './components/common/directives'
// import '@/icons' // icon

Vue.config.productionTip = false
Vue.use(ElementUI, {
  size: 'small',
  i18n: (key, value) => i18n.t(key, value)
})

// 使用钩子函数对路由进行权限跳转
router.beforeEach((to, from, next) => {
  // console.log(to, store.getters.token)
  if (store.getters.token) { // 判断是否有token
    if (to.path === '/login') {
      next({ path: '/' })
    } else {
      const user = store.getters.name
      const hasRoles = store.getters.roles && store.getters.roles.length > 0
      console.log(store.getters.roles)
      if (hasRoles) {
        // store.dispatch('getNowRoutes', to)
        next()
      } else {
        store.dispatch('GetInfo', user).then(res => { // 拉取user_info
          const roles = res.data.role
          // console.log(roles)
          store.dispatch('GenerateRoutes', { roles }).then(() => { // 生成可访问的路由表
            router.addRoutes(store.getters.addRouters) // 动态添加可访问路由表
            // console.log(store.getters.permission_routers)
            next({ ...to, replace: true }) // hack方法 确保addRoutes已完成
          })
        }).catch((error) => {
          store.dispatch('FedLogOut', error).then(() => {
            next({ path: '/login' })
          })
        })
      }
    }
  } else {
    if (to.path === '/login') { // 如果是登录页面路径，就直接next()
      next()
    } else { // 不然就跳转到登录；
      next('/login')
    }
  }
})

router.afterEach(() => {
})

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  i18n,
  components: { App },
  template: '<App/>'
})
