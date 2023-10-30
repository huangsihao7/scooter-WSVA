/*
 * @Author: Xu Ning
 * @Date: 2023-10-30 13:15:41
 * @LastEditors: huangsihao7 1057434651@qq.com
 * @LastEditTime: 2023-10-30 15:01:45
 * @Description:
 * @FilePath: /scooter-WSVA/frontend/src/apis/relation.ts
 */
import { service } from "@/axios";

export function doFollow(to_user_id: number, action_type: number) {
  return service({
    url: "/relation/action",
    method: "post",
    data: {
      to_user_id: to_user_id,
      action_type: action_type,
    },
  });
}
