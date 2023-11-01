<!--
 * @Author: Xu Ning
 * @Date: 2023-10-31 18:42:57
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-01 12:11:23
 * @Description: 查看某个特定video
 * @FilePath: \scooter-WSVA\frontend\src\view\VideoView.vue
-->
<script setup lang="ts">
import { ref, onMounted, computed } from "vue";
import VideoPlus from "@/components/video/VideoPlus.vue";
import { NEmpty, NDrawer, NDrawerContent, NTabs, NTabPane, NInput } from "naive-ui";
import { getVideoById } from "@/apis/video";
import { useRoute } from "vue-router";
import { getCommentList } from "@/apis/comment";
import { getRecommendVideosList } from "@/apis/video";
import CommentListCom from "@/components/comment/CommentListCom.vue";
import VideoRecommendCard from "@/components/video/VideoRecommendCard.vue";

// 评论区域是否可见
const drawerVisible = ref<boolean>(false);
const route = useRoute();
const videoId = computed(() => route.params.id);
const video = ref<any>();
const commentlists = ref<any>();
// 相关推荐列表
const recommendlists = ref<any>();

onMounted(() => {
  let vid = videoId.value.toString();
  let vidNum = parseInt(vid);
  getVideoById(vidNum).then((res: any) => {
    video.value = res.video_info;
    console.log("video", res, res.video_info, video.value);
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
            <NInput class="comment-input" round placeholder="中" />
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
