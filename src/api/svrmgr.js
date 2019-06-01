/* eslint-disable */ 
import request from '@/utils/request'
import qs from 'qs'

export function getSvrCfg (user, page, game) {
    // console.log('getSvrCfg', user, page)
    return request({
        url: '/vms/cfg/req',
        method: 'get',
        params: { user, page, game }
    })
}

export function setSvrCfg (user, form) {
    // console.log('setSvrCfg', user, form)
    const data = {
        game: form.game,
        value: form.jsonvalue
    }
    return request({
        url: '/vms/cfg/edit',
        method: 'post',
        params: { user },
        transformRequest: [function (data) {
            return qs.stringify(data)
        }],
        data
    })
}

export function getSvrSwitch (user, page, game) {
    // console.log('getSvrSwitch', user, page)
    return request({
        url: '/vms/switch/req',
        method: 'get',
        params: { user, page, game }
    })
}

export function setSvrSwitch (user, form) {
    console.log('setSvrSwitch', user, form)
    const data = {
        game: form.game,
        funcname: form.funcname,
        funcswitch : form.funcswitch
    }
    return request({
        url: '/vms/switch/edit',
        method: 'post',
        params: { user },
        transformRequest: [function (data) {
            return qs.stringify(data)
        }],
        data
    })
}

export function removeSvrSwitch (user, form) {
    // console.log('removeSvrSwitch', user, form)
    const data = {
        game: form.game,
        funcname: form.funcname,
        funcswitch : ''
    }
    return request({
        url: '/vms/switch/remove',
        method: 'post',
        params: { user },
        transformRequest: [function (data) {
            return qs.stringify(data)
        }],
        data
    })
}