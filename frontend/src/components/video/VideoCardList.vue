<!--
 * @Author: huangsihao7
 * @Date: 2023-10-30 11:17:41
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-30 23:12:00
 * @FilePath: \scooter-WSVA\frontend\src\components\video\VideoCardList.vue
 * @Description: 
-->
<script lang="ts" setup>
import { onMounted } from "vue";
import { NTag, NIcon, NCard, NEllipsis } from "naive-ui";
import { Play, Heart } from "@vicons/ionicons5";
import { VideoType } from "@/apis/interface";
interface propsType {
  videos: Array<VideoType>;
}
const props = defineProps<propsType>();

const handleShowVedio = () => {
  console.log("show");
};

// 跳转路由，展示视频
const GetVideoLink = () => {};

onMounted(() => {
  //info.content
});
</script>
<template>
  <NCard
    v-for="(info, index) in props.videos"
    :key="index"
    class="box-card"
    style="width: calc((100vw - 260px) / 6)"
  >
    <template #cover>
      <img class="image-css" :src="info.cover_url" @click="handleShowVedio" />
    </template>
    <div class="video-space" style="position: relative" @click="GetVideoLink">
      <NTag class="time" round :bordered="false" type="info">
        {{ info.duration }}
        <template #icon>
          <NIcon color="#fff" :component="Play" />
        </template>
      </NTag>
      <NTag class="like" round :bordered="false" type="error">
        {{ info.favorite_count }}
        <template #icon>
          <NIcon color="#fff" :component="Heart" />
        </template>
      </NTag>
    </div>
    <div class="card-footer">
      <div class="content">
        <NEllipsis
          :tooltip="false"
          style="max-width: calc((100vw - 360px) / 4)"
        >
          {{ info.title }}
        </NEllipsis>
      </div>
      <div class="name">
        <span>{{ info.author == undefined ? "null" : info.author.name }}</span>
        <span class="date">{{ info.create_time }}</span>
      </div>
    </div>
  </NCard>
</template>
<style lang="scss" scoped>
.card-space {
  overflow: scroll-y;
  max-height: calc(100vh - 60px);
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
    height: 100%;
    width: 100%;
    object-fit: fill;
    height: 30vh;
  }

  .n-card .n-card-cover img {
    height: 30vh;
  }

  .video-space {
    height: 100%;
    width: 100%;
  }

  .time {
    position: absolute;
    // bottom: 40px;
    bottom: 10px;
    left: 10px;
    color: white;
  }

  .like {
    position: absolute;
    // bottom: 40px;
    bottom: 10px;
    right: 10px;
    color: white;
  }

  .card-footer {
    span,
    p,
    .content,
    .name {
      margin-left: 10px;
      text-align: left;
    }

    .content span {
      font-weight: bold;
      color: #333333;
    }

    span {
      font-weight: 400;
      font-size: 0.8rem;
      color: rgb(95, 95, 95);
    }
  }
}
</style>
