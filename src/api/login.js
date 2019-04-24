import request from '@/utils/request'

export function login (email, password) {
  const data = {
    email,
    password
  }
  return request({
    url: '/login/login',
    method: 'post',
    data
  })
}

export function logout () {
  return request({
    url: '/login/logout',
    method: 'post'
  })
}

export function getInfo (token) {
  return request({
    url: '/user/info',
    method: 'get',
    params: { token }
  })
}
