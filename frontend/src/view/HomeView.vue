<!--
 * @Author: Xu Ning
 * @Date: 2023-05-08 15:29:52
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-31 19:30:50
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\view\HomeView.vue
-->
<script setup lang="ts">
import VideoList from "@/components/VideoList.vue";
import { routeStore } from "@/stores/route";
import { ref, onMounted} from 'vue'
// 视频队列类别 推荐0 热门1
const videoListType = ref<number>(0)
// 根据路由判断视频队列类别
onMounted(() => {
  if(routeStore().name == 'rec'){
    videoListType.value = 0
  }else{
    videoListType.value = 1
  }
})
</script>

<template>
  <VideoList :videoListType="videoListType"/>
</template>

<style scoped></style>
