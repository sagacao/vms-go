/* eslint-disable */ 
import request from '@/utils/request'
import qs from 'qs'

export function getSvrChannelStatus (user, page, game) {
    // console.log('getSvrChannelStatus', user, page)
    return request({
        url: '/vms/cstatus/req',
        method: 'get',
        params: { user, page, game }
    })
}

export function setSvrChannelStatus (user, form) {
    console.log('setSvrChannelStatus', user, form)
    const data = {
        game: form.game,
        channel: form.channel,
        name : form.name,
        status : form.status,
        action : form.action
    }
    return request({
        url: '/vms/cstatus/edit',
        method: 'post',
        params: { user },
        transformRequest: [function (data) {
            return qs.stringify(data)
        }],
        data
    })
}

export function removeChannelStatus (user, form) {
    // console.log('removeChannelStatus', user, form)
    const data = {
        game: form.game,
        channel: form.channel,
        name : form.name,
    }
    return request({
        url: '/vms/cstatus/remove',
        method: 'post',
        params: { user },
        transformRequest: [function (data) {
            return qs.stringify(data)
        }],
        data
    })
}