/*
 * @Author: Xu Ning
 * @Date: 2023-10-22 19:52:42
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-22 19:52:46
 * @Description: 
 * @FilePath: \myMap\src\service\modules\login\login.ts
 */
import http from '@/service/http'
import * as T from './types'
​
const loginApi: T.ILoginApi = {
    login(params){
        return http.post('/login', params)
    }
​
}
export default loginApi