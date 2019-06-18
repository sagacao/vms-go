/* eslint-disable */ 
import request from '@/utils/request'
import qs from 'qs'

export function getgifts (user, page, game) {
    // console.log('getstats', user, page)
    return request({
        url: '/vms/code/get',
        method: 'get',
        params: { user, page, game }
    })
}

export function gengifts (user, form) {
    // console.log('gengifts', user, form)
    const data = {
        channel : form.channel,
        game: form.game,
        number: form.number,
        atype: form.atype,
        acount: form.acount
    }
    return request({
        url: '/vms/code/gen',
        method: 'post',
        params: { user },
        transformRequest: [function (data) {
            return qs.stringify(data)
        }],
        data
    })
}