<!--
 * @Author: Xu Ning
 * @Date: 2023-10-26 18:39:00
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-30 13:00:15
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\components\video\VideoPlus.vue
-->

<script lang="ts" setup>
import Dplayer from "@/components/video/VideoCom.vue";
import Hls from "hls.js";
import { ref, reactive, onMounted } from "vue";
import { NIcon, NButton, NAvatar } from "naive-ui";
import { Heart, ArrowRedo, ChatbubbleEllipses, Star, Add, Checkmark } from "@vicons/ionicons5";
import { VideoType } from "@/apis/interface";
import { ElMessageBox, ElMessage } from "element-plus";
import useClipboard from "vue-clipboard3";

interface propsType {
  video: VideoType;
  index: number;
  onplay: number;
}

const props = defineProps<propsType>();

const emit = defineEmits(["comment-visible-update"]);

const thisVideo = ref<VideoType>();

const { toClipboard } = useClipboard();

const dplayerObj = reactive({
  video: {
    url: props.video.playUrl, //视频地址
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
    // id: '9E2E3368B56CDBB4',
    // api: 'https://api.prprpr.me/dplayer/',
    // token: 'tokendemo',
    // maximum: 1000,
    // addition: ['https://api.prprpr.me/dplayer/v3/bilibili?aid=4157142'],
    // user: 'DIYgod',
    // bottom: '15%',
    // unlimited: true,
    // speedRate: 0.5,
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

const copy = async (msg: any) => {
  try {
    // 复制
    await toClipboard(msg);
    // 复制成功
  } catch (e) {
    // 复制失败
    ElMessage({
      type: "error",
      message: "复制失败",
    });
  }
};

// 喜欢按钮的操作
const handleLikeBtn = () => {
  if (thisVideo.value) {
    if (!thisVideo.value?.isFavorite) {
      thisVideo.value.favoriteCount++;
    } else {
      thisVideo.value.favoriteCount--;
    }
    if (!thisVideo.value?.isFavorite) {
      likeAnimateClass.value = "animate__heartBeat";
    } else {
      likeAnimateClass.value = "";
    }
    thisVideo.value.isFavorite = !thisVideo.value.isFavorite;
  }
  // TODO: 发请求
};

// 收藏按钮的操作
const handleCollectBtn = () => {
  if (thisVideo.value) {
    if (!thisVideo.value?.starCount) {
      thisVideo.value.starCount++;
    } else {
      thisVideo.value.starCount--;
    }

    // TODO: 等待后端返回collect值
    if (!thisVideo.value?.isStar) {
      collectAnimateClass.value = "animate__heartBeat";
    } else {
      collectAnimateClass.value = "";
    }
    thisVideo.value.isStar = !thisVideo.value.isStar;
  }

  // TODO: 发请求
};

// 评论按钮的操作
const handleCommentBtn = () => {
  commentVisible.value = !commentVisible.value;
  emit("comment-visible-update");
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
        ElMessage({
          type: "info",
          message: `复制成功`,
        });
        copyFlag.value = false;
      }
    },
  });
};

// 关注的操作 
const updateFollow = (flag:boolean) =>{
  if(flag && thisVideo.value){
    thisVideo.value.author.is_follow = flag
  }
  else if(!flag && thisVideo.value){
    thisVideo.value.author.is_follow = flag
  }
  else{
    ElMessage({type:'error', message:'关注失败'})
  }
}

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
          <p class="time">{{ props.video.createTime }}</p>
        </div>
        <div class="content">{{ props.video.title }}</div>
      </div>
      <div class="video-interaction-box">
        <div class="avatar">
          <div >
            <n-avatar
              round
              size="medium"
              :src="thisVideo?.author.avatar"
            />
          </div>
          <n-button color="#ffa51d8f" @click="updateFollow(true)" v-if="!thisVideo?.author.is_follow" class="avatar-btn animate__bounceIn" size="tiny" circle type="warning">
            <template #icon>
              <n-icon><Add /></n-icon>
            </template>
          </n-button>
          <n-button color="#ffa51d8f" @click="updateFollow(false)" v-else class="avatar-btn animate__bounceIn" size="tiny" circle type="warning">
            <template #icon>
              <n-icon><Checkmark /></n-icon>
            </template>
          </n-button>
        </div>
        <div class="like">
          <div :class="likeAnimateClass">
            <NButton
              v-if="thisVideo?.isFavorite"
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
          <p>{{ thisVideo?.favoriteCount }}</p>
        </div>
        <div class="comment">
          <div :class="commentAnimateClass">
            <NButton class="btn" text color="#fff" @click="handleCommentBtn">
              <NIcon>
                <ChatbubbleEllipses />
              </NIcon>
            </NButton>
          </div>
          <p>{{ props.video.commentCount }}</p>
        </div>
        <div class="collect">
          <div :class="collectAnimateClass">
            <NButton
              v-if="thisVideo?.isStar"
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
          <p>{{ props.video.starCount }}</p>
        </div>
        <div class="share">
          <div :class="shareAnimateClass">
            <NButton class="btn" text color="#fff" @click="handleShareBtn">
              <NIcon>
                <ArrowRedo />
              </NIcon>
            </NButton>
          </div>
          <p>{{ thisVideo?.shareCount }}</p>
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
    
    .avatar{
      margin-bottom: -18px;
      .avatar-btn{
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
        color:#fff;
      }
    }
  }
}
</style>
