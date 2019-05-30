/* eslint-disable */ 
import request from '@/utils/request'
import qs from 'qs'

export function getSvrAppInfo (user, page, game) {
    // console.log('getSvrGates', user, page)
    return request({
        url: '/vms/vapp/req',
        method: 'get',
        params: { user, page, game }
    })
}

export function setSvrAppInfo (user, form) {
    console.log('setSvrAppInfo', user, form)
    const data = {
        game: form.game,
        channel: form.channel,
        appid : form.appid,
        secret : form.secret,
        desc : form.desc
    }
    return request({
        url: '/vms/vapp/edit',
        method: 'post',
        params: { user },
        transformRequest: [function (data) {
            return qs.stringify(data)
        }],
        data
    })
}

export function removeSvrAppInfo (user, form) {
    // console.log('removeSvrGates', user, form)
    const data = {
        game: form.game,
        appid: form.appid
    }
    return request({
        url: '/vms/vapp/remove',
        method: 'post',
        params: { user },
        transformRequest: [function (data) {
            return qs.stringify(data)
        }],
        data
    })
}