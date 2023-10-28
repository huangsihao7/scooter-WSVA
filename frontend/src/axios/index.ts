/*
 * @Author: Xu Ning
 * @Date: 2023-10-25 7:08:43
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-28 13:06:34
 * @FilePath: \scooter-WSVA\frontend\src\axios\index.ts
 * @Description:
 *
 */
import axios, { type AxiosRequestHeaders } from 'axios'
import router from '@/router'
import { userStore } from '@/stores/user'
import { ElMessage } from 'element-plus'


// const baseURl = 'http://127.0.0.1:8080';
const baseURl = 'http://172.22.121.53:7070'

const service = axios.create({
  baseURL: baseURl,
  timeout: 15000 // 请求超时时间
})

// const { token } = storeToRefs('user')

// request拦截器 添加token
service.interceptors.request.use(
  (config) => {
    const token = userStore().token
    if (token) {
      const reqHeader = config.headers as AxiosRequestHeaders
      reqHeader.Authorization = token
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

//  response拦截器
service.interceptors.response.use((response) => {
  console.log(response,response.data)
  if (response.status === 200) {
    return response.data
  } else if (response.status === 401) {
    ElMessage({
      message: response.data.status_msg + '，请重新登录',
      type: 'error'
    })
    userStore().isLoggedIn = false
    userStore().token = ''
    userStore().avatar = ''
    userStore().username = ''
    router.push('/')
    return Promise.reject()
  } else {
    ElMessage({
      message: response.data.status_msg,
      type: 'error'
    })
    return Promise.reject()
  }
})
export { service, baseURl }
