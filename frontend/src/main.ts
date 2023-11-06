/*
 * @Author: Xu Ning
 * @Date: 2023-10-22 19:33:20
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-04 17:48:29
 * @Description: 配置挂载
 * @FilePath: \scooter-WSVA\frontend\src\main.ts
 */
import { createApp } from "vue";
import { createPinia } from "pinia";
import piniaPluginPersistedstate from "pinia-plugin-persistedstate";
import VueLazyLoad from "vue3-lazyload";
import App from "./App.vue";
import router from "./router/index";
import "./style.css";
import "animate.css";
import './assets/iconfont/iconfont.css'

const pinia = createPinia();
const app = createApp(App);
pinia.use(piniaPluginPersistedstate);
app.use(pinia);
app.use(VueLazyLoad, {});
app.use(router);
app.mount("#app");
