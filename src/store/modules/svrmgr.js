/* eslint-disable */
import { getSvrCfg, setSvrCfg, getSvrSwitch, setSvrSwitch, removeSvrSwitch } from '@/api/svrmgr'
import { getSvrGates, setSvrGates, removeSvrGates } from '@/api/gates'

const state = {
}

const mutations = {
}

const actions= {
    /////////////////////////////////////////////////
    //// settings
    GetSvrCfg({ commit }, info) {
        const {user, page, game} = info
        return new Promise((resolve, reject) => {
            getSvrCfg(user, page, game).then(response => {
                // console.log(response)
                resolve(response)
            }).catch(error => {
                // console.log(error)
                reject(error)
            })
        })
    },

    SetSvrCfg({ commit }, info) {
        const {user, formdata} = info
        // console.log(formdata)
        return new Promise((resolve, reject) => {
            setSvrCfg(user, formdata).then(response => {
                // console.log(response)
                resolve(response)
            }).catch(error => {
                // console.log(error)
                reject(error)
            })
        })
    },

    GetSvrSwitch({ commit }, info) {
        const {user, page, game} = info
        return new Promise((resolve, reject) => {
            getSvrSwitch(user, page, game).then(response => {
                // console.log(response)
                resolve(response)
            }).catch(error => {
                // console.log(error)
                reject(error)
            })
        })
    },

    SetSvrSwitch({ commit }, info) {
        const {user, formdata} = info
        // console.log(formdata)
        return new Promise((resolve, reject) => {
            setSvrSwitch(user, formdata).then(response => {
                // console.log(response)
                resolve(response)
            }).catch(error => {
                // console.log(error)
                reject(error)
            })
        })
    },

    RemoveSvrSwitch({ commit }, info) {
        const {user, formdata} = info
        // console.log(formdata)
        return new Promise((resolve, reject) => {
            removeSvrSwitch(user, formdata).then(response => {
                // console.log(response)
                resolve(response)
            }).catch(error => {
                // console.log(error)
                reject(error)
            })
        })
    },

    GetSvrGates({ commit }, info) {
        const {user, page, game} = info
        return new Promise((resolve, reject) => {
            getSvrGates(user, page, game).then(response => {
                // console.log(response)
                resolve(response)
            }).catch(error => {
                // console.log(error)
                reject(error)
            })
        })
    },

    SetSvrGates({ commit }, info) {
        const {user, formdata} = info
        // console.log(formdata)
        return new Promise((resolve, reject) => {
            setSvrGates(user, formdata).then(response => {
                // console.log(response)
                resolve(response)
            }).catch(error => {
                // console.log(error)
                reject(error)
            })
        })
    },

    RemoveSvrGates({ commit }, info) {
        const {user, formdata} = info
        // console.log(formdata)
        return new Promise((resolve, reject) => {
            removeSvrGates(user, formdata).then(response => {
                // console.log(response)
                resolve(response)
            }).catch(error => {
                // console.log(error)
                reject(error)
            })
        })
    },

    placeholder() {

    }
}

export default {
    namespaced: true,
    state,
    mutations,
    actions
}
  