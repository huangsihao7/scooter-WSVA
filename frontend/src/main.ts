/*
 * @Author: Xu Ning
 * @Date: 2023-10-22 19:33:20
 * @LastEditors: huangsihao7 1057434651@qq.com
 * @LastEditTime: 2023-11-03 16:33:05
 * @Description:
 * @FilePath: /scooter-WSVA/frontend/src/main.ts
 */
import { createApp } from "vue";
import { createPinia } from "pinia";
import piniaPluginPersistedstate from "pinia-plugin-persistedstate";
import App from "./App.vue";
import router from "./router/index";
import "./style.css";
import "animate.css";

const pinia = createPinia();
const app = createApp(App);
pinia.use(piniaPluginPersistedstate);
app.use(pinia);
app.use(router);
app.mount("#app");
