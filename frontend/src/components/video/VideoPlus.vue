<!--
 * @Author: Xu Ning
 * @Date: 2023-10-26 18:39:00
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-03 19:20:52
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\components\video\VideoPlus.vue
-->

<script lang="ts" setup>
import Dplayer from "@/components/video/VideoCom.vue";
import Hls from "hls.js";
import { computed, reactive, ref } from "vue";
import { NAvatar, NButton, NIcon, useMessage, NModal, NSpin } from "naive-ui";
import {
  Add,
  ArrowRedo,
  ChatbubbleEllipses,
  Checkmark,
  Heart,
  Star,
  Download,
} from "@vicons/ionicons5";
import { VideoType } from "@/apis/interface";
import { doFavourite, doStar } from "@/apis/favourite";
import { doFollow } from "@/apis/relation";
import { userStore } from "@/stores/user";
import { baseURL } from "@/axios";
import { DownloadType } from "@/apis/interface";
import { useClipboard } from "@vueuse/core";
const { copy, copied } = useClipboard();
interface propsType {
  video: VideoType;
  index: number;
  onplay: number;
}

const message = useMessage();
const props = defineProps<propsType>();
const emit = defineEmits(["comment-visible-update"]);
const likeAnimateClass = ref<String>("");
const collectAnimateClass = ref<String>("");
const commentAnimateClass = ref<String>("");
const shareAnimateClass = ref<String>("");
const commentVisible = ref<boolean>(false);
const shareVisible = ref<boolean>(false);
const logedFlag = computed(() => !userStore().isLoggedIn);
const thisVideo = reactive<VideoType>({
  video_id: props.video.video_id,
  author: props.video.author,
  play_url: props.video.play_url,
  cover_url: props.video.cover_url,
  favorite_count: props.video.favorite_count,
  comment_count: props.video.comment_count,
  is_favorite: props.video.is_favorite,
  title: props.video.title,
  create_time: props.video.create_time,
  is_star: props.video.is_star,
  star_count: props.video.star_count,
  share_count: props.video.share_count,
  duration: props.video.duration,
});
const userId = computed(() => userStore().user_id);
const dplayerObj = reactive({
  videoId: props.video.video_id,
  video: {
    url: props.video.play_url, //视频地址
    type: "mp4",
    customType: {
      customHls: function (video: any) {
        const hls = new Hls(); //实例化Hls  用于解析m3u8
        hls.loadSource(video.src);
        hls.attachMedia(video);
      },
    },
  },
  danmaku: {
    id: props.video.video_id,
    api: baseURL + "/danmu/list",
    token: userStore().token,
    user: userStore().user_id,
    bottom: "15%",
    unlimited: true,
    speedRate: 0.5,
  },
  contextmenu: [
    {
      text: "项目GitHub地址",
      link: "https://github.com/huangsihao7/scooter-WSVA",
    },
  ],
});
// 分享链接
const url = computed(() => {
  let currentUrl: string = window.location.href;
  let firstSegment: string = currentUrl.substring(
    0,
    currentUrl.indexOf("/", 8),
  );
  return firstSegment + "/video/" + props.video.video_id;
});

// 喜欢按钮的操作
const handleLikeBtn = () => {
  let action_type = -1;
  if (thisVideo) {
    if (!thisVideo.is_favorite) {
      thisVideo.favorite_count++;
      action_type = 1;
    } else {
      thisVideo.favorite_count--;
      action_type = 2;
    }
    if (!thisVideo?.is_favorite) {
      likeAnimateClass.value = "animate__heartBeat";
    } else {
      likeAnimateClass.value = "";
    }
    thisVideo.is_favorite = !thisVideo.is_favorite;
  }
  doFavourite(props.video.video_id, action_type).then(() => {});
};

// 收藏按钮的操作
const handleCollectBtn = () => {
  let action_type = -1;
  if (thisVideo) {
    if (!thisVideo.is_star) {
      console.log('starcount',thisVideo.star_count)
      thisVideo.star_count++;
      console.log('starcount2',thisVideo.star_count)

      action_type = 1;
    } else {
      thisVideo.star_count--;
      action_type = 2;
    }
    if (!thisVideo?.is_star) {
      collectAnimateClass.value = "animate__heartBeat";
    } else {
      collectAnimateClass.value = "";
    }
    thisVideo.is_star = !thisVideo.is_star;
  }
  doStar(props.video.video_id, action_type).then(() => {});
};

// 评论按钮的操作
const handleCommentBtn = () => {
  commentVisible.value = !commentVisible.value;
  emit("comment-visible-update", thisVideo.video_id);
};

// 分享按钮的操作
const handleShareBtn = () => {
  shareVisible.value = true;
};

// 视频下载函数
const handleDownloadBtn = () => {
  if (thisVideo) {
    let videoUrl = thisVideo.play_url;
    const lastDotIndex = videoUrl.lastIndexOf(".");
    let url = "";
    if (lastDotIndex !== -1) {
      const beforeDot = videoUrl.substring(0, lastDotIndex);
      const afterDot = videoUrl.substring(lastDotIndex + 1);
      url = beforeDot + "-1." + afterDot;
    }
    if (url != "") {
      let title = thisVideo.title.toString();
      const downItem: DownloadType = { url: url, title: title };
      message.info("正在下载");
      downLoad(downItem);
    } else {
      message.error("下载路径错误");
    }
  } else {
    message.error("下载失败");
  }
};

// 仅支持视频下载和图片下载
function downLoad(item: DownloadType) {
  let url = item.url;
  // let fileName = url.slice(url.lastIndexOf("/") + 1); //下载的文件名换成自己的
  let fileName = item.title; //dayjs
  let xhr = new XMLHttpRequest();
  xhr.open("GET", url, true);
  xhr.responseType = "blob"; // 返回类型blob
  xhr.onload = () => {
    if (xhr.readyState === 4 && xhr.status === 200) {
      let blob = xhr.response;
      let downLoadUrl = window.URL.createObjectURL(
        new Blob([blob], {
          // type: item.type === "video" ? "video/mp4" : "image/jpeg",
          type: "video/mp4",
        }),
      );
      let a = document.createElement("a");
      a.download = fileName;
      a.href = downLoadUrl;
      a.style.display = "none";
      document.body.appendChild(a);
      a.click();
      a.remove();
    }
  };
  xhr.send();
}

// 关注的操作
const updateFollow = (flag: boolean) => {
  let action = flag ? 1 : 2;
  if (flag && thisVideo) {
    thisVideo.author.is_follow = flag;
  } else if (!flag && thisVideo) {
    thisVideo.author.is_follow = flag;
  } else {
    message.error("关注失败");
  }
  doFollow(props.video.author.id, action).then(() => {});
};

// 链接复制成功回调
const handleCopy = () => {
  copy(url.value);
  message.success("复制成功");
  shareVisible.value = false;
};
</script>

<template>
  <div>
    <div class="video-container">
      <Dplayer
        :video-id="dplayerObj.videoId"
        :video="dplayerObj.video"
        :danmaku="dplayerObj.danmaku"
        :contextmenu="dplayerObj.contextmenu"
        :video-index="props.index"
        :on-play-index="props.onplay"
      />
      <div class="video-info-box">
        <div class="header">
          <p class="title">@{{ props.video.author.name }}</p>
          <p class="time">{{ props.video.create_time }}</p>
        </div>
        <div class="content">{{ props.video.title }}</div>
      </div>
      <div class="video-interaction-box">
        <div class="avatar">
          <div>
            <NAvatar round size="medium" :src="thisVideo.author.avatar" />
          </div>
          <NButton
            v-if="!thisVideo.author.is_follow && thisVideo.author.id != userId"
            color="#ffa51d8f"
            class="avatar-btn animate__bounceIn"
            size="tiny"
            circle
            type="warning"
            @click="updateFollow(true)"
          >
            <template #icon>
              <NIcon>
                <Add />
              </NIcon>
            </template>
          </NButton>
          <NButton
            v-else-if="
              thisVideo.author.is_follow && thisVideo.author.id != userId
            "
            color="#ffa51d8f"
            class="avatar-btn animate__bounceIn"
            size="tiny"
            circle
            type="warning"
            @click="updateFollow(false)"
          >
            <template #icon>
              <NIcon>
                <Checkmark />
              </NIcon>
            </template>
          </NButton>
          <div v-else style="height: 20px"></div>
        </div>
        <div class="like">
          <div :class="likeAnimateClass">
            <NButton
              v-if="thisVideo.is_favorite"
              class="btn"
              :disabled="logedFlag"
              text
              color="rgb(254, 44, 85)"
              @click="handleLikeBtn"
            >
              <NIcon>
                <Heart />
              </NIcon>
            </NButton>
            <NButton
              v-else
              class="btn"
              text
              :disabled="logedFlag"
              color="#ffffff"
              @click="handleLikeBtn"
            >
              <NIcon>
                <Heart />
              </NIcon>
            </NButton>
          </div>
          <p>{{ thisVideo.favorite_count }}</p>
        </div>
        <div class="comment">
          <div :class="commentAnimateClass">
            <NButton
              class="btn"
              :disabled="logedFlag"
              text
              color="#fff"
              @click="handleCommentBtn"
            >
              <NIcon>
                <ChatbubbleEllipses />
              </NIcon>
            </NButton>
          </div>
          <p>{{ thisVideo.comment_count }}</p>
        </div>
        <div class="collect">
          <div :class="collectAnimateClass">
            <NButton
              v-if="thisVideo.is_star"
              class="btn"
              :disabled="logedFlag"
              text
              color="#ffb802"
              @click="handleCollectBtn"
            >
              <NIcon>
                <Star />
              </NIcon>
            </NButton>
            <NButton
              v-else
              class="btn"
              :disabled="logedFlag"
              text
              color="#fff"
              @click="handleCollectBtn"
            >
              <NIcon>
                <Star />
              </NIcon>
            </NButton>
          </div>
          <p>{{ thisVideo.star_count }}</p>
        </div>
        <div class="share">
          <div :class="shareAnimateClass">
            <NButton class="btn" text color="#fff" @click="handleShareBtn">
              <NIcon>
                <ArrowRedo />
              </NIcon>
            </NButton>
          </div>
        </div>
        <div class="share">
          <div :class="shareAnimateClass">
            <NButton class="btn" text color="#fff" @click="handleDownloadBtn">
              <NIcon>
                <Download />
              </NIcon>
            </NButton>
          </div>
        </div>
        <NModal v-model:show="shareVisible" preset="dialog" title="分享">
          分享链接:{{ url }}
          <template #action>
            <NSpin v-if="copied">
              <NButton round> 复制中 </NButton>
            </NSpin>
            <NButton v-else round @click="handleCopy"> 复制 </NButton>
          </template>
        </NModal>
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
.video-container {
  .dplayer {
    width: calc(100vw - 160px);
    height: calc(100vh - 60px);
  }

  .video-info-box {
    display: block;
    height: auto;
    pointer-events: none;
    position: absolute;
    z-index: 2;
    width: calc(100vw - 200px);
    bottom: 48px;
    height: 100px;
    padding: 0 20px;

    .title,
    .time,
    .content {
      color: white;
    }

    .header {
      display: flex;

      .time {
        margin: 25px 0 0 15px;
        font-size: small;
      }

      .title {
        font-size: larger;
        font-weight: bold;
      }
    }

    .content {
      text-align: left;
    }
  }

  .video-interaction-box {
    background: rgba(255, 255, 255, 0.12);
    border-radius: 25px;
    display: block;
    height: auto;
    position: absolute;
    z-index: 999;
    width: 4vw;
    bottom: calc((100vh - 360px) / 2);
    right: 0;
    height: 400px;
    padding: 0 20px;

    .avatar {
      margin-bottom: -18px;

      .avatar-btn {
        position: relative;
        bottom: 18px;
      }
    }

    .avatar,
    .like,
    .collect,
    .comment,
    .share {
      margin-top: 10px;

      .btn {
        font-size: 35px;
      }

      p {
        display: block;
        margin: 0;
        color: #fff;
      }
    }
  }
}
</style>
