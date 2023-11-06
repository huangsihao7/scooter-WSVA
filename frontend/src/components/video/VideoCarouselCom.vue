<!--
 * @Author: huangsihao7
 * @Date: 2023-10-30 11:17:41
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-06 20:18:41
 * @FilePath: \scooter-WSVA\frontend\src\components\video\VideoCarouselCom.vue
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
const defaultLoad: number = 2;
// 评论列表
const commentlists = ref<Array<CommentType>>([]);
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
  if (userStore().isLoggedIn) {
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
};

const throttledUpdatePage = throttle(updatePage, 200);

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

// 发布评论接口
const doCommentApi = () => {
  doComment(videoStore().video_id, 1, addComment.value, 0).then((res: any) => {
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
      comment_id: res.comment_id,
    };
    commentlists.value?.push(addCommentObj);
    addComment.value = "";
  });
};
const throttledDoComment = throttle(doCommentApi, 500);
// 发布评论
const postComment = (e: any) => {
  if (e.keyCode == 13 && addComment.value) {
    throttledDoComment();
  }
};

// 点击事件发布评论
const postCommentByBtn = () => {
  throttledDoComment();
};

// 动态删除评论数据
const deleteFunc = (comment_id: number) => {
  commentlists.value = commentlists.value?.filter(
    (item: CommentType) => item.comment_id !== comment_id,
  );
};

const tabValue = ref<string>("comment");
</script>

<template>
  <div>
    <NCarousel
      id="drawer-target"
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
  <NDrawer
    v-model:show="drawerVisible"
    :width="400"
    placement="right"
    :trap-focus="false"
    :block-scroll="false"
    to="#drawer-target"
  >
    <NDrawerContent :native-scrollbar="false">
      <NTabs v-model:value="tabValue" type="line" animated>
        <NTabPane name="comment" tab="评论">
          <CommentListCom
            v-if="commentlists.length != 0"
            :commentlists="commentlists"
            @delete-comment="deleteFunc"
          />
          <NEmpty v-else description="还没有评论哦，快来抢沙发~">
            <template #icon>
              <NIcon>
                <ChatbubbleEllipses />
              </NIcon>
            </template>
          </NEmpty>
        </NTabPane>
        <NTabPane name="recommend" tab="相关推荐">
          <VideoRecommendCard
            v-if="recommendlists.length != 0"
            :recommendlists="recommendlists"
          />
          <NEmpty v-else description="没有推荐的视频哦~去别处看看吧~">
            <template #icon>
              <NIcon>
                <VideocamOff />
              </NIcon>
            </template>
          </NEmpty>
        </NTabPane>
      </NTabs>
      <template v-if="tabValue == 'comment'" #footer>
        <NInput
          v-model:value="addComment"
          maxlength="30"
          show-count
          round
          class="comment-input"
          placeholder="留下精彩的评论吧"
          @keydown="postComment"
        >
          <template #prefix>
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

<style scoped lang="scss"></style>
