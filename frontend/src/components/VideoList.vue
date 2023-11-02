<script setup lang="ts">
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
} from "naive-ui";
import VideoPlus from "@/components/video/VideoPlus.vue";
import { onMounted, ref } from "vue";
import CommentListCom from "@/components/comment/CommentListCom.vue";
import {
  getRecommendVideos,
  getPopularVideos,
  getRecommendVideosList,
} from "@/apis/video";
import { UserType, VideoType } from "@/apis/interface";
import { getCommentList, doComment } from "@/apis/comment";
import { videoStore } from "@/stores/video";
import VideoRecommendCard from "@/components/video/VideoRecommendCard.vue";
import { CommentType } from "@/apis/interface";
import { userStore } from "@/stores/user";
import { ArrowUpCircle } from "@vicons/ionicons5";

interface propsType {
  videoListType: number;
}

const props = defineProps<propsType>();
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
const commentlists = ref<Array<CommentType>>();
// 相关推荐列表
const recommendlists = ref<any>();
// 绑定轮播器
const carouselRef = ref<any>();
// 添加评论的内容
const addComment = ref<string>("");
// 看到第几个视频的标记
const visitedIndex = ref<number>(-1);

// 获取视频队列
onMounted(() => {
  if (props.videoListType == 0) {
    getRecommendVideos(0, 0).then((res: any) => {
      videos.value = res.video_list;
    });
  } else {
    getPopularVideos(0, 0).then((res: any) => {
      videos.value = res.video_list;
    });
  }
  visitedIndex.value = 0;
  console.log(visitedIndex.value);
});

// 更新评论区可见状态
const updateVisible = (thisVideo: any) => {
  drawerVisible.value = !drawerVisible.value;
  getCommentList(thisVideo.value.video_id).then((res: any) => {
    commentlists.value = res.comment_list;
  });
  getRecommendVideosList(thisVideo.value.video_id).then((res: any) => {
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
  console.log(visitedIndex.value);
  if (currentIndex > lastIndex && currentIndex > visitedIndex.value) {
    let offset = defaultLoad + currentIndex;
    let readedVideo = videoStore().video_id;
    visitedIndex.value = currentIndex;
    if (props.videoListType == 0) {
      getRecommendVideos(offset, readedVideo).then((res: any) => {
        if (res.video_list[0]) {
          videos.value?.push(res.video_list[0]);
        }
      });
    } else {
      getPopularVideos(offset, readedVideo).then((res: any) => {
        if (res.video_list[0]) {
          videos.value?.push(res.video_list[0]);
        }
      });
    }
    if (videos.value) {
      getCommentList(videos.value[currentIndex].video_id).then((res: any) => {
        if (res.comment_list[0]) {
          commentlists.value = res.comment_list;
        }
      });
      getRecommendVideosList(videos.value[currentIndex].video_id).then(
        (res: any) => {
          if (res.video_list[0]) {
            recommendlists.value = res.video_list;
          }
        },
      );
    }
  }
};

// 时间获取
const formattedDate = () => {
  const currentDate = new Date();
  const year = currentDate.getFullYear();
  const month = String(currentDate.getMonth() + 1).padStart(2, "0");
  const day = String(currentDate.getDate()).padStart(2, "0");
  const hours = String(currentDate.getHours()).padStart(2, "0");
  const minutes = String(currentDate.getMinutes()).padStart(2, "0");
  const formattedDate = `${year}-${month}-${day} ${hours}:${minutes}`;
  return formattedDate;
};

const doCommentApi = () => {
  doComment(videoStore().video_id, 1, addComment.value, 0).then((res: any) => {
    if (res.status_code == 200) {
      let userInfo = userStore();
      let userObj: UserType = {
        id: userInfo.user_id,
        name: userInfo.name,
        gender: userInfo.gender,
        avatar: userInfo.avatar,
        signature: userInfo.signature,
        phoneNum: userInfo.phoneNum,
        background_image: userInfo.background_image,
        follow_count: 0,
        follower_count: 0,
        total_favorited: 0,
        work_count: 0,
        favorite_count: 0,
        is_follow: false,
      };
      let addCommentObj: CommentType = {
        content: addComment.value,
        create_date: formattedDate(),
        user: userObj,
        comment_id: 11,
      };
      commentlists.value?.push(addCommentObj);
      addComment.value = "";
    }
  });
};

// 发布评论
const postComment = (e: any) => {
  if (e.keyCode == 13 && addComment.value) {
    doCommentApi();
  }
};

const postCommentByBtn = () => {
  doCommentApi();
};

// 动态删除评论数据
const deleteFunc = (comment_id: number) => {
  console.log("before", commentlists.value);
  commentlists.value = commentlists.value?.filter(
    (item: CommentType) => item.comment_id !== comment_id,
  );
  console.log("after", commentlists.value);
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
    placement="right"
    :trap-focus="false"
    :block-scroll="false"
    to="#drawer-target"
  >
    <NDrawerContent :native-scrollbar="false">
      <NTabs type="line" animated>
        <NTabPane name="comment" tab="评论">
          <CommentListCom
            v-if="commentlists"
            :commentlists="commentlists"
            @delete-comment="deleteFunc"
          />
        </NTabPane>
        <NTabPane name="recommend" tab="相关推荐">
          <VideoRecommendCard :recommendlists="recommendlists" />
        </NTabPane>
      </NTabs>
      <template #footer>
        <NInput
          v-model:value="addComment"
          class="comment-input"
          round
          placeholder="留下精彩的评论吧"
          @keydown="postComment"
        >
          <template #suffix>
            <NButton text @click="postCommentByBtn">
              <template #icon>
                <NIcon :component="ArrowUpCircle" />
              </template>
            </NButton>
          </template>
        </NInput>
      </template>
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

.n-carousel__arrow-group {
  bottom: 50px;
}
.n-carousel.n-carousel--right .n-carousel__arrow-group {
  bottom: 100px;
}
</style>
