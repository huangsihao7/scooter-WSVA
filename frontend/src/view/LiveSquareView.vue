<!--
 * @Author: Xu Ning
 * @Date: 2023-11-05 18:42:02
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-06 16:56:04
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\view\LiveSquareView.vue
-->
<script setup lang="ts">
import { ref, onMounted } from "vue";
import LiveCard from "@/components/cards/LiveCard.vue";
import { getLiveList } from "@/apis/live";
const videos = ref<any>([]);

onMounted(() => {
  // videos.value = video_list
  getLiveList().then((res: any) => {
    videos.value = res.user_list;
  });
});
</script>
<template>
  <LiveCard :videos="videos" />
</template>

<style scoped></style>
