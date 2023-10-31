/*
 * @Author: Xu Ning
 * @Date: 2023-10-27 23:44:28
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-30 20:59:53
 * @Description:
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
