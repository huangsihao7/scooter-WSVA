<script setup lang="ts">
import { NCarousel, NDrawer, NDrawerContent, NCarouselItem } from 'naive-ui';
import VideoPlus from './VideoPlus.vue'
import { onMounted, ref } from 'vue'
import CommentListCom from '@/components/comment/CommentListCom.vue'
import { getVideosList } from '@/apis/video'
import { keys } from 'lodash';


// 评论区域是否可见
const drawerVisible = ref<boolean>(false)

// 动态变化轮播视频类
const carouselClass = ref<string>('wide-carousel')

const videos = ref<any>()

onMounted(() => {
    getVideosList().then((res:any)=>{
        videos.value = res.videos
    })
})

// 更新评论区可见状态
const updateVisible = () =>{
    drawerVisible.value = !drawerVisible.value
}

const carouselRef = ref<any>()

const upPage = () =>{
    carouselRef.value.prev()
}

const downPage = () =>{
    carouselRef.value.next()
}

</script>

<template>
    <div>
        <n-carousel
            ref="carouselRef"
            @keydown.arrow-up="upPage"
            @keydown.arrow-down="downPage"
            class= 'wide-carousel'
            id = "drawer-target"
            direction="vertical"
            dot-placement="right"
            mousewheel
            show-arrow
            :show-dots="false"
            style="width: 100%; height: calc(100vh - 60px);"
        >
            <n-carousel-item v-for="(video,index) in videos"> <VideoPlus :key="index" :video="video" @comment-visible-update="updateVisible"/></n-carousel-item>
            
        </n-carousel>
    </div>
    <n-drawer
        v-model:show="drawerVisible"
        :width="400"
        :height="200"
        placement="right"
        :trap-focus="false"
        :block-scroll="false"
        to="#drawer-target"
    >
    
        <n-drawer-content title="评论" :native-scrollbar="false">
            <CommentListCom />
        </n-drawer-content>
  </n-drawer>
</template>

<style scoped lang="scss">
.carousel-img {
width: 100%;
height: 240px;
object-fit: cover;
}

.short-carousel{
    .video-com{
        width: auto;
        height: calc(100vh - 60px);
    }
}

.wide-carousel{
    .video-com{
        width: calc(100vw - 560px);
        height: calc(100vh - 60px);
    }
}

.n-carousel.n-carousel--right .n-carousel__arrow-group{
    bottom: 100px;
}
</style>