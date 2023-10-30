<!--
 * @Author: Xu Ning
 * @Date: 2023-10-26 20:05:39
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-30 18:39:43
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\view\ClassifiedView.vue
-->

<template>
  <VideoCard :is-scroll="true" :videos="videos" />
</template>

<script lang="ts" setup>
import VideoCard from "@/components/VideoCard.vue";
import { onMounted, ref, computed, watch } from 'vue'
import { VideoType } from "@/apis/interface";
import { getCategoryVideosList } from '@/apis/video'
import { routeStore } from '@/stores/route'
const videos = ref<Array<VideoType>>()
const videoType = computed(()=> routeStore().name)
const VideoTypeMap:any = {
    'hot':0, 
    'recreation':1, 
    'sports':2, 
    'food':3, 
    'cartoon':4, 
    'knowledge':5
}

// 挂载首次打开页面的视频流
onMounted(() => {
  let typeCode = VideoTypeMap[videoType.value]
  console.log(typeCode)
  if(typeCode == 0){

  }else{
    getCategoryVideosList(typeCode).then((res:any)=>{
      videos.value = res.videos
    })
  }
})

// 更新切换页面后的视频流数据
watch(()=>videoType.value,
(newValue:any)=>{
  console.log('videoType',newValue,VideoTypeMap['hot'])
  let typeCode = VideoTypeMap[videoType.value]
  getCategoryVideosList(typeCode).then((res:any)=>{
    videos.value = res.videos
  })
})
</script>
