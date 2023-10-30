<!--
 * @Author: Xu Ning
 * @Date: 2023-10-28 12:30:41
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-30 22:58:33
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\components\myinfo\MyInteractCom.vue
-->
<template>
  
    <NTabs
      default-value="work"
      justify-content="space-evenly"
      type="line"
      @update-value="handleUpdate"
    >
      <NTabPane name="work" tab="作品">
        <n-scrollbar style="max-height: 300px">
          <VideoCard :is-scroll="false" :videos="videos" />
        </n-scrollbar>
      </NTabPane>
      <NTabPane name="favourite" tab="喜欢">
        <VideoCard :is-scroll="false" :videos="videos" />
      </NTabPane>
      <NTabPane name="collect" tab="收藏">
        <VideoCard :is-scroll="false" :videos="videos" />
      </NTabPane>
      <NTabPane name="history" tab="观看历史">
        <n-scrollbar style="max-height: 50vh">
          <VideoCard :is-scroll="false" :videos="videos" />
        </n-scrollbar>
      </NTabPane>
    </NTabs>
</template>
<script lang="ts" setup>
import { NTabs, NTabPane, NScrollbar } from "naive-ui";
import VideoCard from "../VideoCard.vue";
// import { userStore } from "@/stores/user";
import { userVideoListReq, getHistoryVideosListReq } from "@/apis/video";
import { userFavouriteListReq, userStarListReq } from "@/apis/favourite";
import { onMounted, ref } from "vue";
const videos = ref<any>();
interface propsType {
  userId: number;
}

const props = defineProps<propsType>();
const myid = ref<number>(-1)
onMounted(() => {
  getMyWork();
});
const getMyWork = () => {
  let uid = props.userId.toString()
  let uid_num = parseInt(uid)
  myid.value = uid_num
  userVideoListReq(myid.value).then((res: any) => {
    videos.value = res.video_list;
  });
};
const getMyFavourite = () => {
  userFavouriteListReq(myid.value).then((res: any) => {
    videos.value = res.video_list;
  });
};
const getMyCollect = () => {
  userStarListReq(myid.value).then((res: any) => {
    videos.value = res.video_list;
  });
};
const getMyHistory = () => {
  getHistoryVideosListReq().then((res: any) => {
    videos.value = res.video_list;
  });
};
const handleUpdate = (paneName: string) => {
  if (paneName === "work") {
    getMyWork();
  } else if (paneName === "favourite") {
    getMyFavourite();
  } else if (paneName === "collect") {
    getMyCollect();
  } else if (paneName === "history") {
    getMyHistory();
  }
};
</script>
<style>
.demo-tabs > .el-tabs__content {
  padding: 32px;
  color: #6b778c;
  font-size: 32px;
  font-weight: 600;
}
</style>
