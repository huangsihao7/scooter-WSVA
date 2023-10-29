<!--
 * @Author: Xu Ning
 * @Date: 2023-10-27 14:13:32
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-29 16:14:25
 * @Description: 各个视频分类的视频卡片渲染
 * @FilePath: \scooter-WSVA\frontend\src\components\VideoCard.vue
-->
<script lang="ts" setup>
import { reactive, onMounted } from "vue";
import { NTag, NIcon } from "naive-ui";
import { Play, Heart } from "@vicons/ionicons5";
import { NEllipsis } from "naive-ui";
interface propsType {
  isScroll: boolean;
}
const props = defineProps<propsType>();

const res: any = reactive([
  {
    avatar: "http://127.0.0.1:8080/OIP-C.jpg",
    url: "http://127.0.0.1:8080/1.jpg",
    name: "我今天好帅",
    content:
      "奇奇怪怪的视频有增加了s9.5最强空城玩法-空城司令莎狗！兄弟们点赞收藏赶紧冲",
    isfollowed: true,
  },
  {
    avatar: "http://127.0.0.1:8080/OIP-C.jpg",
    url: "http://127.0.0.1:8080/2.jpg",
    name: "吃大餐饱了",
    content: "明日复明日，ooooo",
    isfollowed: true,
  },
  {
    avatar: "http://127.0.0.1:8080/OIP-C.jpg",
    url: "http://127.0.0.1:8080/3.jpg",
    name: "吃大餐饱了",
    content: "明日复明日，ooooo",
    isfollowed: true,
  },
  {
    avatar: "http://127.0.0.1:8080/OIP-C.jpg",
    url: "http://127.0.0.1:8080/18.jpg",
    name: "原皮",
    content: "明日复明日，ooooo",
    isfollowed: true,
  },
  {
    avatar: "http://127.0.0.1:8080/OIP-C.jpg",
    url: "http://127.0.0.1:8080/5.jpg",
    name: "大师傅",
    content: "学习是终身事业！！！！！！！！！",
    isfollowed: true,
  },
  {
    avatar: "http://127.0.0.1:8080/OIP-C.jpg",
    url: "http://127.0.0.1:8080/6.jpg",
    name: "大发噶是否",
    content: "学习是终身事业！！！！！！！！！",
    isfollowed: true,
  },
  {
    avatar: "http://127.0.0.1:8080/OIP-C.jpg",
    url: "http://127.0.0.1:8080/7.jpg",
    name: "爱上的人服务",
    content: "学习是终身事业！！！！！！！！！",
    isfollowed: true,
  },
  {
    avatar: "http://127.0.0.1:8080/OIP-C.jpg",
    url: "http://127.0.0.1:8080/8.jpg",
    name: "暴怒改变",
    content: "学习是终身事业！！！！！！！！！",
    isfollowed: true,
  },
  {
    avatar: "http://127.0.0.1:8080/OIP-C.jpg",
    url: "http://127.0.0.1:8080/9.jpg",
    name: "大师傅",
    content: "学习是终身事业！！！！！！！！！",
    isfollowed: true,
  },
  {
    avatar: "http://127.0.0.1:8080/OIP-C.jpg",
    url: "http://127.0.0.1:8080/10.jpg",
    name: "大发噶是否",
    content: "学习是终身事业！！！！！！！！！",
    isfollowed: true,
  },
  {
    avatar: "http://127.0.0.1:8080/OIP-C.jpg",
    url: "http://127.0.0.1:8080/11.jpg",
    name: "爱上的人服务",
    content: "学习是终身事业！！！！！！！！！",
    isfollowed: true,
  },
  {
    avatar: "http://127.0.0.1:8080/OIP-C.jpg",
    url: "http://127.0.0.1:8080/12.jpg",
    name: "暴怒改变",
    content: "学习是终身事业！！！！！！！！！",
    isfollowed: true,
  },
]);
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
  <ElScrollbar v-if="props.isScroll" class="card-space">
    <ElSpace wrap>
      <ElCard
        v-for="(info, index) in res"
        :key="index"
        class="box-card"
        style="width: calc((100vw - 260px) / 4)"
        @click="GetVideoLink"
      >
        <div style="position: relative">
          <ElImage :src="info.url" fit="fill" @click="handleShowVedio" />
          <NTag class="time" round :bordered="false" type="info">
            7：26
            <template #icon>
              <NIcon color="#fff" :component="Play" />
            </template>
          </NTag>
          <NTag class="like" round :bordered="false" type="error">
            10
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
              {{ info.content }}
            </NEllipsis>
          </div>
          <div class="name">
            <span>@ {{ info.name }}</span>
            <span class="date">7月8日</span>
          </div>
        </div>
      </ElCard>
    </ElSpace>
  </ElScrollbar>
  <ElSpace v-else wrap>
    <ElCard
      v-for="(info, index) in res"
      :key="index"
      class="box-card"
      style="width: calc((100vw - 260px) / 4)"
    >
      <div style="position: relative" @click="GetVideoLink">
        <ElImage :src="info.url" fit="fill" @click="handleShowVedio" />
        <NTag class="time" round :bordered="false" type="info">
          7：26
          <template #icon>
            <NIcon color="#fff" :component="Play" />
          </template>
        </NTag>
        <NTag class="like" round :bordered="false" type="error">
          10
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
            {{ info.content }}
          </NEllipsis>
        </div>
        <div class="name">
          <span>@ {{ info.name }}</span>
          <span class="date">7月8日</span>
        </div>
      </div>
    </ElCard>
  </ElSpace>
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
