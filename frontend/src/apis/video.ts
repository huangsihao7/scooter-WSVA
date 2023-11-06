/*
 * @Author: huangsihao7
 * @Date: 2023-10-29 13:04:21
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-06 11:36:49
 * @FilePath: \scooter-WSVA\frontend\src\apis\video.ts
 * @Description: 视频接口
 */
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

export function userVideoListReq(to_user_id: number) {
  return service({
    url: "/feed/UserVideosList",
    method: "post",
    data: { to_user_id },
  });
}

export function getRecommendVideos(offset: number, readed_videoId: number) {
  return service({
    url: "/feed/recommends",
    method: "post",
    data: { offset, readed_videoId },
  });
}

export function getPopularVideos(offset: number, readed_videoId: number) {
  return service({
    url: "/feed/populars",
    method: "post",
    data: { offset, readed_videoId },
  });
}

export function getVideosByKeyWords(content: string) {
  return service({
    url: "/feed/searcheEs",
    method: "post",
    data: { content },
  });
}

export function getVideosList() {
  return service({
    url: `/feed/VideosList`,
    method: "get",
  });
}

export function getCategoryVideosList(category: number) {
  return service({
    url: `/feed/CategoryVideosList?category=${category}`,
    method: "get",
  });
}

export function getRecommendVideosList(video_id: number) {
  return service({
    url: `/feed/neighbors?video_id=${video_id}`,
    method: "get",
  });
}

export function getHistoryVideosListReq() {
  return service({
    url: `/feed/history`,
    method: "get",
  });
}

export function getVideoById(video_id: number) {
  return service({
    url: `/feed/videoinfo?video_id=${video_id}`,
    method: "get",
  });
}

export function postDeleteVideo(video_id: number) {
  return service({
    url: "/feed/deleteViedo",
    method: "post",
    data: { video_id },
  });
}
