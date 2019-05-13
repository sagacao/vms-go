/* eslint-disable */
import axios from 'axios'
import store from '../store'
import { Notification } from 'element-ui'
// import vue from 'vue'
// import router from '../router';

// 创建axios实例
const service = axios.create({
  baseURL: process.env.BASE_API, // api的base_url //"http://localhost:8081/",
  timeout: 6000 // 请求超时时间
})

// service.defaults.retry = 4;
// service.defaults.retryDelay = 1000;

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
    console.log('axios err:', error)// for debug
    // console.log(JSON.stringify(error));
    // console.log(error.config)
    // vue.$message.error({
    //   message: error.message,
    //   duration: 5 * 1000,
    //   closable: true
    // })
    // var originalRequest = error.config
    console.log(originalRequest)
    if (error.code === 'ECONNABORTED' && error.config.hasOwnProperty('timeout')) {
      Notification.error('接口请求超时，请刷新!')
      // originalRequest._retry = true
      // return service.request(originalRequest);
    } else {
      Notification.error(error.msg || '接口返回出错!')
    }
    return Promise.reject(error)
  }
)

// service.interceptors.response.use(undefined, function axiosRetryInterceptor(err) {
//   var config = err.config;
//   console.log(config)
//   // If config does not exist or the retry option is not set, reject
//   if(!config || !config.retry) return Promise.reject(err);
  
//   // Set the variable for keeping track of the retry count
//   config.__retryCount = config.__retryCount || 0;
  
//   // Check if we've maxed out the total number of retries
//   if(config.__retryCount >= config.retry) {
//       // Reject with the error
//       return Promise.reject(err);
//   }
  
//   // Increase the retry count
//   config.__retryCount += 1;
  
//   // Create new promise to handle exponential backoff
//   var backoff = new Promise(function(resolve) {
//       setTimeout(function() {
//           resolve();
//       }, config.retryDelay || 1);
//   });
  
//   // Return the promise in which recalls axios to retry the request
//   return backoff.then(function() {
//       return axios(config);
//   });
// });

export default service
