<script setup lang="ts">
import { NCarousel, NDrawer, NDrawerContent } from 'naive-ui';
import VideoPlus from './VideoPlus.vue'
import { onMounted, ref } from 'vue'
import Comment from '@/components/comment/Comment.vue'

// 评论区域是否可见
const drawerVisible = ref<boolean>(true)

// 动态变化轮播视频类
const carouselClass = ref<string>('wide-carousel')

onMounted(() => {
    
})

// 更新评论区可见状态
const updateVisible = () =>{
    drawerVisible.value = !drawerVisible.value
}


</script>

<template>
    <div  >
        <n-carousel
            class= 'wide-carousel'
            id = "drawer-target"
            direction="vertical"
            dot-placement="right"
            mousewheel
            show-arrow
            :show-dots="false"
            style="width: 100%; height: calc(100vh - 60px);"
        >
            <VideoPlus @comment-visible-update="updateVisible"/>
            <VideoPlus @comment-visible-update="updateVisible"/>
            <VideoPlus @comment-visible-update="updateVisible"/>
            <VideoPlus @comment-visible-update="updateVisible"/>
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
            <Comment/>
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