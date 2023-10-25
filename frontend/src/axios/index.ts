/*
 * @Author: Xu Ning
 * @Date: 2023-10-25 7:08:43
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-25 14:17:52
 * @FilePath: \scooter-wsva\frontend\src\axios\index.ts
 * @Description:
 *
 */
import { computed, ref } from 'vue'
import axios, { type AxiosRequestHeaders } from 'axios'
import { createDiscreteApi, lightTheme, darkTheme, type ConfigProviderProps } from 'naive-ui'
import router from '@/router'
import { userStore } from '@/stores/user'

// const baseURl = 'http://127.0.0.1:8080';
const baseURl = 'http://172.22.121.53:8080'

const service = axios.create({
  baseURL: baseURl,
  timeout: 15000 // 请求超时时间
})

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

const themeRef = ref<'light' | 'dark'>('dark')
const configProviderPropsRef = computed<ConfigProviderProps>(() => ({
  theme: themeRef.value === 'light' ? lightTheme : darkTheme
}))
const { message } = createDiscreteApi(['message'], { configProviderProps: configProviderPropsRef })

//  response拦截器
service.interceptors.response.use((response) => {
  if (response.data.status === 200) {
    return response.data.data
  } else if (response.data.status === 401) {
    message.error(response.data.msg + '，请重新登录')
    userStore().logout()
    router.push('/login')
    return Promise.reject()
  } else {
    message.error(response.data.msg)
    return Promise.reject()
  }
})
export { service, baseURl }
