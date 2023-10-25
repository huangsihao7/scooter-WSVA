/*
 * @Author: Xu Ning
 * @Date: 2023-10-22 19:52:56
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-22 19:53:02
 * @Description: 
 * @FilePath: \myMap\src\service\modules\login\types.ts
 */
export interface ILoginParams {
    userName: string
    passWord: string | number
   }
   export interface ILoginApi {
    login: (params: ILoginParams)=> Promise<any>
   }
   