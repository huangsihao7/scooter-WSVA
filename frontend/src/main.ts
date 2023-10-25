/*
 * @Author: Xu Ning
 * @Date: 2023-10-22 19:33:20
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-22 20:03:18
 * @Description: 
 * @FilePath: \myMap\src\main.ts
 */
import {createApp} from 'vue'
import App from './App.vue'
import router from './router/index'
import {createPinia} from "pinia"
import SvgIcon from './components/SvgIcon/index.vue'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import "element-plus/dist/index.css"
import './style.css'

const app = createApp(App)
Object.keys(ElementPlusIconsVue).forEach(key => {
    app.component(key, ElementPlusIconsVue[key as keyof typeof ElementPlusIconsVue]);
});
app.use(createPinia())
app.use(router)
app.mount('#app')
