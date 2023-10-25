/*
 * @Author: Xu Ning
 * @Date: 2023-10-22 19:49:44
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-26 00:10:34
 * @Description: 
 * @FilePath: \scooter-wsva\frontend\src\router\index.ts
 */
import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../view/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/rec',
      name: 'rec',
      component: HomeView
    },
    {
      path: '/user',
      name: 'user',
      component: () => import('../view/User.vue')
    },
    {
      path: '/follow',
      name: 'follow',
      component: () => import('../view/Follow.vue')
    },
    {
      path: '/',
      redirect: '/rec'
    }
    
  ]
})

export default router

