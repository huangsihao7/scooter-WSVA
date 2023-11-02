<!--
 * @Author: Xu Ning
 * @Date: 2023-10-27 14:13:32
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-02 14:47:13
 * @Description: 各个视频分类的视频卡片渲染
 * @FilePath: \scooter-WSVA\frontend\src\components\VideoCard.vue
-->
<script lang="ts" setup>
import { onMounted } from "vue";
import { NScrollbar, NSpace } from "naive-ui";
import VideoCardList from "@/components/video/VideoCardList.vue";
interface propsType {
  isScroll: boolean;
  videos: any;
  deletable?: boolean;
}
const props = defineProps<propsType>();

onMounted(() => {
  //info.content
});
</script>

<template>
  <NScrollbar v-if="props.isScroll" class="card-space">
    <NSpace wrap>
      <VideoCardList :videos="props.videos" />
    </NSpace>
  </NScrollbar>
  <NSpace v-else wrap>
    <VideoCardList :videos="props.videos" :deletable="$props.deletable" />
  </NSpace>
</template>

<style lang="scss" scoped>

.card-space {
  overflow: scroll-y;
  max-height: calc(100vh - 60px);
}

</style>
