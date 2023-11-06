<!--
 * @Author: huangsihao7
 * @Date: 2023-10-30 11:17:41
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-06 16:36:12
 * @FilePath: \scooter-WSVA\frontend\src\components\live\LiveCardListCom.vue
 * @Description: 卡片形式的视频信息展示组件
-->
<script lang="ts" setup>
import { NCard, NEllipsis, NAvatar, NGrid, NGridItem } from "naive-ui";
import { LiveType } from "@/apis/interface";
import { liveStore } from "@/stores/live";
import { useRouter } from "vue-router";

interface propsType {
  videos: any;
  // videos: Array<LiveType>;
}

const props = defineProps<propsType>();
const router = useRouter();

// 跳转直播
const goLive = (info: LiveType) => {
  liveStore().live_url = info.live_url;
  liveStore().live_id = info.uid;
  router.push({
    name: "live",
    params: { id: info.uid },
  });
};

// 跳转用户信息
const goUserInfo = (info: LiveType) => {
  router.push({ name: "userinfo", params: { id: info.uid } });
};
</script>
<template>
  <NCard
    v-for="(info, index) in props.videos"
    :key="index"
    hoverable
    class="box-card"
  >
    <template #cover>
      <img
        v-lazy="info.live_cover_url"
        class="image-css"
        alt="img"
        @click="goLive(info)"
      />
    </template>
    <div class="card-footer">
      <NGrid item-responsive :x-gap="12" responsive="screen" cols="8">
        <NGridItem span="2">
          <div class="out-avatar">
            <div id="effect-music">
              <NAvatar
                class="avatar image"
                :src="info.avatar"
                round
                @click="goUserInfo(info)"
              />
              <div class="wave"></div>
              <div class="wave"></div>
              <div class="wave"></div>
              <div class="wave"></div>
            </div>
          </div>
        </NGridItem>
        <NGridItem span="6">
          <div class="content">
            <NEllipsis :tooltip="false">
              {{ info.signature }}
            </NEllipsis>
          </div>
          <div class="name">
            <NEllipsis :tooltip="false">
              {{ info.name }}
            </NEllipsis>
          </div>
        </NGridItem>
      </NGrid>
    </div>
  </NCard>
</template>
<style lang="scss" scoped>
.card-space {
  max-height: calc(100vh - 60px);
}
@media (max-width: 1200px) {
  /* 在这里设置适应小屏幕的样式规则 */
  .box-card {
    width: calc((100vw - 260px) / 2);
  }
  .title-eli {
    max-width: calc((100vw - 360px) / 4);
  }
}

@media (min-width: 1200px) and (max-width: 1800px) {
  /* 在这里设置适应中等屏幕的样式规则 */
  .box-card {
    width: calc((100vw - 260px) / 4);
  }
}

/* 在页面宽度大于1200px时应用以下样式 */
@media (min-width: 1800px) {
  .box-card {
    width: calc((100vw - 260px) / 6);
  }
}

.box-card {
  margin-top: 15px;
  margin-left: 10px;
  display: inline-block;
  .image-css {
    width: 100%;
    object-fit: fill;
    height: 40vh;
  }

  .n-card .n-card-cover img {
    height: 50vh;
  }

  .video-space {
    height: 100%;
    width: 100%;
  }

  .time {
    position: absolute;
    // bottom: 40px;
    bottom: 10px;
    left: 0;
    color: white;
  }

  .like {
    position: absolute;
    bottom: 10px;
    right: 0;
    color: white;
  }

  .card-footer {
    padding-top: 12px;
    .avatar {
      height: 100%;
      width: auto;
    }

    .out-avatar {
      height: 100%;
      width: 100%;
      border-radius: 10px;
    }

    .content,
    .name {
      text-align: left;
    }
    .content {
      font-weight: bold;
      color: #333333;
    }
    .name {
      font-weight: 400;
      font-size: 0.8rem;
      color: rgb(95, 95, 95);
    }
  }
  .btn {
    position: absolute;
    top: 0;
    right: 0;
  }
}

svg {
  width: 3.75em;
  animation: 1.5s spin ease infinite;
}

.ring {
  fill: none;
  stroke: hsla(341, 97%, 59%, 0.3);
  stroke-width: 2;
}

.ball {
  fill: #fc2f70;
  stroke: none;
}

#effect-music {
  position: relative;
  margin: auto;
  width: 100%;
  height: 60px;
  overflow: hidden;
}

#effect-music > .image {
  position: absolute;
  left: 0;
  right: 0;
  top: 0;
  bottom: 0;
  margin: auto;
  width: 40px;
  height: 40px;
  // background: url(https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1575820971101&di=935fabf797c3b30f45f7ed90666f9399&imgtype=0&src=http%3A%2F%2Fimglf1.ph.126.net%2FJceVstVIHMeAlbMvdbYwlA%3D%3D%2F2456713596830921578.jpg);
  background-size: cover;
  background-position: center center;
  border-radius: 50%;
  -webkit-border-radius: 50%;
  -moz-border-radius: 50%;
  -ms-border-radius: 50%;
  -o-border-radius: 50%;
  animation: rotate 10s linear 0s infinite;
  -webkit-animation: rotate 10s linear 0s infinite;
}

#effect-music > .wave {
  position: absolute;
  opacity: 0;
  left: 0;
  right: 0;
  top: 0;
  bottom: 0;
  margin: auto;
  width: 26px;
  height: 26px;
  border-radius: 50%;
  border: 2px solid #95c9ff;
  animation: wave 4s linear 0s infinite;
  -webkit-animation: wave 4s linear 0s infinite;
}

#effect-music > .wave::after {
  content: "";
  position: absolute;
  top: -4px;
  left: 50%;
  width: 6px;
  height: 6px;
  overflow: hidden;
  border-radius: 5px;
  -webkit-border-radius: 5px;
  -moz-border-radius: 5px;
  -ms-border-radius: 5px;
  -o-border-radius: 5px;
  background-color: #4678db;
}

#effect-music > .wave:nth-child(2) {
  animation-delay: 1s;
}

#effect-music > .wave:nth-child(3) {
  animation-delay: 2s;
}

#effect-music > .wave:nth-child(4) {
  animation-delay: 3s;
}

@keyframes rotate {
  from {
    transform: rotate(0deg);
    -webkit-transform: rotate(0deg);
    -moz-transform: rotate(0deg);
    -ms-transform: rotate(0deg);
    -o-transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
    -webkit-transform: rotate(360deg);
    -moz-transform: rotate(360deg);
    -ms-transform: rotate(360deg);
    -o-transform: rotate(360deg);
  }
}

@keyframes wave {
  from {
    opacity: 1;
    transform: rotate(0deg) scale(1);
    -webkit-transform: rotate(0deg) scale(1);
    -moz-transform: rotate(0deg) scale(1);
    -ms-transform: rotate(0deg) scale(1);
    -o-transform: rotate(0deg) scale(1);
  }
  to {
    opacity: 0;
    transform: rotate(-300deg) scale(2.2);
    -webkit-transform: rotate(-300deg) scale(2.2);
    -moz-transform: rotate(-300deg) scale(2.2);
    -ms-transform: rotate(-300deg) scale(2.2);
    -o-transform: rotate(-300deg) scale(2.2);
  }
}
</style>
