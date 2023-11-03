<!--
 * @Author: Xu Ning
 * @Date: 2023-10-27 14:13:32
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-03 18:51:06
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\components\UserCard.vue
-->
<script lang="ts" setup>
import { ref, onMounted, watch, reactive } from "vue";
import {
  useMessage,
  NCard,
  NAvatar,
  NButton,
  NGrid,
  NGi,
  NEmpty,
  NEllipsis,
NIcon,
} from "naive-ui";
import { useRouter } from "vue-router";
import { FollowCardType } from "@/apis/interface";
import {
  getFollowsList,
  getFollowersList,
  getFriendsList,
  canclefollowOne,
  followOne,
} from "@/apis/follow";
import { routeStore } from "@/stores/route";
import { PersonAddOutline } from "@vicons/ionicons5";
interface propsType {
  userId: number;
}
const props = defineProps<propsType>();
const message = useMessage();
const router = useRouter();
// 用户列表渲染数据
const usersList = reactive<Array<FollowCardType>>([]);
// 卡片类型
const cardType = ref<string>();

// 获取路由最后一个数
function getLastSegmentFromRoute(route: string): string {
  const segments = route.split("/");
  return segments[segments.length - 1];
}

// 获取路由倒数第二个数
function getSecondSegmentFromRoute(route: string): string {
  const segments = route.split("/");
  return segments[segments.length - 2];
}

// 监听路由变化信息
watch(
  () => routeStore().name,
  (newValue: any) => {
    getUserCardsInfo(newValue);
  },
);

// 获取用户卡片信息接口
const getUserCardsInfo = (finalRoute: string) => {
  let userId = props.userId;
  switch (finalRoute) {
    case "friends":
      getFriendsList(userId).then((res: any) => {
        usersList.splice(0, usersList.length, ...res.list);
        console.log(usersList);
        if (usersList) {
          usersList.forEach((element: any) => {
            element.isfriends = true;
          });
        }
      });
      break;
    case "followers":
      getFollowersList(userId).then((res: any) => {
        usersList.splice(0, usersList.length, ...res.list);
      });
      break;
    case "follows":
      getFollowsList(userId).then((res: any) => {
        usersList.splice(0, usersList.length, ...res.list);
        if (usersList) {
          usersList.forEach((element: any) => {
            element.is_follow = true;
          });
        }
      });
      break;
    case "follow":
      getFollowsList(userId).then((res: any) => {
        usersList.splice(0, usersList.length, ...res.list);
        if (usersList) {
          usersList.forEach((element: any) => {
            element.is_follow = true;
          });
        }
      });
      break;
  }
};

// 获取关注的人的信息卡片
onMounted(() => {
  let path = window.location.href;
  console.log(path);
  let lastRoute = getLastSegmentFromRoute(path);
  let secondRoute = getSecondSegmentFromRoute(path);
  let finalRoute = "";
  if (lastRoute === "follow") {
    finalRoute = lastRoute;
  } else if (
    secondRoute === "friends" ||
    secondRoute === "followers" ||
    secondRoute === "follows"
  ) {
    finalRoute = secondRoute;
  }
  routeStore().name = finalRoute;
  cardType.value = finalRoute;
  getUserCardsInfo(finalRoute);
});

// 跳转视频
const handleShowVedio = (info: FollowCardType) => {
  const video_id = info.video_id;
  routeStore().name = "video";
  router.push({ name: "video", params: { id: video_id } });
};

// 跳转到关注人的页面
const handleShowUser = (userId: number) => {
  // routeStore().name = "userinfo";//不能写，写了回退有问题
  router.push({ name: "userinfo", params: { id: userId } });
};

// 取消关注
const cancleFollow = (item: any, _index: any) => {
  if (item.is_follow) {
    canclefollowOne(item.id).then(() => {
      message.success("取消关注成功");
      window.location.reload();
    });
    item.is_follow = false;
  } else {
    followOne(item.id).then(() => {
      message.success("关注成功");
    });
    item.is_follow = true;
  }
};
</script>

<template>
  <NGrid
    v-if="usersList"
    class="space"
    :x-gap="12"
    cols="2 s:3 m:4 l:5 xl:6 2xl:7"
    responsive="screen"
  >
    <NGi v-for="(info, index) in usersList" :key="index">
      <NCard class="card" style="padding: 0" :hoverable="true">
        <template #cover>
          <img
            v-if="info.cover_url"
            class="img"
            :src="info.cover_url"
            @click="handleShowVedio(info)"
          />
          <NEmpty
            v-else
            class="empty"
            description="Ta最近没有发布视频哦"
          ></NEmpty>
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
            <span class="name">
              <NEllipsis expand-trigger="click" line-clamp="1" :tooltip="false">
                {{ info.name }}
              </NEllipsis>
            </span>
            <NButton
              v-if="cardType == 'follows'"
              class="btn"
              @click="cancleFollow(info, index)"
            >
              {{ info.is_follow ? "已关注" : "关注" }}
            </NButton>
            <NButton
              v-if="cardType == 'follow'"
              class="btn"
              @click="cancleFollow(info, index)"
            >
              {{ info.is_follow ? "已关注" : "关注" }}
            </NButton>
            <NButton
              v-if="cardType == 'friends'"
              class="btn"
              @click="cancleFollow(info, index)"
            >
              {{ info.is_friends ? "已互关" : "关注" }}
            </NButton>
            <p class="sig">{{ info.signature }}</p>
          </div>
        </div>
      </NCard>
    </NGi>
  </NGrid>
  <NEmpty description="没有关注的用户哦~去别处看看吧~">
    <template #icon>
      <NIcon>
        <PersonAddOutline />
      </NIcon>
    </template>
  </NEmpty>
</template>

<style lang="scss" scoped>
.space {
  margin: 2vh 2vw;
}
.card {
  height: 35vh;
  width: 100%;

  .header-info {
    display: flex;
    margin-top: 20px;
    .other-info {
      margin-left: 16px;
      .name {
        font-weight: bold;
        font-size: 1.2rem;
      }
      .btn {
        margin-left: 16px;
        position: absolute;
        bottom: 10px;
        right: 10px;
      }
      .sig {
        text-align: left;
      }
    }
  }

  .img {
    width: 100%;
    height: calc(30vh - 40px);
    object-fit: fill;
  }
  .empty {
    width: 100%;
    height: calc(22vh - 40px);
    position: relative;
    top: 30%;
  }

}
</style>
