<!--
 * @Author: huangsihao7
 * @Date: 2023-10-30 11:17:41
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-04 17:46:37
 * @FilePath: \scooter-WSVA\frontend\src\components\video\VideoCardListCom.vue
 * @Description: 卡片形式的视频信息展示组件
-->
<script lang="ts" setup>
import { onMounted } from "vue";
import { NButton, NCard, NEllipsis, NIcon, NTag } from "naive-ui";
import { Heart, Play, Trash } from "@vicons/ionicons5";
import { VideoType } from "@/apis/interface";
import { useRouter } from "vue-router";
import { myDeleteVideo } from "@/apis/video";
import { useMessage } from "naive-ui";

interface propsType {
  videos: Array<VideoType>;
  deletable?: boolean;
}

const props = defineProps<propsType>();
const router = useRouter();
const message = useMessage();

const handleShowVideo = (info: VideoType) => {
  router.push({ name: "video", params: { id: info.video_id } });
};

// 跳转路由，展示视频
const GetVideoLink = () => {};

// 删除我的视频
const deleteVideo = (info: any) => {
  myDeleteVideo(info.video_id).then(() => {
    message.success("删除成功");
    window.location.reload();
  });
};

onMounted(() => {
  //info.content
});
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
        v-lazy="info.cover_url"
        class="image-css"
        alt="img"
        @click="handleShowVideo(info)"
      />
    </template>
    <NButton
      v-if="props.deletable"
      class="btn"
      quaternary
      circle
      type="error"
      @click="deleteVideo(info)"
    >
      <template #icon>
        <NIcon><Trash /></NIcon>
      </template>
    </NButton>
    <div class="video-space" style="position: relative" @click="GetVideoLink">
      <NTag
        class="time"
        round
        :bordered="false"
        :color="{ color: '#e6e6e64a' }"
      >
        {{ info.duration }}
        <template #icon>
          <NIcon color="#fff" :component="Play" />
        </template>
      </NTag>
      <NTag
        class="like"
        round
        :bordered="false"
        :color="{ color: '#f93b3b4a' }"
      >
        {{ info.favorite_count }}
        <template #icon>
          <NIcon color="#fff" :component="Heart" />
        </template>
      </NTag>
    </div>
    <div class="card-footer">
      <div class="content">
        <NEllipsis :tooltip="false" class="title-eli">
          {{ info.title }}
        </NEllipsis>
      </div>
      <div class="other-info">
        <span class="name">
          <NEllipsis :tooltip="false" class="name-eli">
            {{ info.author == undefined ? "null" : info.author.name }}
          </NEllipsis>
        </span>
        <span class="date">
          <NEllipsis :tooltip="false" class="date-eli">
            {{ info.create_time }}
          </NEllipsis>
        </span>
      </div>
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
  .n-card__content {
    height: 200px;
    width: 100%;
  }

  .image-css {
    width: 100%;
    object-fit: fill;
    height: 50vh;
  }

  .n-card .n-card-cover img {
    height: 60vh;
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
    // bottom: 40px;
    bottom: 10px;
    right: 0;
    color: white;
  }

  .card-footer {
    height: 5vh;
    span,
    p,
    .content,
    .other-info {
      text-align: left;
    }

    .other-info {
      position: absolute;
      bottom: 0;
      width: 100%;
      .name {
        display: inline-block;
        font-weight: 400;
        font-size: 0.8rem;
        color: rgb(95, 95, 95);
        width: 37%;
      }
      .date {
        display: inline-block;
        margin-left: 6px;
        text-align: right;
        font-weight: 200;
        font-size: 0.6rem;
        color: rgb(95, 95, 95);
        width: 45%;
      }
    }
    .content {
      padding-top: 6px;
      span {
        font-weight: bold;
        color: #333333;
      }
    }
  }
  .btn {
    position: absolute;
    top: 0;
    right: 0;
  }
}
</style>
