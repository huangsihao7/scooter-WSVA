<!--
 * @Author: Xu Ning
 * @Date: 2023-10-27 14:13:32
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-30 20:10:40
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\components\UserCard.vue
-->
<script lang="ts" setup>
import { ref, onMounted } from "vue";
import { userStore } from "@/stores/user";
import { getFollowList, canclefollowOne, followOne } from "@/apis/follow";
import { useRouter } from "vue-router";
import { useMessage, NCard, NAvatar, NButton, NGrid, NGi } from "naive-ui";

const message = useMessage();
const router = useRouter();
const folllowList: any = ref([]);

// 获取关注的人的信息卡片
onMounted(() => {
  getFollowList(userStore().user_id).then((res: any) => {
    folllowList.value = res.list;
    folllowList.value.forEach((element: any) => {
      element.isfollowed = true;
    });
  });
});

const handleShowVedio = () => {
  console.log("show");
};

// 跳转到关注人的页面
const handleShowUser = (userId: number) => {
  router.push({ name: "following", params: { id: userId } });
};

const cancleFollow = (item: any, _index: any) => {
  if (item.isfollowed) {
    canclefollowOne(item.id).then((res: any) => {
      if (res.status_code == 200) {
        message.success("取消关注成功");
      } else {
        message.error(res.status_msg);
      }
    });
    item.isfollowed = false;
  } else {
    followOne(item.id).then((res: any) => {
      if (res.status_code == 200) {
        message.success("关注成功");
      } else {
        message.error(res.status_msg);
      }
    });
    item.isfollowed = true;
  }

  // item.isfollowed = false
  // TODO 取消关注接口
};
</script>

<template>
  <!-- <ElSpace wrap> -->
  <NGrid class="space" x-gap="12" :cols="4">
    <NGi v-for="(info, index) in folllowList" :key="index">
      <NCard class="card" style="padding: 0" :hoverable="true">
        <template #cover>
          <img
            class="img"
            :src="info.background_image"
            @click="handleShowVedio"
          />
        </template>
        <div class="header-info">
          <NAvatar
            class="avatar"
            round
            :size="60"
            :src="info.avatar"
            @click="handleShowUser(info.id)"
          />
          <div class="other-info">
            <span class="name">{{ info.name }}</span>
            <NButton class="btn" @click="cancleFollow(info, index)">
              {{ info.isfollowed == true ? "已关注" : "关注" }}
            </NButton>
            <p class="sig">{{ info.dec }}</p>
          </div>
        </div>
      </NCard>
    </NGi>
  </NGrid>

  <!-- <ElCard
      v-for="(info, index) in folllowList"
      :key="index"
      class="box-card"
      style="width: calc((100vw - 260px) / 4)"
    >
      <template #header>
        <div class="card-header">
          <div class="avatar">
            <ElAvatar :size="60" :src="info.avatar" />
          </div>
          <div>
            <span>{{ info.name }}</span>
            <ElButton @click="cancleFollow(info, index)">{{
              info.isfollowed == true ? "已关注" : "关注"
            }}</ElButton>
            <p>{{ info.dec }}</p>
          </div>
        </div>
      </template>
      <div>
        <ElImage
          :src="info.background_image"
          fit="fill"
          @click="handleShowVedio"
        />
      </div>
    </ElCard> -->
  <!-- </ElSpace> -->
</template>

<style lang="scss" scoped>
.space {
  margin: 2vh 2vw;
}
.card {
  height: 30vh;
  width: 100%;

  .header-info {
    display: flex;
    margin-top: 20px;
    height: 8vh;
    .other-info {
      margin-left: 16px;
      .name {
        font-weight: bold;
        font-size: 1.2rem;
      }
      .btn {
        margin-left: 16px;
      }
      .sig {
        text-align: left;
      }
    }
  }

  .img {
    width: 100%;
    height: calc(22vh - 40px);
    object-fit: fill;
  }
}

.box-card {
  margin-top: 15px;
  margin-left: 10px;

  .el-card__body {
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

  .card-header {
    margin-top: 15px;
    display: flex;
    justify-content: left;

    .avatar,
    span,
    p,
    .el-button {
      margin-left: 10px;
    }

    span {
      font-weight: bold;
      font-size: 1.2rem;
    }

    p {
      font-size: 0.8rem;
    }
  }
}
</style>
