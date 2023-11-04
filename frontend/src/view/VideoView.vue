<!--
 * @Author: Xu Ning
 * @Date: 2023-10-31 18:42:57
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-04 17:48:04
 * @Description: 查看某个特定video的页面
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
  NButton,
  useMessage,
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
import VideoRecommendCard from "@/components/cards/VideoRecommendCard.vue";
import { ArrowUpCircle } from "@vicons/ionicons5";
import { VideocamOff, ChatbubbleEllipses } from "@vicons/ionicons5";

// 评论区域是否可见
const drawerVisible = ref<boolean>(false);
const route = useRoute();
const videoId = computed(() => route.params.id);
const video = ref<any>();
const commentlists = ref<Array<CommentType>>([]);
// 相关推荐列表
const recommendlists = ref<any>();
// 添加评论的内容
const addComment = ref<string>("");
const message = useMessage();
const tabValue = ref<string>("comment");

// 获取视频
onMounted(() => {
  let vid = videoId.value.toString();
  let vidNum = parseInt(vid);
  getVideoById(vidNum).then((res: any) => {
    video.value = res.video_info;
  });
});

// 更新评论区可见状态
const updateVisible = (video_id: number) => {
  drawerVisible.value = !drawerVisible.value;
  getCommentList(video_id).then((res: any) => {
    commentlists.value = res.comment_list;
  });
  getRecommendVideosList(video_id).then((res: any) => {
    recommendlists.value = res.video_list;
  });
};

// 发布评论接口
const doCommentApi = () => {
  if (addComment.value != "") {
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
      },
    );
  } else {
    message.error("评论不能为空");
  }
};

// 发布评论
const postComment = (e: any) => {
  if (e.keyCode == 13 && addComment.value) {
    doCommentApi();
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

// 动态删除评论数据
const deleteFunc = (comment_id: number) => {
  commentlists.value = commentlists.value?.filter(
    (item: CommentType) => item.comment_id !== comment_id,
  );
};

// 点击事件发布评论
const postCommentByBtn = () => {
  doCommentApi();
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
