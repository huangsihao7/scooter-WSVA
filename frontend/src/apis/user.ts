import { service } from "@/axios";

export function getUserInfo(uid: number) {
  return service({
    url: "/user/userinfo",
    method: "post",
    data: { uid },
  });
}


export function updateUserInfo(name: string, gender: number, avatar:string, dec: string, background_image: string) {
  return service({
    url: "/user/update",
    method: "post",
    data: { name, gender, avatar, dec, background_image },
  });
}