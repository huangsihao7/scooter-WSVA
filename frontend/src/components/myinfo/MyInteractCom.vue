<!--
 * @Author: Xu Ning
 * @Date: 2023-10-28 12:30:41
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-06 20:39:39
 * @Description: 我的页面的视频信息展示组件
 * @FilePath: \scooter-WSVA\frontend\src\components\myinfo\MyInteractCom.vue
-->
<template>
  <NTabs
    default-value="work"
    justify-content="space-evenly"
    @update-value="handleUpdate"
  >
    <NTabPane name="work">
      <template #tab>
        <span class="iconfont icon-shipin2 tab-icon"></span>
        作品
      </template>
      <VideoCard
        v-if="!showEmpty"
        :is-scroll="false"
        :videos="videos"
        :deletable="deletable"
      />
      <NEmpty v-else description="没有发布过视频哦~">
        <template #icon>
          <NIcon>
            <VideocamOff />
          </NIcon>
        </template>
      </NEmpty>
    </NTabPane>
    <NTabPane name="favourite">
      <template #tab>
        <span class="iconfont icon-dianzan1 tab-icon"></span>
        喜欢
      </template>
      <VideoCard v-if="!showEmpty" :is-scroll="false" :videos="videos" />
      <NEmpty v-else description="没有喜欢的视频哦~">
        <template #icon>
          <NIcon>
            <HeartDislike />
          </NIcon>
        </template>
      </NEmpty>
    </NTabPane>
    <NTabPane name="collect">
      <template #tab>
        <span class="iconfont icon-shoucang tab-icon"></span>
        收藏
      </template>
      <VideoCard v-if="!showEmpty" :is-scroll="false" :videos="videos" />
      <NEmpty v-else description="没有收藏的视频哦~">
        <template #icon>
          <NIcon>
            <RemoveCircle />
          </NIcon>
        </template>
      </NEmpty>
    </NTabPane>
    <NTabPane name="history" v-if="userId == userStore().user_id">
      <template #tab>
        <span class="iconfont icon-lishijilu tab-icon"></span>
        浏览记录
      </template>
      <VideoCard v-if="!showEmpty" :is-scroll="false" :videos="videos" />
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
import { NTabPane, NTabs, NEmpty, NIcon } from "naive-ui";
import VideoCard from "../cards/VideoCard.vue";
import { getHistoryVideosListReq, userVideoListReq } from "@/apis/video";
import { userFavouriteListReq, userStarListReq } from "@/apis/favourite";
import { onMounted, ref, computed } from "vue";
import { userStore } from "@/stores/user";
import {
  HeartDislike,
  EyeOff,
  VideocamOff,
  RemoveCircle,
} from "@vicons/ionicons5";

const videos = ref<Array<any>>([]);

interface propsType {
  userId: number;
}

const deletable = computed(() =>
  props.userId == userStore().user_id ? true : false,
);
const props = defineProps<propsType>();
const myid = ref<number>(-1);
const showEmpty = ref<boolean>(false);
onMounted(() => {
  getMyWork();
});

const getMyWork = () => {
  let uid = props.userId.toString();
  let uid_num = parseInt(uid);
  myid.value = uid_num;
  userVideoListReq(myid.value).then((res: any) => {
    videos.value = res.video_list;
    if (videos.value.length == 0) {
      showEmpty.value = true;
    }
  });
};

const getMyFavourite = () => {
  userFavouriteListReq(myid.value).then((res: any) => {
    videos.value = res.video_list;
    if (videos.value.length == 0) {
      showEmpty.value = true;
    }
  });
};

const getMyCollect = () => {
  userStarListReq(myid.value).then((res: any) => {
    videos.value = res.video_list;
    if (videos.value.length == 0) {
      showEmpty.value = true;
    }
  });
};

const getMyHistory = () => {
  getHistoryVideosListReq().then((res: any) => {
    videos.value = res.video_list;
    if (videos.value.length == 0) {
      showEmpty.value = true;
    }
  });
};

const handleUpdate = (paneName: string) => {
  showEmpty.value = false;
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
.tab-icon {
  margin-right: 6px;
  font-size: 20px;
}
</style>
