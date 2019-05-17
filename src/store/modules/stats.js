/* eslint-disable */
import { getstats, setstats } from '@/api/stats'

const state = {
}

const mutations = {
}

const actions = {
  GetStats({ commit }, info) {
    const {user, page, stime, etime} = info
    // console.log(info)
    // console.log(user, page)
    return new Promise((resolve, reject) => {
      getstats(user, page, stime, etime).then(response => {
        // console.log(response)
        resolve(response)
      }).catch(error => {
        // console.log(error)
        reject(error)
      })
    })
  },

  SetStats({ commit }, info) {
    const {user, formdata} = info
    // console.log(formdata)
    return new Promise((resolve, reject) => {
      setstats(user, formdata).then(response => {
        // console.log(response)
        resolve(response)
      }).catch(error => {
        // console.log(error)
        reject(error)
      })
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
