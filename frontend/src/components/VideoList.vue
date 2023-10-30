<script setup lang="ts">
import { NCarousel, NDrawer, NDrawerContent, NCarouselItem } from "naive-ui";
import VideoPlus from "@/components/video/VideoPlus.vue";
import { onMounted, ref } from "vue";
import CommentListCom from "@/components/comment/CommentListCom.vue";
import { getRecommendVideos } from "@/apis/video";
import { VideoType } from "@/apis/interface";
import { getCommentList } from "@/apis/comment";

// 评论区域是否可见
const drawerVisible = ref<boolean>(false);
const currentVideoIndex = ref<number>(0);
const lastVideoIndex = ref<number>(0);
const videos = ref<Array<VideoType>>();
const defaultLoad: number = 4;
const commentlists = ref<any>();

onMounted(() => {
  getRecommendVideos(0).then((res: any) => {
    videos.value = res.videos;
  });
});

// 更新评论区可见状态
const updateVisible = (thisVideo: any) => {
  drawerVisible.value = !drawerVisible.value;
  console.log(thisVideo.value.id)
  getCommentList(thisVideo.value.id).then((res:any)=>{
    commentlists.value = res.comment_list
  })
}

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
  if (currentIndex > lastIndex) {
    let offset = defaultLoad + currentIndex;
    getRecommendVideos(offset).then((res: any) => {
      videos.value?.push(res.videos[0]);
      console.log(videos.value);
    });
  }
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
      <NCarouselItem v-for="(video, index) in videos" :key="index">
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
      <CommentListCom :commentlists="commentlists" />
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
