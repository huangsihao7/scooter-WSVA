<!--
 * @Author: huangsihao7
 * @Date: 2023-10-30 11:17:41
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-05 22:37:27
 * @FilePath: \scooter-WSVA\frontend\src\components\live\LiveCarouselCom.vue
 * @Description: 实现上下滑动的短视频组件
-->
<script setup lang="ts">
import { reactive } from "vue";
import Hls from "hls.js";
import LiveCom from "@/components/live/LiveCom.vue";

interface propsType{
  userId: number,
  url: string
}
const props = defineProps<propsType>()

const dplayerObj = reactive({
  live: true,
  videoId: 0,
  video: {
    url: 'http://pili-live-hls.scooter.peterli.club/scooter/scooter-live.m3u8', //视频地址
    type: "customHls",
    customType: {
      customHls: function (video: any) {
        const hls = new Hls(); //实例化Hls  用于解析m3u8
        hls.loadSource(video.src);
        hls.attachMedia(video);
      },
    },
  },
 
});
</script>

<template>
  <div>
    <LiveCom
        :video-id="dplayerObj.videoId"
        :video="dplayerObj.video"
        :live = "dplayerObj.live"
      />
  </div>
</template>

<style scoped lang="scss">

.dplayer {
    width: calc(100vw - 160px);
    height: calc(100vh - 60px);
  }
  </style>
