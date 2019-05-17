/* eslint-disable */ 
import request from '@/utils/request'
import qs from 'qs'

export function getstats (user, page, stime, etime) {
    // console.log('getstats', user, page)
    return request({
        url: '/vms/stats/get',
        method: 'get',
        params: { user, page, stime, etime }
    })
}

export function setstats (user, form) {
    // console.log('setstats', user, form)
    const data = {
        channel : form.channel,
        game: form.game,
        logdate: form.logdate,
        newly: form.newly,
        tow_pr: form.tow_pr,
        three_pr: form.three_pr,
        seven_pr: form.seven_pr,
        retention: form.retention
    }
    return request({
        url: '/vms/stats/edit',
        method: 'post',
        params: { user },
        transformRequest: [function (data) {
            return qs.stringify(data)
        }],
        data
    })
}

export function delstats (user, form) {
    // console.log('setstats', user, form)
    const data = {
        channel : form.channel,
        game: form.game,
        logdate: form.logdate
    }
    return request({
        url: '/vms/stats/rm',
        method: 'post',
        params: { user },
        transformRequest: [function (data) {
            return qs.stringify(data)
        }],
        data
    })
}