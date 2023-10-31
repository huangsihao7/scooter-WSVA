<!--
 * @Author: Xu Ning
 * @Date: 2023-10-31 18:42:57
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-31 19:28:02
 * @Description: 查看某个特定video
 * @FilePath: \scooter-WSVA\frontend\src\view\VideoView.vue
-->
<script setup lang="ts">
import { ref, onMounted } from "vue";
import { VideoType } from "@/apis/interface";
import VideoPlus from "@/components/video/VideoPlus.vue";
import { NEmpty } from "naive-ui";
const video = ref<VideoType>()

onMounted(() => {
//   getvideol  
})

</script>

<template>
    <div class="video">
        <VideoPlus v-if="video" :video="video" :index=0 :onplay=0 />
        <n-empty size="huge" class="empty-state" v-else description="你找的视频下架啦！！！">
        </n-empty>
    </div>
</template>

<style scoped lang="scss">
.video{
    height: calc(100vh - 60px);
    width: 100%;
    .empty-state{
        height: 100%;
        text-align: center;
        top: 45%;
        left: 45%;
        position: fixed;
    }
}
</style>
