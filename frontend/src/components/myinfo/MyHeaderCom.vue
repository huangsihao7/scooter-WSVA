<!--
 * @Author: Xu Ning
 * @Date: 2023-10-28 12:30:12
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-02 19:09:06
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\components\myinfo\MyHeaderCom.vue
-->
<script setup lang="ts">
import { onMounted, ref } from "vue";
import { userStore } from "@/stores/user";
import { getUserInfo } from "@/apis/user";
import { NAvatar, NButton, NDivider, NGrid, NGridItem, NIcon } from "naive-ui";
import InfoEditCom from "./InfoEditCom.vue";
import { CashOutline as CashIcon } from "@vicons/ionicons5";
import { useRouter } from "vue-router";
import { routeStore } from "@/stores/route";

interface propsType {
  userId: number;
}

const props = defineProps<propsType>();
// 用户头像信息
const userInfo = ref<any>({
});
const editVisible = ref<boolean>(false);
const router = useRouter();

// 获取用户信息
const getUserInfoFunc = () => {
  // 别删 很诡异的错误 userId会显示string，必须转两下才能变成int
  let uid = props.userId.toString();
  let uid_num = parseInt(uid);
  getUserInfo(uid_num).then((res: any) => {
    userInfo.value = res.user;
    let routeName = routeStore().name;
    if(routeName == 'user'){
      userStore().name = res.user.name;
      userStore().avatar = res.user.avatar;
      userStore().gender = res.user.gender;
      userStore().signature = res.user.signature;
      userStore().background_image = res.user.background_image;
    }
  });
};

// 更新编辑资料是否可见
const UpdateVisible = () => {
  editVisible.value = false;
};

onMounted(() => {
  getUserInfoFunc();
});

// 跳转到关注的人的列表
const goFollowing = () => {
  routeStore().name = "follows";
  router.push({ name: "follows" , params: { id: props.userId } });
};

// 跳到粉丝列表
const goFollowers = () =>{
  routeStore().name = "followers";
  router.push({ name: "followers" , params: { id: props.userId } });
}

// 跳到朋友列表
const goFriends = () =>{
  routeStore().name = "friends";
  router.push({ name: "friends" , params: { id: props.userId } });
}
</script>

<template>
  <div
    class="header"
    :style="{
      'background-image':
        'linear-gradient( to right, rgba(255, 255, 255, 1), rgba(255, 255, 255, 0)), url(' +
        userInfo.background_image +
        ')',
    }"
  >
    <NGrid>
      <NGridItem :span="4">
        <NAvatar :src="userInfo.avatar" round />
      </NGridItem>
      <NGridItem v-if="userInfo" :span="20" class="info-tab">
        <p class="name">{{ userInfo.name }}</p>
        <p class="signature">{{ userInfo.signature }}</p>
        <div class="follow">
          <NButton color="#606266" text @click="goFollowing">
            关注 {{ userInfo.follow_count }}
          </NButton>
          <NDivider vertical />
          <NButton color="#606266" text @click="goFollowers">
            粉丝 {{ userInfo.follower_count }}
          </NButton>
          <NDivider vertical />
          <NButton color="#606266" text @click="goFriends">
            朋友 {{ userInfo.follower_count }}
          </NButton>
          <NDivider vertical />
          <NButton color="#606266" text >
            获赞 {{ userInfo.favorite_count }}
          </NButton>
        </div>
        <NButton
          strong
          round
          class="edit-info"
          color="#409eff85"
          @click="editVisible = true"
        >
          <template #icon>
            <NIcon>
              <CashIcon />
            </NIcon>
          </template>
          编辑资料
        </NButton>
      </NGridItem>
    </NGrid>
    <InfoEditCom
      v-if="userInfo"
      :user-info="userInfo"
      :is-visible="editVisible"
      @visible-update="UpdateVisible"
    />
  </div>
</template>

<style scoped lang="scss">
.header {
  background-color: aquamarine;
  text-align: left;
  display: flex;
  padding: 3vh 0;
  border-radius: 25px;
  background: no-repeat center top / 100% 100%;

  .n-avatar {
    float: right;
    font-size: 5rem;
    width: calc((80vw - 260px) / 6);
    height: calc((80vw - 260px) / 6);
  }

  .info-tab {
    display: flex;
    flex-direction: column;
    justify-content: space-evenly;

    .edit-info {
      position: absolute;
      top: 70px;
      right: calc(10vw + 20px);
    }

    p,
    .follow {
      margin: 10px 20px;
      display: block;
    }

    p {
      width: calc(100% - 40px);
      margin: 10px 20px;
    }

    .follow {
      .n-button {
        padding-right: 10px;
      }
    }

    p:nth-child(1) {
      color: black;
      font-size: 2.5rem;
      font-weight: bold;
    }
  }
}

.header > div {
  flex: 1;
}

.header > div:not(:last-child) {
  border-right: 1px solid var(--el-border-color);
}
</style>
