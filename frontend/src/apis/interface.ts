/*
 * @Author: Xu Ning
 * @Date: 2023-10-27 23:44:28
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-29 20:13:40
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
  isStar: boolean;
  starCount: number;
  shareCount: number;
}

interface CommentType {
  content: string;
  createDate: string;
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

export type { VideoType, UserType, CommentType };
