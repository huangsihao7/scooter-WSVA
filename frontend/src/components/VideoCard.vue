<!--
 * @Author: Xu Ning
 * @Date: 2023-10-27 14:13:32
 * @LastEditors: huangsihao7 1057434651@qq.com
 * @LastEditTime: 2023-10-30 10:37:16
 * @Description: 各个视频分类的视频卡片渲染
 * @FilePath: /scooter-WSVA/frontend/src/components/VideoCard.vue
-->
<script lang="ts" setup>
import { onMounted } from "vue";
import { NTag, NIcon, NCard, NSpace, NImage, NScrollbar } from "naive-ui";
import { Play, Heart } from "@vicons/ionicons5";
import { NEllipsis } from "naive-ui";
interface propsType {
  isScroll: boolean;
  videos: any;
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
  <NScrollbar v-if="props.isScroll" class="card-space">
    <NSpace wrap>
      <NCard
        v-for="(info, index) in props.videos"
        :key="index"
        class="box-card"
        style="width: calc((100vw - 260px) / 4)"
        @click="GetVideoLink"
      >
        <div style="position: relative">
          <NImage
            :src="info.coverUrl"
            width="100"
            preview-disabled
            @click="handleShowVedio"
          />
          <NTag class="time" round :bordered="false" type="info">
            7：26
            <template #icon>
              <NIcon color="#fff" :component="Play" />
            </template>
          </NTag>
          <NTag class="like" round :bordered="false" type="error">
            {{ info.favoriteCount }}
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
            <span>@ {{ info.author.name }}</span>
            <span class="date">7月8日</span>
          </div>
        </div>
      </NCard>
    </NSpace>
  </NScrollbar>
  <NSpace v-else wrap>
    <NCard
      v-for="(info, index) in props.videos"
      :key="index"
      class="box-card"
      style="width: calc((100vw - 260px) / 4)"
    >
      <div style="position: relative" @click="GetVideoLink">
        <NImage
          :src="info.coverUrl"
          width="100"
          preview-disabled
          @click="handleShowVedio"
        />
        <NTag class="time" round :bordered="false" type="info">
          7：26
          <template #icon>
            <NIcon color="#fff" :component="Play" />
          </template>
        </NTag>
        <NTag class="like" round :bordered="false" type="error">
          {{ info.favoriteCount }}
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
          <span>@ {{ info.author.name }}</span>
          <span class="date">7月8日</span>
        </div>
      </div>
    </NCard>
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
