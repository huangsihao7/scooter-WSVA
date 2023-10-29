<script setup lang="ts">
import { NCarousel, NDrawer, NDrawerContent, NCarouselItem } from "naive-ui";
import VideoPlus from "@/components/video/VideoPlus.vue";
import { onMounted, ref } from "vue";
import CommentListCom from "@/components/comment/CommentListCom.vue";
import { getVideosList } from "@/apis/video";

// 评论区域是否可见
const drawerVisible = ref<boolean>(false);
const currentVideoIndex = ref<number>(0);
const lastVideoIndex = ref<number>(0);
const videos = ref<any>();

onMounted(() => {
  getVideosList().then((res: any) => {
    videos.value = res.videos;
  });
});

// 更新评论区可见状态
const updateVisible = () => {
  drawerVisible.value = !drawerVisible.value;
};

const carouselRef = ref<any>();

const upPage = () => {
  carouselRef.value.prev();
};

const downPage = () => {
  carouselRef.value.next();
};

const updatePage = (currentIndex: number, lastIndex: number) => {
  console.log("hello", currentIndex, lastIndex);
  currentVideoIndex.value = currentIndex;
  lastVideoIndex.value = lastIndex;
};
</script>

<template>
  <div>
    <NCarousel
      id="drawer-target"
      ref="carouselRef"
      :loop="false"
      :on-update:current-index="updatePage"
      class="wide-carousel"
      direction="vertical"
      dot-placement="right"
      mousewheel
      show-arrow
      :show-dots="false"
      style="width: 100%; height: calc(100vh - 60px)"
      @keydown.arrow-up="upPage"
      @keydown.arrow-down="downPage"
    >
      <NCarouselItem v-for="(video, index) in videos">
        <VideoPlus
          :onplay="currentVideoIndex"
          :index="index"
          :video="video"
          @comment-visible-update="updateVisible"
        />
      </NCarouselItem>
    </NCarousel>
  </div>
  <NDrawer
    v-model:show="drawerVisible"
    :width="400"
    :height="200"
    placement="right"
    :trap-focus="false"
    :block-scroll="false"
    to="#drawer-target"
  >
    <NDrawerContent title="评论" :native-scrollbar="false">
      <CommentListCom />
    </NDrawerContent>
  </NDrawer>
</template>

<style scoped lang="scss">
.carousel-img {
  width: 100%;
  height: 240px;
  object-fit: cover;
}

.short-carousel {
  .video-com {
    width: auto;
    height: calc(100vh - 60px);
  }
}

.wide-carousel {
  .video-com {
    width: calc(100vw - 560px);
    height: calc(100vh - 60px);
  }
}

.n-carousel.n-carousel--right .n-carousel__arrow-group {
  bottom: 100px;
}
</style>
