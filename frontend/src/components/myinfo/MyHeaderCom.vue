<!--
 * @Author: Xu Ning
 * @Date: 2023-10-28 12:30:12
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-30 20:29:49
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\components\myinfo\myHeaderCom.vue
-->
<script setup lang="ts">
import { ref, onMounted } from "vue";
import { userStore } from "@/stores/user";
import { getUserInfo } from "@/apis/user";
import { NButton, NIcon } from "naive-ui";
import InfoEditCom from "./InfoEditCom.vue";
import { CashOutline as CashIcon } from "@vicons/ionicons5";

interface propsType {
  userId: number;
}

const props = defineProps<propsType>();
const avatar = userStore().avatar;
const userInfo = ref<any>({
  background_image: userStore().avatar,
});
const editVisible = ref<boolean>(false);

// 获取用户信息
const getUserInfoFunc = () => {
  getUserInfo(props.userId).then((res: any) => {
    userInfo.value = res.user;
  });
};

// 编辑资料开启
const editInfo = () => {
  editVisible.value = true;
};

// 更新编辑资料是否可见
const UpdateVisible = () => {
  editVisible.value = false;
};

onMounted(() => {
  console.log(props.userId);
  getUserInfoFunc();
});
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
    <ElRow>
      <ElCol :span="4">
        <ElAvatar :src="avatar" />
      </ElCol>
      <ElCol v-if="userInfo" :span="20" class="info-tab">
        <ElText tag="b">{{ userInfo.name }}</ElText>
        <ElText tag="p">{{ userInfo.signature }}</ElText>
        <div class="follow">
          <NButton color="#606266" text>
            关注 {{ userInfo.follow_count }}
          </NButton>
          <ElDivider direction="vertical" />
          <NButton color="#606266" text>
            粉丝 {{ userInfo.follower_count }}
          </NButton>
          <ElDivider direction="vertical" />
          <NButton color="#606266" text>
            获赞 {{ userInfo.favorite_count }}
          </NButton>
        </div>
        <NButton
          strong
          round
          class="edit-info"
          color="#409eff85"
          @click="editInfo"
        >
          <template #icon>
            <NIcon>
              <CashIcon />
            </NIcon>
          </template>
          编辑资料
        </NButton>
      </ElCol>
    </ElRow>
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

  .el-avatar {
    float: right;
    font-size: 5rem;
    width: calc((80vw - 260px) / 6);
    // height: calc((80vw - 260px) / 6);
    // height: calc((90vh - 70px) / 5);
    height: calc((80vw - 260px) / 6);
  }

  .info-tab {
    display: flex;
    flex-direction: column;
    justify-content: space-evenly;

    .edit-info {
      position: absolute;
      top: 0;
      right: 20px;
    }

    .el-text,
    .follow {
      margin: 10px 20px;
      display: block;
    }

    .el-text {
      width: calc(100% - 40px);
      margin: 10px 20px;
    }

    .follow {
      .n-button {
        padding-right: 10px;
      }
    }

    .el-text:nth-child(1) {
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
