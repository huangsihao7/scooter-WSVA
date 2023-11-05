<!--
 * @Author: Xu Ning
 * @Date: 2023-10-27 14:13:32
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-05 21:59:07
 * @Description: 直播卡片列表渲染
 * @FilePath: \scooter-WSVA\frontend\src\components\cards\LiveCard.vue
-->
<script lang="ts" setup>
import { NEmpty, NScrollbar, NSpace, NIcon } from "naive-ui";
import LiveCardListCom from "@/components/live/LiveCardListCom.vue";
import { VideocamOff } from "@vicons/ionicons5";
interface propsType {
  videos: any;
}
const props = defineProps<propsType>();
const video_list = [{
  name:'小红',
  id:1,
  live_play_url:'http://pili-live-hls.scooter.peterli.club/scooter/scooter-live.m3u8',
  live_cover_url:"http://s327crbzf.hn-bkt.clouddn.com/3d2be50e68424af8fe4f2b365884863864a774785fe38cddf405c4afb78e85fe.jpg",
  avatar:"http://s327crbzf.hn-bkt.clouddn.com/3d2be50e68424af8fe4f2b365884863864a774785fe38cddf405c4afb78e85fe.jpg",
  is_follow:true,
  signature:'我是你的 你是我的 谁'
},{
  name:'小黑',
  id:2,
  live_play_url:'http://pili-live-hls.scooter.peterli.club/scooter/scooter-live.m3u8',
  live_cover_url:"http://s327crbzf.hn-bkt.clouddn.com/3d2be50e68424af8fe4f2b365884863864a774785fe38cddf405c4afb78e85fe.jpg",
  avatar:"http://s327crbzf.hn-bkt.clouddn.com/3d2be50e68424af8fe4f2b365884863864a774785fe38cddf405c4afb78e85fe.jpg",
  is_follow:true,
  signature:'我是你的 你是我的kdjashkljghwake 谁'
},{
  name:'小栏',
  id:3,
  live_play_url:'http://pili-live-hls.scooter.peterli.club/scooter/scooter-live.m3u8',
  live_cover_url:"http://s327crbzf.hn-bkt.clouddn.com/3d2be50e68424af8fe4f2b365884863864a774785fe38cddf405c4afb78e85fe.jpg",
  avatar:"http://s327crbzf.hn-bkt.clouddn.com/3d2be50e68424af8fe4f2b365884863864a774785fe38cddf405c4afb78e85fe.jpg",
  is_follow:true,
  signature:'我是你的 你是我的 谁'
},{
  name:'小摆',
  id:4,
  live_play_url:'http://pili-live-hls.scooter.peterli.club/scooter/scooter-live.m3u8',
  live_cover_url:"http://s327crbzf.hn-bkt.clouddn.com/3d2be50e68424af8fe4f2b365884863864a774785fe38cddf405c4afb78e85fe.jpg",
  avatar:"http://s327crbzf.hn-bkt.clouddn.com/3d2be50e68424af8fe4f2b365884863864a774785fe38cddf405c4afb78e85fe.jpg",
  is_follow:true,
  signature:'我是你的 你是我的 谁'
},
]
</script>

<template>
  <NScrollbar v-if="video_list.length != 0" class="card-space">
  <!-- <NScrollbar v-if="videos.length != 0" class="card-space"> -->
    <NSpace wrap>
      <LiveCardListCom  :videos="video_list" />
    </NSpace>
  </NScrollbar>
  <NEmpty class="empty-space" v-else description="没有用户在直播哦~去别处看看吧~">
    <template #icon>
      <NIcon>
        <VideocamOff />
      </NIcon>
    </template>
  </NEmpty>
</template>

<style lang="scss" scoped>
.card-space {
  overflow: scroll-y;
  max-height: calc(100vh - 60px);
}

.empty-space{
  height: calc(100vh - 60px);;
  margin: 40vh;
}
</style>
