<!--
 * @Author: Xu Ning
 * @Date: 2023-11-05 18:42:02
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-05 22:59:48
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\view\LiveView.vue
-->
<script setup lang="ts">
import { onMounted, computed } from "vue";
import LivePlus from "@/components/live/LivePlus.vue";
import { useRoute } from "vue-router";
import { liveStore } from "@/stores/live";
const route = useRoute();
// 播放视频号
const liveId = computed(() =>
  route.params.id ? parseInt(route.params.id.toString()) : liveStore().live_id,
);
// 播放路由
const url = computed(() =>
  route.params.url ? route.params.url.toString() : liveStore().live_play_url,
);

onMounted(() => {
  //TODO getliveinfo
});
</script>
<template>
  <LivePlus :live-id="liveId" :url="url" />
</template>

<style scoped></style>
