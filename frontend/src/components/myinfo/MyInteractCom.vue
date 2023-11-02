<!--
 * @Author: Xu Ning
 * @Date: 2023-10-28 12:30:41
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-02 14:50:21
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
      <NScrollbar style="max-height: 50vh">
        <VideoCard :is-scroll="false" :videos="videos" :deletable="deletable" />
      </NScrollbar>
    </NTabPane>
    <NTabPane name="favourite" tab="喜欢">
      <NScrollbar style="max-height: 50vh">
        <VideoCard :is-scroll="false" :videos="videos" />
      </NScrollbar>
    </NTabPane>
    <NTabPane name="collect" tab="收藏">
      <NScrollbar style="max-height: 50vh">
        <VideoCard :is-scroll="false" :videos="videos" />
      </NScrollbar>
    </NTabPane>
    <NTabPane name="history" tab="观看历史">
      <NScrollbar style="max-height: 50vh">
        <VideoCard :is-scroll="false" :videos="videos" />
      </NScrollbar>
    </NTabPane>
  </NTabs>
</template>
<script lang="ts" setup>
import { NScrollbar, NTabPane, NTabs } from "naive-ui";
import VideoCard from "../VideoCard.vue";
import { getHistoryVideosListReq, userVideoListReq } from "@/apis/video";
import { userFavouriteListReq, userStarListReq } from "@/apis/favourite";
import { onMounted, ref, computed } from "vue";
import { userStore } from "@/stores/user";

const videos = ref<Array<any>>([]);

interface propsType {
  userId: number;
}

const deletable = computed(()=>(props.userId==userStore().user_id)?true:false)
const props = defineProps<propsType>();
const myid = ref<number>(-1);

onMounted(() => {
  getMyWork();
});

const getMyWork = () => {
  let uid = props.userId.toString();
  let uid_num = parseInt(uid);
  myid.value = uid_num;
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

</style>
