/* eslint-disable */
import { login, logout, getInfo } from '@/api/login'
import { getlogger, setlogger } from '@/api/logger'
import { getName, setName, getToken, setToken } from '@/utils/dao'

const user = {
  state: {
    user: '',
    status: '',
    token: getToken(),
    name: getName(),
    avatar: '',
    roles: [],
    setting: {
      articlePlatform: []
    }
  },

  mutations: {
    SET_TOKEN: (state, token) => {
      state.token = token
      setToken(token)
    },
    SET_SETTING: (state, setting) => {
      state.setting = setting
    },
    SET_STATUS: (state, status) => {
      state.status = status
    },
    SET_NAME: (state, name) => {
      state.name = name
      setName(name)
    },
    SET_AVATAR: (state, avatar) => {
      state.avatar = avatar
    },
    SET_ROLES: (state, roles) => {
      state.roles = roles
    },
    LOGIN_SUCCESS: () => {
      console.log('login success')
    },
    LOGOUT_USER: state => {
      state.user = ''
    }
  },

  actions: {
    // 登录
    Login ({ commit }, userInfo) {
      //console.log("userInfo", userInfo)
      const username = userInfo.username
      return new Promise((resolve, reject) => {
        login(username, userInfo.password).then(response => {
          const data = response.data
          // console.log(response.data)
          if (data.code == 0) {
            commit('SET_TOKEN', data.token)
            commit('SET_NAME', username)
          } 
          // console.log("resolve")
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },

    // 获取用户信息
    GetInfo ({ commit, state }, user) {
      return new Promise((resolve, reject) => {
        getInfo(user).then(response => {
          const data = response.data
          if (!data) {
            reject('Verification failed, please Login again.')
          }

          // role must be a non-empty array
          if (!data.role || data.role.length <= 0) {
            reject('getInfo: role must be a non-null array!')
          }
  
          commit('SET_ROLES', data.role)
          commit('SET_NAME', data.name)
          commit('SET_AVATAR', data.avatar)
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },

    // 登出
    LogOut ({ commit, state }) {
      return new Promise((resolve, reject) => {
        logout(state.token).then(() => {
          commit('SET_TOKEN', '')
          commit('SET_ROLES', [])
          resolve()
        }).catch(error => {
          reject(error)
        })
      })
    },

    // 前端 登出
    FedLogOut ({ commit }, err) {
      return new Promise(resolve => {
        commit('SET_TOKEN', '')
        commit('SET_ROLES', [])
        alert('has logout', err)
        resolve()
      })
    },

    GetLogger({ commit }, info) {
      const {user, page, stime, etime} = info
      // console.log(info)
      // console.log(user, page)
      return new Promise((resolve, reject) => {
        getlogger(user, page, stime, etime).then(response => {
          // console.log(response)
          resolve(response)
        }).catch(error => {
          // console.log(error)
          reject(error)
        })
      })
    },

    SetLogger({ commit }, info) {
      const {user, formdata} = info
      // console.log(formdata)
      return new Promise((resolve, reject) => {
        setlogger(user, formdata).then(response => {
          // console.log(response)
          resolve(response)
        }).catch(error => {
          // console.log(error)
          reject(error)
        })
      })
    },

    // 动态修改权限
    ChangeRole ({ commit }, role) {
      return new Promise(resolve => {
        commit('SET_ROLES', [role])
        commit('SET_TOKEN', role)
        resolve()
      })
    }
  }
}

export default user
