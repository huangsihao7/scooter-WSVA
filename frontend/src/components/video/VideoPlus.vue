<!--
 * @Author: Xu Ning
 * @Date: 2023-10-26 18:39:00
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-01 01:00:10
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\components\video\VideoPlus.vue
-->

<script lang="ts" setup>
import Dplayer from "@/components/video/VideoCom.vue";
import Hls from "hls.js";
import { computed, onMounted, reactive, ref } from "vue";
import { NAvatar, NButton, NIcon, useMessage } from "naive-ui";
import {
  Add,
  ArrowRedo,
  ChatbubbleEllipses,
  Checkmark,
  Heart,
  Star,
} from "@vicons/ionicons5";
import { VideoType } from "@/apis/interface";
import { ElMessageBox } from "element-plus";
import useClipboard from "vue-clipboard3";
import { doFavourite, doStar } from "@/apis/favourite";
import { doFollow } from "@/apis/relation";
import { userStore } from "@/stores/user";
import { baseURL } from "@/axios";

interface propsType {
  video: VideoType;
  index: number;
  onplay: number;
}

const message = useMessage();
const props = defineProps<propsType>();

const emit = defineEmits(["comment-visible-update"]);

const thisVideo = ref<VideoType>();
const userId = computed(() => userStore().user_id);
const { toClipboard } = useClipboard();

const dplayerObj = reactive({
  videoId: props.video.video_id,
  video: {
    url: props.video.play_url, //视频地址
    type: "mp4",
    customType: {
      customHls: function (video: any, player: any) {
        console.log(player);
        const hls = new Hls(); //实例化Hls  用于解析m3u8
        hls.loadSource(video.src);
        hls.attachMedia(video);
      },
    },
  },
  danmaku: {
    id: props.video.video_id,
    // api: "https://angustar.tech/api/dplayer/",
    api: baseURL + '/danmu/list',
    token: userStore().token,
    // maximum: 1000,
    // addition: ['https://api.prprpr.me/dplayer/v3/bilibili?aid=4157142'],
    // addition: [baseURL + '/danmu/listv3?/video_id=' + props.video.video_id],
    user: userStore().user_id,
    bottom: "15%",
    unlimited: true,
    speedRate: 0.5,
  },
  contextmenu: [
    {
      text: "custom1",
      link: "https://github.com/DIYgod/DPlayer",
    },
    {
      text: "custom2",
      click: (player: any) => {
        console.log(player);
      },
    },
  ],
  highlight: [
    {
      time: 20,
      text: "这是第 20 秒",
    },
    {
      time: 120,
      text: "这是 2 分钟",
    },
  ],
});

// 复制分享链接
const copy = async (msg: any) => {
  try {
    // 复制
    await toClipboard(msg);
    // 复制成功
  } catch (e) {
    // 复制失败
    message.error("复制失败");
  }
};

// 喜欢按钮的操作
const handleLikeBtn = () => {
  let action_type = -1;
  if (thisVideo.value) {
    if (!thisVideo.value?.is_favorite) {
      thisVideo.value.favorite_count++;
      action_type = 1;
    } else {
      thisVideo.value.favorite_count--;
      action_type = 2;
    }
    if (!thisVideo.value?.is_favorite) {
      likeAnimateClass.value = "animate__heartBeat";
    } else {
      likeAnimateClass.value = "";
    }
    thisVideo.value.is_favorite = !thisVideo.value.is_favorite;
  }
  // TODO: 发请求
  doFavourite(props.video.video_id, action_type).then((res: any) => {
    console.log(res);
  });
};

// 收藏按钮的操作
const handleCollectBtn = () => {
  let action_type = -1;
  if (thisVideo.value) {
    if (!thisVideo.value?.star_count) {
      thisVideo.value.star_count++;
      action_type = 1;
    } else {
      thisVideo.value.star_count--;
      action_type = 2;
    }
    if (!thisVideo.value?.is_star) {
      collectAnimateClass.value = "animate__heartBeat";
    } else {
      collectAnimateClass.value = "";
    }
    thisVideo.value.is_star = !thisVideo.value.is_star;
  }
  doStar(props.video.video_id, action_type).then((res: any) => {
    console.log(res);
  });
};

// 评论按钮的操作
const handleCommentBtn = () => {
  commentVisible.value = !commentVisible.value;
  emit("comment-visible-update", thisVideo);
};

const copyFlag = ref<boolean>(false);

// 分享按钮的操作
const handleShareBtn = () => {
  shareVisible.value = !shareVisible.value;
  const currentUrl: string = window.location.href;
  copy(currentUrl);
  ElMessageBox.alert(currentUrl, "分享", {
    confirmButtonText: "复制",
    center: true,
    beforeClose: (action, instance, done) => {
      if (action === "confirm") {
        instance.confirmButtonText = "复制中...";
        instance.confirmButtonLoading = true;
        console.log("action", action);
        copy(currentUrl);
        copyFlag.value = true;
        setTimeout(() => {
          instance.confirmButtonLoading = false;
          done();
        }, 300);
      } else {
        copyFlag.value = false;
        done();
      }
    },
    callback: () => {
      console.log("cpf", copyFlag.value);
      if (copyFlag.value) {
        message.success("复制成功");
        copyFlag.value = false;
      }
    },
  });
};

// 关注的操作
const updateFollow = (flag: boolean) => {
  let action = flag ? 1 : 2;
  if (flag && thisVideo.value) {
    thisVideo.value.author.is_follow = flag;
  } else if (!flag && thisVideo.value) {
    thisVideo.value.author.is_follow = flag;
  } else {
    message.error("关注失败");
  }
  doFollow(props.video.author.id, action).then((res: any) => {
    console.log(res);
  });
};

const likeAnimateClass = ref<String>("");
const collectAnimateClass = ref<String>("");
const commentAnimateClass = ref<String>("");
const shareAnimateClass = ref<String>("");
const commentVisible = ref<boolean>(false);
const shareVisible = ref<boolean>(false);
onMounted(() => {
  thisVideo.value = props.video;
});
</script>

<template>
  <div>
    <div class="video-container">
      <Dplayer
        :video-id="dplayerObj.videoId"
        :video="dplayerObj.video"
        :danmaku="dplayerObj.danmaku"
        :contextmenu="dplayerObj.contextmenu"
        :highlight="dplayerObj.highlight"
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
            <NAvatar round size="medium" :src="thisVideo?.author.avatar" />
          </div>
          <NButton
            v-if="
              !thisVideo?.author.is_follow && thisVideo?.author.id != userId
            "
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
              thisVideo?.author.is_follow && thisVideo?.author.id != userId
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
              v-if="thisVideo?.is_favorite"
              class="btn"
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
              color="#ffffff"
              @click="handleLikeBtn"
            >
              <NIcon>
                <Heart />
              </NIcon>
            </NButton>
          </div>
          <p>{{ thisVideo?.favorite_count }}</p>
        </div>
        <div class="comment">
          <div :class="commentAnimateClass">
            <NButton class="btn" text color="#fff" @click="handleCommentBtn">
              <NIcon>
                <ChatbubbleEllipses />
              </NIcon>
            </NButton>
          </div>
          <p>{{ props.video.comment_count }}</p>
        </div>
        <div class="collect">
          <div :class="collectAnimateClass">
            <NButton
              v-if="thisVideo?.is_star"
              class="btn"
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
              text
              color="#fff"
              @click="handleCollectBtn"
            >
              <NIcon>
                <Star />
              </NIcon>
            </NButton>
          </div>
          <p>{{ props.video.star_count }}</p>
        </div>
        <div class="share">
          <div :class="shareAnimateClass">
            <NButton class="btn" text color="#fff" @click="handleShareBtn">
              <NIcon>
                <ArrowRedo />
              </NIcon>
            </NButton>
          </div>
          <p>{{ thisVideo?.share_count }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
.video-container {
  .el-button {
    z-index: 9999;
  }

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
    bottom: calc((100vh - 420px) / 2);
    right: 0;
    height: 360px;
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
