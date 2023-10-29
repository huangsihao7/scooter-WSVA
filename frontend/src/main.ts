/*
 * @Author: Xu Ning
 * @Date: 2023-10-22 19:33:20
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-27 12:54:21
 * @Description:
 * @FilePath: \scooter-WSVA\frontend\src\main.ts
 */
import { createApp } from "vue";
import { createPinia } from "pinia";
import piniaPluginPersistedstate from "pinia-plugin-persistedstate";
import App from "./App.vue";
import router from "./router/index";
import ElementPlus from "element-plus";
import * as ElementPlusIconsVue from "@element-plus/icons-vue";
import "element-plus/dist/index.css";
import "./style.css";
import "animate.css";

const pinia = createPinia();
const app = createApp(App);
Object.keys(ElementPlusIconsVue).forEach((key) => {
  app.component(
    key,
    ElementPlusIconsVue[key as keyof typeof ElementPlusIconsVue],
  );
});
pinia.use(piniaPluginPersistedstate);
app.use(pinia);
app.use(ElementPlus);
app.use(router);
app.mount("#app");
