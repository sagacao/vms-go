/* eslint-disable */ 
import request from '@/utils/request'
import qs from 'qs'

export function getPlayerInfo (user, page, game, key, ktype) {
    console.log('getPlayerInfo', user, page, game, key, ktype)
    return request({
        url: '/vms/player/req',
        method: 'get',
        params: { user, page, game, key, ktype }
    })
}

export function editPlayerInfo (user, form) {
    console.log('editPlayerInfo', user, form)
    const data = {
        game : form.game,
        uid : form.uid,
        data : form.data
    }
    return request({
        url: '/vms/player/edit',
        method: 'post',
        params: { user },
        transformRequest: [function (data) {
            return qs.stringify(data)
        }],
        data
    })
}