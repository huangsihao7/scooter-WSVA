/*
 * @Author: Xu Ning
 * @Date: 2023-10-25 7:08:43
 * @LastEditors: huangsihao7 1057434651@qq.com
 * @LastEditTime: 2023-10-30 16:19:02
 * @FilePath: /scooter-WSVA/frontend/src/axios/index.ts
 * @Description:
 *
 */
import axios, { type AxiosRequestHeaders } from "axios";
import router from "@/router";
import { userStore } from "@/stores/user";
import { useMessage } from "naive-ui";

// const baseURl = 'http://127.0.0.1:8080';
const baseURL = "http://172.22.121.53:7070";
const message = useMessage();
const service = axios.create({
  baseURL: baseURL,
  timeout: 15000, // 请求超时时间
});

// const { token } = storeToRefs('user')

// request拦截器 添加token
service.interceptors.request.use(
  (config) => {
    const token = userStore().token;
    if (token) {
      const reqHeader = config.headers as AxiosRequestHeaders;
      reqHeader.Authorization = token;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  },
);

//  response拦截器
service.interceptors.response.use((response) => {
  // console.log(response, response.data);
  if (response.status === 200) {
    return response.data;
  } else if (response.status === 401) {
    message.error(response.data.status_msg + "，请重新登录");
    userStore().isLoggedIn = false;
    userStore().token = "";
    userStore().avatar = "";
    userStore().phoneNum = "";
    router.push("/");
    return Promise.reject();
  } else {
    message.error(response.data.status_msg);
    return Promise.reject();
  }
});
export { service, baseURL };
