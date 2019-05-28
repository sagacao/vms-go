/* eslint-disable */ 
import request from '@/utils/request'
import qs from 'qs'

export function getSvrGates (user, page, game) {
    // console.log('getSvrGates', user, page)
    return request({
        url: '/vms/gates/req',
        method: 'get',
        params: { user, page, game }
    })
}

export function setSvrGates (user, form) {
    console.log('setSvrGates', user, form)
    const data = {
        game: form.game,
        name: form.name,
        value : form.value,
        svr : form.svr
    }
    return request({
        url: '/vms/gates/edit',
        method: 'post',
        params: { user },
        transformRequest: [function (data) {
            return qs.stringify(data)
        }],
        data
    })
}

export function removeSvrGates (user, form) {
    // console.log('removeSvrGates', user, form)
    const data = {
        game: form.game,
        name: form.name
    }
    return request({
        url: '/vms/gates/remove',
        method: 'post',
        params: { user },
        transformRequest: [function (data) {
            return qs.stringify(data)
        }],
        data
    })
}