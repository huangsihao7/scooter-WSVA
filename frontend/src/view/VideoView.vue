<!--
 * @Author: Xu Ning
 * @Date: 2023-10-31 18:42:57
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-02 21:03:55
 * @Description: 查看某个特定video
 * @FilePath: \scooter-WSVA\frontend\src\view\VideoView.vue
-->
<script setup lang="ts">
import { ref, onMounted, computed } from "vue";
import VideoPlus from "@/components/video/VideoPlus.vue";
import {
  NEmpty,
  NDrawer,
  NDrawerContent,
  NTabs,
  NTabPane,
  NInput,
  NIcon,
} from "naive-ui";
import { getVideoById } from "@/apis/video";
import { useRoute } from "vue-router";
import { getCommentList, doComment } from "@/apis/comment";
import { getRecommendVideosList } from "@/apis/video";
import { videoStore } from "@/stores/video";
import { CommentType } from "@/apis/interface";
import { userStore } from "@/stores/user";
import { UserType } from "@/apis/interface";
import CommentListCom from "@/components/comment/CommentListCom.vue";
import VideoRecommendCard from "@/components/video/VideoRecommendCard.vue";
import { ColorWand } from "@vicons/ionicons5";

// 评论区域是否可见
const drawerVisible = ref<boolean>(false);
const route = useRoute();
const videoId = computed(() => route.params.id);
const video = ref<any>();
const commentlists = ref<any>();
// 相关推荐列表
const recommendlists = ref<any>();
// 添加评论的内容
const addComment = ref<string>("");

// 获取视频
onMounted(() => {
  let vid = videoId.value.toString();
  let vidNum = parseInt(vid);
  getVideoById(vidNum).then((res: any) => {
      video.value = res.video_info;
  });
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

// 发布评论
const postComment = (e: any) => {
  if (e.keyCode == 13 && addComment.value) {
    doComment(videoStore().video_id, 1, addComment.value, 0).then(
      (res: any) => {
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
        }
    );
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
</script>

<template>
  <div class="video">
    <VideoPlus
      v-if="video"
      id="drawer-target"
      :video="video"
      :index="0"
      :onplay="0"
      @comment-visible-update="updateVisible"
    />
    <NEmpty
      v-else
      size="huge"
      class="empty-state"
      description="你找的视频下架啦！！！"
    >
    </NEmpty>
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
            <CommentListCom :commentlists="commentlists" />
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
            >
            <template #suffix>
              <NIcon :component="ColorWand" />
            </template>
          </NInput>
        </template>
      </NDrawerContent>
    </NDrawer>
  </div>
</template>

<style scoped lang="scss">
.video {
  height: calc(100vh - 60px);
  width: 100%;
  .empty-state {
    height: 100%;
    text-align: center;
    top: 45%;
    left: 45%;
    position: fixed;
  }
}
</style>
