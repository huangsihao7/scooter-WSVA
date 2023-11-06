<!--
 * @Author: Xu Ning
 * @Date: 2023-10-31 20:58:41
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-06 17:36:32
 * @Description: 相关视频推荐卡片
 * @FilePath: \scooter-WSVA\frontend\src\components\cards\VideoRecommendCard.vue
-->
<script setup lang="ts">
import { NCard, NImage } from "naive-ui";
import { NButton, NIcon, NEllipsis } from "naive-ui";
import { Heart } from "@vicons/ionicons5";
import { VideoType } from "@/apis/interface";
import { routeStore } from "@/stores/route";
import { useRouter } from "vue-router";
interface propsType {
  recommendlists: Array<VideoType>;
}
const props = defineProps<propsType>();
const router = useRouter();
// 查看视频
const goVideo = (video: VideoType) => {
  routeStore().name = "video";
  router.push({ name: "video", params: { id: video.video_id } });
};
</script>

<template>
  <NCard
    v-for="(video, index) in props.recommendlists"
    :key="index"
    class="video-card"
    hoverable
  >
    <NImage
      :src="video.cover_url"
      width="100"
      preview-disabled
      @click="goVideo(video)"
    />
    <div class="video-info">
      <span class="title">
        <NEllipsis expand-trigger="click" line-clamp="2" :tooltip="false">
          {{ video.title }}
        </NEllipsis>
      </span>
      <div class="footer">
        <div class="like">
          <NButton class="btn" text color="#3f3f3f">
            <NIcon>
              <Heart />
            </NIcon>
          </NButton>
          <p>{{ video.favorite_count }}</p>
        </div>
        <span>
          <NEllipsis style="max-width: 140px" :tooltip="true">
            @ {{ video.author.name }}
          </NEllipsis>
        </span>
      </div>
    </div>
  </NCard>
</template>

<style lang="scss">
.video-card {
  height: 100%;
  margin-top: 10px;
  .n-card__content {
    display: flex;
    flex-direction: row;
    justify-content: start;
  }

  .n-image {
    width: 100px;
    height: 100px;
  }
  .video-info {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    margin-left: 14px;
    .title {
      display: block;
      font-size: 1.2rem;
      font-weight: bold;
      text-align: left;
    }
    .footer {
      color: #333;
      display: flex;
      justify-content: space-between;
      .like {
        width: 20%;
        .btn,
        p {
          display: inline;
        }
        p {
          margin-left: 5px;
        }
      }
    }
  }
}
</style>
