<!--
 * @Author: Xu Ning
 * @Date: 2023-10-27 14:13:32
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-05 21:06:02
 * @Description: 各个视频分类的视频卡片列表渲染
 * @FilePath: \scooter-WSVA\frontend\src\components\cards\VideoCard.vue
-->
<script lang="ts" setup>
import { NScrollbar, NSpace } from "naive-ui";
import VideoCardListCom from "@/components/video/VideoCardListCom.vue";
interface propsType {
  // 用户视频卡片
  isScroll: boolean;
  videos: any;
  // 分类视频卡片
  deletable?: boolean;
}
const props = defineProps<propsType>();
</script>

<template>
  <NScrollbar v-if="props.isScroll" class="card-space">
    <NSpace wrap>
      <VideoCardListCom :videos="props.videos" />
    </NSpace>
  </NScrollbar>
  <NSpace v-else wrap>
    <VideoCardListCom :videos="props.videos" :deletable="$props.deletable" />
  </NSpace>
</template>

<style lang="scss">
.card-space {
  overflow: scroll-y;
  max-height: calc(100vh - 60px);
}
.n-space{
  text-align: left;
}
</style>
