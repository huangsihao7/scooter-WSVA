/*
 * @Author: Xu Ning
 * @Date: 2023-10-27 23:44:28
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-28 13:51:15
 * @Description:
 * @FilePath: \scooter-WSVA\frontend\src\apis\interface.ts
 */

interface VideoType {
  id: number;
  author: UserType;
  playUrl: string;
  coverUrl: string;
  favoriteCount: number;
  commentCount: number;
  isFavorite: boolean;
  title: number;
  createTime: number;
  isCollect: boolean;
  collectCount: number;
  shareCount: number;
}

interface UserType {
  id: number;
  name: any;
  gender?: number;
  avatar: string;
  signature?: string;
  follow_count: number;
  follower_count: number;
  total_favorited: number;
  work_count: number;
  favorite_count: number;
  is_follow: boolean;
}

export type { VideoType, UserType };
