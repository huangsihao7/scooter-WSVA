/*
 * @Author: Xu Ning
 * @Date: 2023-10-25 7:08:43
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-05 16:52:18
 * @FilePath: \scooter-WSVA\frontend\src\axios\index.ts
 * @Description: 请求配置
 *
 */
import axios, { type AxiosRequestHeaders } from "axios";
import router from "@/router";
import { userStore } from "@/stores/user";
import { createDiscreteApi } from "naive-ui";
import { routeStore } from "@/stores/route";

// const baseURl = 'http://127.0.0.1:8080';
const baseURL = "http://172.22.121.53:7070";

const service = axios.create({
  baseURL: baseURL,
  timeout: 15000, // 请求超时时间
});
const { message } = createDiscreteApi(["message"]);

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
  console.log(response);
  console.log(response.data);
  if (response.data.status_code === 200) {
    return response.data;
  } else if (response.status === 401) {
    message.error(response.data.status_message + "，请重新登录");
    userStore().isLoggedIn = false;
    userStore().user_id = -1;
    userStore().token = "";
    userStore().avatar = "";
    userStore().phoneNum = "";
    userStore().gender = 1;
    userStore().signature = "";
    userStore().background_image = "";
    router.push("/rec");
    routeStore().name = "rec";
    window.location.reload();
    return Promise.reject();
  } else {
    message.error(response.data.status_message);
    return Promise.reject();
  }
});
export { service, baseURL };
