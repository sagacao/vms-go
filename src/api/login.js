/* eslint-disable */ 
import request from '@/utils/request'
import md5 from 'js-md5'
import qs from 'qs'

export function login (user, password) {
  const data = {
    user : user,
    password : md5(password)
  }
  return request({
    url: '/login',
    method: 'post',
    transformRequest: [function (data) {
      return qs.stringify(data)
    }],
    data
  })
}

export function logout () {
  return request({
    url: '/logout',
    method: 'post'
  })
}

export function getInfo (user) {
  return request({
    url: '/vms/user/info',
    method: 'get',
    params: { user }
  })
}
