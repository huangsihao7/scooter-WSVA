/*
 * @Author: Xu Ning
 * @Date: 2023-10-22 19:49:44
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-26 12:13:18
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
      component: () => import('../view/UserView.vue')
    },
    {
      path: '/follow',
      name: 'follow',
      component: () => import('../view/FollowView.vue')
    },
    {
      path: '/hot',
      name: 'hot',
      component: () => import('../view/ClassifiedView.vue')
    },
    {
      path: '/recreation',
      name: 'recreation',
      component: () => import('../view/ClassifiedView.vue')
    },
    {
      path: '/sports',
      name: 'sports',
      component: () => import('../view/ClassifiedView.vue')
    },
    {
      path: '/food',
      name: 'food',
      component: () => import('../view/ClassifiedView.vue')
    },
    {
      path: '/cartoon',
      name: 'cartoon',
      component: () => import('../view/ClassifiedView.vue')
    },
    {
      path: '/knowledge',
      name: 'knowledge',
      component: () => import('../view/ClassifiedView.vue')
    },
    {
      path: '/',
      redirect: '/rec'
    }
    
  ]
})

export default router

