/* eslint-disable */
import { getSvrCfg, setSvrCfg, getSvrSwitch, setSvrSwitch, removeSvrSwitch } from '@/api/svrmgr'
import { getSvrGates, setSvrGates, removeSvrGates } from '@/api/gates'
import { getSvrAppInfo, setSvrAppInfo, removeSvrAppInfo } from '@/api/appctor'
import { getSvrChannelStatus, setSvrChannelStatus, removeChannelStatus } from '@/api/channelstatus'
import { gengifts, getgifts } from '@/api/gifts'

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

    GetSvrAppInfo({ commit }, info) {
        const {user, page, game} = info
        return new Promise((resolve, reject) => {
            getSvrAppInfo(user, page, game).then(response => {
                // console.log(response)
                resolve(response)
            }).catch(error => {
                // console.log(error)
                reject(error)
            })
        })
    },

    SetSvrAppInfo({ commit }, info) {
        const {user, formdata} = info
        // console.log(formdata)
        return new Promise((resolve, reject) => {
            setSvrAppInfo(user, formdata).then(response => {
                // console.log(response)
                resolve(response)
            }).catch(error => {
                // console.log(error)
                reject(error)
            })
        })
    },

    RemoveSvrAppInfo({ commit }, info) {
        const {user, formdata} = info
        // console.log(formdata)
        return new Promise((resolve, reject) => {
            removeSvrAppInfo(user, formdata).then(response => {
                // console.log(response)
                resolve(response)
            }).catch(error => {
                // console.log(error)
                reject(error)
            })
        })
    },

    GetSvrChannelStatus({ commit }, info) {
        const {user, page, game} = info
        return new Promise((resolve, reject) => {
            getSvrChannelStatus(user, page, game).then(response => {
                // console.log(response)
                resolve(response)
            }).catch(error => {
                // console.log(error)
                reject(error)
            })
        })
    },

    SetSvrChannelStatus({ commit }, info) {
        const {user, formdata} = info
        // console.log(formdata)
        return new Promise((resolve, reject) => {
            setSvrChannelStatus(user, formdata).then(response => {
                // console.log(response)
                resolve(response)
            }).catch(error => {
                // console.log(error)
                reject(error)
            })
        })
    },

    RemoveSvrChannelStatus({ commit }, info) {
        const {user, formdata} = info
        // console.log(formdata)
        return new Promise((resolve, reject) => {
            removeChannelStatus(user, formdata).then(response => {
                // console.log(response)
                resolve(response)
            }).catch(error => {
                // console.log(error)
                reject(error)
            })
        })
    },
    
    Getgifts({ commit }, info) {
        const {user, page, game} = info
        return new Promise((resolve, reject) => {
            getgifts(user, page, game).then(response => {
                // console.log(response)
                resolve(response)
            }).catch(error => {
                // console.log(error)
                reject(error)
            })
        })
    },

    Gengifts({ commit }, info) {
        const {user, formdata} = info
        // console.log(formdata)
        return new Promise((resolve, reject) => {
            gengifts(user, formdata).then(response => {
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
  