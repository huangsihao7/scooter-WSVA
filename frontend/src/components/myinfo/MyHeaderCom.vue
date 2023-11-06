<!--
 * @Author: Xu Ning
 * @Date: 2023-10-28 12:30:12
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-06 20:58:54
 * @Description: 我的页面的个人信息展示组件
 * @FilePath: \scooter-WSVA\frontend\src\components\myinfo\MyHeaderCom.vue
-->
<script setup lang="ts">
import { onMounted, ref } from "vue";
import { userStore } from "@/stores/user";
import { getUserInfo } from "@/apis/user";
import { NAvatar, NButton, NDivider, NGrid, NGridItem, useMessage } from "naive-ui";
import InfoEditCom from "./InfoEditCom.vue";
import { useRouter } from "vue-router";
import { routeStore } from "@/stores/route";
import { doFollow } from '@/apis/follow'

interface propsType {
  userId: number;
}

const props = defineProps<propsType>();
// 用户头像信息
const userInfo = ref<any>({});
const editVisible = ref<boolean>(false);
const router = useRouter();
// 是否关注用户
const isFollow = ref<boolean>(false)

// 获取用户信息
const getUserInfoFunc = () => {
  // 别删 很诡异的错误 userId会显示string，必须转两下才能变成int
  let uid = props.userId.toString();
  let uid_num = parseInt(uid);
  getUserInfo(uid_num).then((res: any) => {
    userInfo.value = res.user;
    let routeName = routeStore().name;
    isFollow.value = res.user.is_follow;
    if (routeName == "user") {
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
  router.push({ name: "follows", params: { id: props.userId } });
};

// 跳到粉丝列表
const goFollowers = () => {
  routeStore().name = "followers";
  router.push({ name: "followers", params: { id: props.userId } });
};

// 跳到朋友列表
const goFriends = () => {
  routeStore().name = "friends";
  router.push({ name: "friends", params: { id: props.userId } });
};

const message = useMessage()
// 更新关注状态
const changeFollowState = () =>{
  let uid = parseInt(props.userId.toString())
  if(isFollow.value){
    isFollow.value = false;
    // 2 取关
    doFollow(uid,2).then(()=>{
      message.success('取关成功')
    })
  }else{
    isFollow.value = true;
    doFollow(uid,1).then(()=>{
      message.success('关注成功')
    })
  }
}
</script>

<template>
  <div
    class="header"
    :style="{
      'background-image':
        'linear-gradient( 100deg , rgba(255, 255, 255, 1), rgba(255, 255, 255, 0)), url(' +
        userInfo.background_image +
        ')',
    }"
  >
    <div class="sub-header">
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
              朋友 {{ userInfo.friend_count }}
            </NButton>
            <NDivider vertical />
            <NButton color="#606266" text>
              获赞 {{ userInfo.favorite_count }}
            </NButton>
          </div>
        </NGridItem>
      </NGrid>
      <NButton
        v-if="userInfo.id == userStore().user_id"
        strong
        round
        class="edit-info"
        color="#409eff85"
        @click="editVisible = true"
      >
        编辑资料
        <template #icon>
          <span class="iconfont icon-bianji"></span>
        </template>
      </NButton>
      <NButton
        v-else
        strong
        round
        class="edit-info"
        color="#409eff85"
        @click="changeFollowState"
      >
        {{ isFollow?'已关注':'关注' }}
      </NButton>
      <InfoEditCom
        v-if="userInfo"
        :user-info="userInfo"
        :is-visible="editVisible"
        @visible-update="UpdateVisible"
      />
    </div>
  </div>
</template>

<style scoped lang="scss">
.edit-info {
  // position: absolute;
  // top: 70px;
  // right: calc(10vw + 20px);
  // float: right;
  margin-right: 20px;
}

.header {
  text-align: left;
  display: flex;
  // padding: 2vh 0;
  // border-radius: 25px;
  background: no-repeat center top / 100% 100%;
  .sub-header {
    text-align: left;
    display: flex;
    background: #ffffff24;
    backdrop-filter: blur(3px);
    // border-radius: 25px;
    // margin: 0 2vw;
    padding: 3vh 2vw;
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
}

.header > div {
  flex: 1;
}

.header > div:not(:last-child) {
  border-right: 1px solid var(--el-border-color);
}
</style>
