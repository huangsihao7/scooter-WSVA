import { service } from "@/axios";

export function getUserInfo(user_id: number) {
  return service({
    url: "/user/userinfo",
    method: "post",
    data: { user_id },
  });
}
