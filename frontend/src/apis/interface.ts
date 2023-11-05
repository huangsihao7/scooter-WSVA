/*
 * @Author: Xu Ning
 * @Date: 2023-10-27 23:44:28
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-05 22:45:33
 * @Description: 数据结构
 * @FilePath: \scooter-WSVA\frontend\src\apis\interface.ts
 */

interface VideoType {
  video_id: number;
  author: UserType;
  play_url: string;
  cover_url: string;
  favorite_count: number;
  comment_count: number;
  is_favorite: boolean;
  title: number;
  create_time: number;
  is_star: boolean;
  star_count: number;
  share_count: number;
  duration: string;
}

interface LiveType {
  uid: number;
  avatar: string;
  name: string;
  live_play_url: string;
  live_cover_url: string;
  is_follow: boolean;
  signature: string;
}

interface CommentType {
  comment_id: number;
  content: string;
  create_date: string;
  user: UserType;
}

interface UserType {
  id: number;
  name: string;
  gender: number;
  avatar: string;
  signature: string;
  follow_count: number;
  follower_count: number;
  total_favorited: number;
  work_count: number;
  favorite_count: number;
  is_follow: boolean;
  phoneNum?: string;
  background_image: string;
}

interface DownloadType {
  url: string;
  title: string;
}

interface FollowCardType {
  id: number;
  name: string;
  gender: number;
  avatar: string;
  signature: string;
  follow_count: number;
  follower_count: number;
  total_favorited: number;
  work_count: number;
  favorite_count: number;
  is_follow?: boolean;
  is_friends?: boolean;
  phoneNum?: string;
  background_image: string;
  cover_url: string;
  video_id: number;
}

export type {
  VideoType,
  DownloadType,
  UserType,
  CommentType,
  FollowCardType,
  LiveType,
};
