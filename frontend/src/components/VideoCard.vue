<!--
 * @Author: Xu Ning
 * @Date: 2023-10-27 14:13:32
 * @LastEditors: huangsihao7 1057434651@qq.com
 * @LastEditTime: 2023-10-30 13:55:05
 * @Description: 各个视频分类的视频卡片渲染
 * @FilePath: /scooter-WSVA/frontend/src/components/VideoCard.vue
-->
<script lang="ts" setup>
import { onMounted } from "vue";
import { NScrollbar, NSpace } from "naive-ui";
import VideoCardList from "@/components/video/VideoCardList.vue";
interface propsType {
  isScroll: boolean;
  videos: any;
}
const props = defineProps<propsType>();

onMounted(() => {
  //info.content
});
</script>

<template>
  <NScrollbar v-if="props.isScroll" class="card-space">
    <NSpace wrap>
      <VideoCardList :videos="props.videos" />
    </NSpace>
  </NScrollbar>
  <NSpace v-else wrap>
    <VideoCardList :videos="props.videos" />
  </NSpace>
</template>

<style lang="scss" scoped>
.el-card {
  --el-card-padding: 0;
}

.card-space {
  overflow: scroll-y;
  max-height: calc(100vh - 60px);
}

.box-card {
  margin-top: 15px;
  margin-left: 10px;

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

  .el-card__body {
    margin-top: -15px;

    .el-image {
      width: 100%;
      height: 30vh;
    }

    .play-btn {
      width: 100%;
      text-align: center;
      position: relative;
      bottom: 125px;
      top: 35.5%;
      z-index: 2;
      font-size: 50px;
    }
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
