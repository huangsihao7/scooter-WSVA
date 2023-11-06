<!--
 * @Author: Xu Ning
 * @Date: 2023-11-05 18:42:02
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-06 12:55:18
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\view\LiveSquareView.vue
-->
<script setup lang="ts">
import { ref, onMounted } from "vue";
import LiveCard from "@/components/cards/LiveCard.vue";
// const video_list = [
//   {
//     name: "小红",
//     uid: 1,
//     live_url:
//       "http://pili-live-hls.scooter.peterli.club/scooter/scooter-live.m3u8",
//     live_cover_url:
//       "http://s327crbzf.hn-bkt.clouddn.com/3d2be50e68424af8fe4f2b365884863864a774785fe38cddf405c4afb78e85fe.jpg",
//     avatar:
//       "http://s327crbzf.hn-bkt.clouddn.com/3d2be50e68424af8fe4f2b365884863864a774785fe38cddf405c4afb78e85fe.jpg",
//     is_follow: true,
//     signature: "我是你的 你是我的 谁",
//   },
//   {
//     name: "小黑",
//     uid: 2,
//     live_url:
//       "http://pili-live-hls.scooter.peterli.club/scooter/scooter-live.m3u8",
//     live_cover_url:
//       "http://s327crbzf.hn-bkt.clouddn.com/3d2be50e68424af8fe4f2b365884863864a774785fe38cddf405c4afb78e85fe.jpg",
//     avatar:
//       "http://s327crbzf.hn-bkt.clouddn.com/3d2be50e68424af8fe4f2b365884863864a774785fe38cddf405c4afb78e85fe.jpg",
//     is_follow: true,
//     signature: "我是你的 你是我的kdjashkljghwake 谁",
//   },
//   {
//     name: "小栏",
//     uid: 3,
//     live_url:
//       "http://pili-live-hls.scooter.peterli.club/scooter/scooter-live.m3u8",
//     live_cover_url:
//       "http://s327crbzf.hn-bkt.clouddn.com/3d2be50e68424af8fe4f2b365884863864a774785fe38cddf405c4afb78e85fe.jpg",
//     avatar:
//       "http://s327crbzf.hn-bkt.clouddn.com/3d2be50e68424af8fe4f2b365884863864a774785fe38cddf405c4afb78e85fe.jpg",
//     is_follow: true,
//     signature: "我是你的 你是我的 谁",
//   },
//   {
//     name: "小摆",
//     uid: 4,
//     live_url:
//       "http://pili-live-hls.scooter.peterli.club/scooter/scooter-live.m3u8",
//     live_cover_url:
//       "http://s327crbzf.hn-bkt.clouddn.com/3d2be50e68424af8fe4f2b365884863864a774785fe38cddf405c4afb78e85fe.jpg",
//     avatar:
//       "http://s327crbzf.hn-bkt.clouddn.com/3d2be50e68424af8fe4f2b365884863864a774785fe38cddf405c4afb78e85fe.jpg",
//     is_follow: true,
//     signature: "我是你的 你是我的 谁",
//   },
// ];
import { getLiveList } from "@/apis/live";
const videos = ref<any>([]);

onMounted(() => {
  getLiveList().then((res:any)=>{
    videos.value = res.user_list;
  })
});

</script>
<template>
  <LiveCard :videos="videos" />
</template>

<style scoped></style>
