/* eslint-disable */ 
import request from '@/utils/request'

export function getPayDetail (user, page, channel, stime, etime) {
    // console.log('getPayDetail', user, page, channel, stime, etime)
    return request({
        url: '/vms/pay/detail',
        method: 'get',
        params: { user, page, channel, stime, etime }
    })
}

export function getPaySum (user, page, channel, stime, etime) {
    // console.log('getPaySum', user, page, channel, stime, etime)
    return request({
        url: '/vms/pay/sum',
        method: 'get',
        params: { user, page, channel, stime, etime }
    })
}