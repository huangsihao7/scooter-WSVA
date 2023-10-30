import { service } from "@/axios";

export function doStar(video_id: number, action_type: number) {
  return service({
    url: "/star/action",
    method: "post",
    data: {
      video_id: video_id,
      action_type: action_type,
    },
  });
}

export function doFavourite(video_id: number, action_type: number) {
  return service({
    url: "/favorite/action",
    method: "post",
    data: {
      video_id: video_id,
      action_type: action_type,
    },
  });
}

export function userFavouriteListReq(uid: number) {
  return service({
    url: `/favorite/list?uid=${uid}`,
    method: "get",
  });
}
