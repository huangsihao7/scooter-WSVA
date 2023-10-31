<script setup lang="ts">
import { NCarousel, NDrawer, NDrawerContent, NCarouselItem } from "naive-ui";
import VideoPlus from "@/components/video/VideoPlus.vue";
import { onMounted, ref } from "vue";
import CommentListCom from "@/components/comment/CommentListCom.vue";
import { getRecommendVideos, getPopularVideos } from "@/apis/video";
import { VideoType } from "@/apis/interface";
import { getCommentList } from "@/apis/comment";
import { videoStore } from "@/stores/video";

interface propsType{
  videoListType: number
}

const props = defineProps<propsType>()
// 评论区域是否可见
const drawerVisible = ref<boolean>(false);
// 正在播放视频的下标
const currentVideoIndex = ref<number>(0);
// 上一个播放过的视频的下标
const lastVideoIndex = ref<number>(0);
// 视频队列
const videos = ref<Array<VideoType>>();
// 默认预加载数
const defaultLoad: number = 4;
// 评论列表
const commentlists = ref<any>();
// 绑定轮播器
const carouselRef = ref<any>();

// 获取视频队列
onMounted(() => {
  if(props.videoListType == 0){
    getRecommendVideos(0, 0).then((res: any) => {
      videos.value = res.video_list;
    });
  }else{
    getPopularVideos(0, 0).then((res: any) => {
      videos.value = res.video_list;
    });
  }
});

// 更新评论区可见状态
const updateVisible = (thisVideo: any) => {
  drawerVisible.value = !drawerVisible.value;
  console.log(thisVideo.value.video_id);
  getCommentList(thisVideo.value.video_id).then((res: any) => {
    commentlists.value = res.comment_list;
  });
};

// 向上翻视频
const upPage = () => {
  carouselRef.value.prev();
};

// 向下翻视频
const downPage = () => {
  carouselRef.value.next();
};

// 轮播图切换效果
const updatePage = (currentIndex: number, lastIndex: number) => {
  currentVideoIndex.value = currentIndex;
  lastVideoIndex.value = lastIndex;
  if (currentIndex > lastIndex) {
    let offset = defaultLoad + currentIndex;
    let readedVideo = videoStore().video_id;
    console.log("readedVideo", readedVideo);
    if(props.videoListType == 0){
    getRecommendVideos(offset, readedVideo).then((res: any) => {
      videos.value?.push(res.video_list[0]);
    });
    }else{
      getPopularVideos(offset, readedVideo).then((res: any) => {
        videos.value?.push(res.video_list[0]);
      });
    }
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
