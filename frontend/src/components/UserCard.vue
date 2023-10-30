<!--
 * @Author: Xu Ning
 * @Date: 2023-10-27 14:13:32
 * @LastEditors: huangsihao7 1057434651@qq.com
 * @LastEditTime: 2023-10-30 16:12:37
 * @Description: 
 * @FilePath: /scooter-WSVA/frontend/src/components/UserCard.vue
-->
<script lang="ts" setup>
import { ref, onMounted } from "vue";
import { userStore } from "@/stores/user";
import { getFollowList, canclefollowOne, followOne } from "@/apis/follow";
import console from "console";
import { useMessage } from "naive-ui";
const folllowList: any = ref([]);
const message = useMessage();

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
  <ElSpace wrap>
    <ElCard
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
    </ElCard>
  </ElSpace>
</template>

<style lang="scss" scoped>
.el-card {
  --el-card-padding: 0;
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
