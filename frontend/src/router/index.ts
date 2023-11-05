/*
 * @Author: Xu Ning
 * @Date: 2023-10-22 19:49:44
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-02 16:38:10
 * @Description: 路由表
 * @FilePath: \scooter-WSVA\frontend\src\router\index.ts
 */
import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../view/HomeView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/rec",
      name: "rec",
      component: HomeView,
    },
    {
      // 用户资料页面
      path: "/user",
      name: "user",
      component: () => import("../view/UserView.vue"),
    },
    {
      // 用户关注的人
      path: "/follow",
      name: "follow",
      component: () => import("../view/FollowView.vue"),
    },
    {
      path: "/hot",
      name: "hot",
      component: () => import("../view/HotView.vue"),
    },
    {
      path: "/recreation",
      name: "recreation",
      component: () => import("../view/ClassifiedView.vue"),
    },
    {
      path: "/sports",
      name: "sports",
      component: () => import("../view/ClassifiedView.vue"),
    },
    {
      path: "/food",
      name: "food",
      component: () => import("../view/ClassifiedView.vue"),
    },
    {
      path: "/cartoon",
      name: "cartoon",
      component: () => import("../view/ClassifiedView.vue"),
    },
    {
      path: "/knowledge",
      name: "knowledge",
      component: () => import("../view/ClassifiedView.vue"),
    },
    {
      // 某个人的个人资料页面
      path: "/userinfo/:id",
      name: "userinfo",
      props: true,
      component: () => import("../view/UserView.vue"),
    },
    {
      // 关注的人的卡片列表
      path: "/follows/:id",
      name: "follows",
      props: true,
      component: () => import("../view/FollowView.vue"),
    },
    {
      // 粉丝的卡片列表
      path: "/followers/:id",
      name: "followers",
      props: true,
      component: () => import("../view/FollowView.vue"),
    },
    {
      // 朋友的卡片列表
      path: "/friends/:id",
      name: "friends",
      props: true,
      component: () => import("../view/FollowView.vue"),
    },
    {
      path: "/video/:id",
      name: "video",
      props: true,
      component: () => import("../view/VideoView.vue"),
    },
    {
      path: "/search",
      name: "search",
      component: () => import("../view/ClassifiedView.vue"),
    },
    {
      path: "/liveSquare",
      name: "liveSquare",
      component: () => import("../view/LiveSquareView.vue"),
    },
    {
      path: "/live/:id",
      name: "live",
      props: true,
      component: () => import("../view/LiveView.vue"),
    },
    {
      path: "/",
      redirect: "/rec",
    },
  ],
});

export default router;
