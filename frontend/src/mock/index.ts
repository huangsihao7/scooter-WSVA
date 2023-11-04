/*
 * @Author: Xu Ning
 * @Date: 2023-11-04 13:22:45
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-04 13:33:22
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\mock\index.ts
 */
import { mock, Random } from 'mockjs'
// 使用mock.js生成avatar图片URL
const avatar = Random.image("200x200", "#50B347", "#FFF", "Avatar");

// 使用mock.js生成User对象的Mock数据
const mockUser = mock({
"avatar": avatar,
"favorite_count|1-100": 1,
"follow_count|1-100": 1,
"follower_count|1-100": 1,
"friend_count|1-100": 1,
"gender|0-1": 1,
"id|1-1000": 1,
"is_follow|1": true,
"name": Random.name(),
"signature": Random.sentence(),
"total_favorited|1-100": 1,
"work_count|1-10": 1,
});

export default [
    {
        url: "/user/login", // 模拟登录的链接
        method: "post", // 请求方式
        timeout: 3000, // 超时时间
        statusCode: 200, // 返回的http状态码
        response: { // 返回的结果集
                status_code: 200,
                status_message: '登录成功',
                user: mockUser
        },
    },
  {
        url: "/mock/getMenuList", // 模拟登录的链接
        method: "get", // 请求方式
        timeout: 3000, // 超时时间
        statusCode: 200, // 返回的http状态码
        response: { // 返回的结果集
            code: 200,
            message: "获取菜单",
            data: {
                content: "获取菜单成功"
            },
        },
    },
]