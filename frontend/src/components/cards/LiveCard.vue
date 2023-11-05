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
</script>

<template>
  <NScrollbar v-if="props.videos.length != 0" class="card-space">
    <!-- <NScrollbar v-if="videos.length != 0" class="card-space"> -->
    <NSpace wrap>
      <LiveCardListCom :videos="props.videos" />
    </NSpace>
  </NScrollbar>
  <NEmpty
    v-else
    class="empty-space"
    description="没有用户在直播哦~去别处看看吧~"
  >
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

.empty-space {
  height: calc(100vh - 60px);
  margin: 40vh;
}
</style>
