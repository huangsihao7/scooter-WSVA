/*
 * @Author: Xu Ning
 * @Date: 2023-10-22 19:49:44
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-25 14:28:32
 * @Description: 
 * @FilePath: \scooter-wsva\frontend\src\router\index.ts
 */
import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../view/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    // {
    //   path: '/login',
    //   name: 'login',
    //   component: () => import('../views/LoginView.vue')
    // },
    // {
    //   path: '/detect/:id',
    //   name: 'detect',
    //   props: true,
    //   component: () => import('../views/DetectView.vue')
    // },
    // {
    //   path: '/customize-detect',
    //   name: 'customize-detect',
    //   component: () => import('../views/CustomizeDetectView.vue')
    // },
    // {
    //   path: '/execute/:id',
    //   name: 'execute',
    //   props: true,
    //   component: () => import('../views/ExecuteView.vue')
    // },
    // {
    //   path: '/reinforcement/:id',
    //   name: 'reinforcement',
    //   props: true,
    //   component: () => import('../views/ReinforceView.vue')
    // }
  ]
})

export default router

