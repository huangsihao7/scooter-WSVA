<!--
 * @Author: huangsihao7
 * @Date: 2023-10-30 11:17:41
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-05 22:49:43
 * @FilePath: \scooter-WSVA\frontend\src\components\live\LiveCardListCom.vue
 * @Description: 卡片形式的视频信息展示组件
-->
<script lang="ts" setup>
import { onMounted } from "vue";
import { NCard, NEllipsis, NAvatar, NGrid, NGridItem } from "naive-ui";
import { LiveType } from "@/apis/interface";
import { useRouter } from "vue-router";

interface propsType {
  videos: any;
  // videos: Array<LiveType>;
}

const props = defineProps<propsType>();
const router = useRouter();

// 跳转直播
const goLive = (info: LiveType) => {
  router.push({
    name: "live",
    params: { id: info.uid, url: info.live_play_url },
  });
};

// 跳转用户信息
const goUserInfo = (info: LiveType) => {
  router.push({ name: "userinfo", params: { id: info.uid } });
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
        v-lazy="info.live_cover_url"
        class="image-css"
        alt="img"
        @click="goLive(info)"
      />
    </template>
    <div class="card-footer">
      <NGrid item-responsive :x-gap="12" responsive="screen" cols="4">
        <NGridItem>
          <NAvatar
            class="avatar"
            :src="info.avatar"
            round
            @click="goUserInfo(info)"
          />
        </NGridItem>
        <NGridItem span="3">
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
</style>
