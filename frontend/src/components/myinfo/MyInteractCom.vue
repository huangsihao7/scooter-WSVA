<!--
 * @Author: Xu Ning
 * @Date: 2023-10-28 12:30:41
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-03 19:13:22
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
      <NScrollbar  v-if="!showEmpty" style="max-height: 50vh">
        <VideoCard :is-scroll="false" :videos="videos" :deletable="deletable" />
      </NScrollbar>
      <NEmpty v-else description="没有发布过视频哦~">
        <template #icon>
          <NIcon>
            <VideocamOff />
          </NIcon>
        </template>
      </NEmpty>
    </NTabPane>
    <NTabPane name="favourite" tab="喜欢">
      <NScrollbar v-if="!showEmpty" style="max-height: 50vh">
        <VideoCard :is-scroll="false" :videos="videos" />
      </NScrollbar>
      <NEmpty v-else description="没有喜欢的视频哦~">
        <template #icon>
          <NIcon>
            <HeartDislike />
          </NIcon>
        </template>
      </NEmpty>
    </NTabPane>
    <NTabPane name="collect" tab="收藏">
      <NScrollbar  v-if="!showEmpty" style="max-height: 50vh">
        <VideoCard :is-scroll="false" :videos="videos" />
      </NScrollbar>
      <NEmpty v-else description="没有收藏的视频哦~">
        <template #icon>
          <NIcon>
            <RemoveCircle />
          </NIcon>
        </template>
      </NEmpty>
    </NTabPane>
    <NTabPane name="history" tab="观看历史">
      <NScrollbar  v-if="!showEmpty" style="max-height: 50vh">
        <VideoCard :is-scroll="false" :videos="videos" />
      </NScrollbar>
      <NEmpty v-else description="没有浏览过视频哦~">
        <template #icon>
          <NIcon>
            <EyeOff />
          </NIcon>
        </template>
      </NEmpty>
    </NTabPane>
  </NTabs>
</template>
<script lang="ts" setup>
import { NScrollbar, NTabPane, NTabs, NEmpty, NIcon } from "naive-ui";
import VideoCard from "../VideoCard.vue";
import { getHistoryVideosListReq, userVideoListReq } from "@/apis/video";
import { userFavouriteListReq, userStarListReq } from "@/apis/favourite";
import { onMounted, ref, computed } from "vue";
import { userStore } from "@/stores/user";
import { HeartDislike, EyeOff, VideocamOff, RemoveCircle } from "@vicons/ionicons5";

const videos = ref<Array<any>>([]);

interface propsType {
  userId: number;
}

const deletable = computed(() =>
  props.userId == userStore().user_id ? true : false,
);
const props = defineProps<propsType>();
const myid = ref<number>(-1);
const showEmpty = ref<boolean>(false)
onMounted(() => {
  getMyWork();
  
});

const getMyWork = () => {
  let uid = props.userId.toString();
  let uid_num = parseInt(uid);
  myid.value = uid_num;
  userVideoListReq(myid.value).then((res: any) => {
    videos.value = res.video_list;
    if(videos.value.length == 0){
      showEmpty.value = true
    }
  });
};

const getMyFavourite = () => {
  userFavouriteListReq(myid.value).then((res: any) => {
    videos.value = res.video_list;
    if(videos.value.length == 0){
      showEmpty.value = true
    }
  });
};

const getMyCollect = () => {
  userStarListReq(myid.value).then((res: any) => {
    videos.value = res.video_list;
    if(videos.value.length == 0){
      showEmpty.value = true
    }
  });
};

const getMyHistory = () => {
  getHistoryVideosListReq().then((res: any) => {
    videos.value = res.video_list;
    if(videos.value.length == 0){
      showEmpty.value = true
    }
  });
};

const handleUpdate = (paneName: string) => {
  showEmpty.value = false
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
<style></style>
