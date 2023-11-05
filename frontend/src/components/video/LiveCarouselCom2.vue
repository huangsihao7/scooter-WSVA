<!--
 * @Author: huangsihao7
 * @Date: 2023-10-30 11:17:41
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-05 18:52:31
 * @FilePath: \scooter-WSVA\frontend\src\components\video\LiveCarouselCom.vue
 * @Description: 实现上下滑动的短视频组件
-->
<script setup lang="ts">
import { onMounted, ref } from "vue";
import {
  NCarousel,
  NDrawer,
  NDrawerContent,
  NCarouselItem,
  NTabs,
  NInput,
  NTabPane,
  NIcon,
  NButton,
  NEmpty,
} from "naive-ui";
import { ArrowUpCircle } from "@vicons/ionicons5";
import {
  getRecommendVideos,
  getPopularVideos,
  getRecommendVideosList,
  getVideosList,
} from "@/apis/video";
import { UserType, VideoType, CommentType } from "@/apis/interface";
import { getCommentList, doComment } from "@/apis/comment";
import { userStore } from "@/stores/user";
import { videoStore } from "@/stores/video";
import VideoPlus from "@/components/video/VideoPlus.vue";
import CommentListCom from "@/components/comment/CommentListCom.vue";
import VideoRecommendCard from "@/components/cards/VideoRecommendCard.vue";
import { throttle } from "lodash";
import { VideocamOff, ChatbubbleEllipses } from "@vicons/ionicons5";
// 评论区域是否可见
const drawerVisible = ref<boolean>(false);
// 正在播放视频的下标
const currentVideoIndex = ref<number>(0);
// 上一个播放过的视频的下标
const lastVideoIndex = ref<number>(0);
// 视频队列
const videos = ref<Array<VideoType>>();
// const videos = reactive<VideoType[]>([]);
// 默认预加载数
const defaultLoad: number = 4;
// 评论列表
const commentlists = ref<Array<CommentType>>([]);
// 相关推荐列表
const recommendlists = ref<any>();
// 绑定轮播器
const carouselRef = ref<any>();
// 看到第几个视频的标记
const visitedIndex = ref<number>(-1);

// 获取视频队列
onMounted(() => {
  if (userStore().isLoggedIn) {
    getRecommendVideos(0, 0).then((res: any) => {
      videos.value = res.video_list;
    });
    visitedIndex.value = 0;
  } else {
    // 游客试看功能 前10条
    getVideosList().then((res: any) => {
      if (res.status_code == 200) {
        let loadNum = res.video_list.length > 10 ? 10 : res.video_list.length;
        videos.value = res.video_list.slice(0, loadNum);
      }
    });
  }
});

// 更新评论区可见状态
const updateVisible = (video_id: any) => {
  drawerVisible.value = !drawerVisible.value;
  getCommentList(video_id).then((res: any) => {
    commentlists.value = res.comment_list;
  });
  getRecommendVideosList(video_id).then((res: any) => {
    recommendlists.value = res.video_list;
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
  if (currentIndex > lastIndex && currentIndex > visitedIndex.value) {
    let offset = defaultLoad + currentIndex;
    let readedVideo = videoStore().video_id;
    visitedIndex.value = currentIndex;
    getRecommendVideos(offset, readedVideo).then((res: any) => {
      if (res.video_list[0]) {
        videos.value?.push(res.video_list[0]);
      }
    });
  }
};

const throttledUpdatePage = throttle(updatePage, 200);

</script>

<template>
  <div>
    <NCarousel
      ref="carouselRef"
      :loop="false"
      :on-update:current-index="throttledUpdatePage"
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
</template>

<style scoped lang="scss"></style>
