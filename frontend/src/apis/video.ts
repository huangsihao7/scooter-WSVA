import { service } from "@/axios";

export function postVideo(
  url: string,
  coverUrl: string,
  title: string,
  category: number,
) {
  return service({
    url: "/feed/create",
    method: "post",
    data: { url, coverUrl, title, category },
  });
}

export function userVideoListReq(toUid: string) {
  return service({
    url: "/feed/UserVideosList",
    method: "post",
    data: { toUid },
  });
}

export function getVideosList() {
  return service({
    url: `/feed/VideosList`, // 使用字符串模板来拼接runId
    method: "get",
  });
}

export function getCategoryVideosList(category: string) {
  return service({
    url: `/feed/CategoryVideosList/${category}`, // 使用字符串模板来拼接runId
    method: "get",
  });
}
