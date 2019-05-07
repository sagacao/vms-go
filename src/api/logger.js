/* eslint-disable */ 
import request from '@/utils/request'
import qs from 'qs'

export function getlogger (user, page, stime, etime) {
    // console.log('getlogger', user, page)
    return request({
        url: '/vms/logger',
        method: 'get',
        params: { user, page, stime, etime }
    })
}

export function setlogger (user, form) {
    // console.log('setlogger', user, form)
    const data = {
        channel : form.channel,
        game: form.game,
        logdate: form.date,
        newly: form.newly,
        tow_pr: form.tow_pr,
        three_pr: form.three_pr,
        seven_pr: form.seven_pr,
        retention: form.retention
    }
    return request({
        url: '/vms/logger',
        method: 'post',
        params: { user },
        transformRequest: [function (data) {
            return qs.stringify(data)
        }],
        data
    })
}
