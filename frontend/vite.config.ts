/*
 * @Author: huangsihao7
 * @Date: 2023-10-29 11:31:39
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-04 15:03:23
 * @FilePath: \scooter-WSVA\frontend\vite.config.ts
 * @Description: 
 */
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'
import {viteMockServe} from "vite-plugin-mock";

const resolve = (dir: string) => path.join(__dirname, dir)
export default defineConfig({
  plugins: [vue(),
    viteMockServe({
      mockPath: "./src/mock",
      localEnabled: true,
    }),
    AutoImport({
      resolvers: [ElementPlusResolver()],
    }),
    Components({
      resolvers: [ElementPlusResolver()],
    }),
  ],
  resolve: {
    alias: {
      '@': resolve('src'),
      comps: resolve('src/components'),
      apis: resolve('src/apis'),
      views: resolve('src/views'),
      utils: resolve('src/utils'),
      routes: resolve('src/routes'),
      styles: resolve('src/styles')
    }
  },
  server: {
    // 配置前端服务地址和端口
    //服务器主机名
    host: '127.0.0.1',
    //端口号
    port: 3088,
    //设为 true 时若端口已被占用则会直接退出，而不是尝试下一个可用端口
    strictPort: false,
    //服务器启动时自动在浏览器中打开应用程序,当此值为字符串时，会被用作 URL 的路径名
    open: false,
    //自定义代理规则
    proxy: {
      // 选项写法
      '/api': {
        target: 'http://localhost:3000',
        changeOrigin: true,
        rewrite: path => path.replace(/^\/api/, '')
      }
    }
  }
})
