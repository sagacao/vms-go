/* eslint-disable */
import axios from 'axios'
import store from '../store'
import vue from 'vue'
// import router from '../router';

// 创建axios实例
const service = axios.create({
  baseURL: process.env.BASE_API, // api的base_url //"http://localhost:8081/",
  timeout: 5000 // 请求超时时间
})

//service.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded';

function checkStatus (response) {//检查状态码的方法
  if (response && (response.status == 200 || response.status == 304 || response.status == 400)){
      return response;
  }
  return {
      status:-404,
      msg:'网络异常'
  }
}

function checkCode(response) {//如果失败，弹出错误信息
  if (response.status === -404){
      alert(response.msg);
  }
  return res;
}

// request拦截器
service.interceptors.request.use(config => {
  // Do something before request is sent
  if (store.getters.token) {
    config.headers.Authorization = `Bearer ${store.getters.token}`
  }
  return config
}, error => {
  // Do something with request error
  console.log(error) // for debug
  Promise.reject(error)
})

// respone拦截器
service.interceptors.response.use(
  response => response,
  error => {
    console.log('err' + error)// for debug
    vue.$message.error({
      message: error.message,
      duration: 5 * 1000,
      closable: true
    })
    return Promise.reject(error)
  }
)

export default service
