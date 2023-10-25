/*
 * @Author: Xu Ning
 * @Date: 2023-10-22 19:49:44
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-22 20:58:14
 * @Description: 
 * @FilePath: \myMap\src\router\index.ts
 */
import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
const routes: RouteRecordRaw[] = [
  // {
  //   path: '/',
  //   name: 'Login',
  //   component: () => import('@/pages/login/Login.vue'), // 注意这里要带上 文件后缀.vue
  // },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router
