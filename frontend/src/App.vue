<!--
 * @Author: Xu Ning
 * @Date: 2023-10-22 19:33:20
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-04 17:48:10
 * @Description: 入口
 * @FilePath: \scooter-WSVA\frontend\src\App.vue
-->
<script setup lang="ts">
import { RouterView } from "vue-router";
import SubMenu from "@/components/menu/SubMenu.vue";
import HeaderMenu from "@/components/menu/HeaderMenu.vue";
import { ref, onMounted } from "vue";
import {
  NConfigProvider,
  NLayout,
  NLayoutHeader,
  NLayoutSider,
  NMessageProvider,
  NModal,
} from "naive-ui";
import themeOverrides from "./theme";
import { userStore } from "./stores/user";
let showModal = ref<boolean>(false);
const onPositiveClick = () => {
  showModal.value = false;
};
onMounted(() => {
  if (!userStore().isLoggedIn) {
    showModal.value = true;
  }
});
</script>

<template>
  <!-- App.vue -->
  <NConfigProvider :theme-overrides="themeOverrides">
    <NMessageProvider placement="top-right">
      <div class="common-layout">
        <NLayout class="common-layout">
          <NLayoutHeader class="header">
            <HeaderMenu />
          </NLayoutHeader>
          <NLayout has-sider class="sider">
            <NLayoutSider width="160">
              <SubMenu />
            </NLayoutSider>
            <NLayout class="main">
              <RouterView />
            </NLayout>
          </NLayout>
        </NLayout>
      </div>
      <NModal
        v-model:show="showModal"
        class="modal-prompt"
        :mask-closable="false"
        preset="dialog"
        title="提醒"
        :show-icon="false"
        :closable="false"
        positive-text="好的"
        @positive-click="onPositiveClick"
      >
        <p>
          1.未登录用户仅可浏览<span>10</span>条视频，并仅开放视频<span>分享</span>和<span>下载</span>功能；
        </p>
        <p>
          2.视频切换支持三种方式：<span>上下滚轮滚动</span>，<span>右下方按钮切换</span>，<span>键盘上下键切换</span>。
        </p>
      </NModal>
    </NMessageProvider>
  </NConfigProvider>
</template>

<style scoped>
.common-layout {
  width: 100vw;
  height: 100vh;
}

.header {
  padding: 0;
  z-index: 10;
}

.sider {
  z-index: 2;
  height: calc(100vh - 60px);
}

/* .modal-prompt{
  width: 50% !important;
} */

.n-layout-header {
  box-shadow: 9px 9px 9px 10px rgb(172 215 255 / 26%);
  background-color: #ffffff;
  /* background-image: linear-gradient(133deg, #ffffff 12%, #d6e4ff 100%); */
}

.n-layout-sider {
  box-shadow: 9px 2px 4px linear-gradient(180deg, #ffffff 0%, #d6e4ff 100%);
  background-color: #ffffff;
  /* background-image: linear-gradient(227deg, #ffffff 42%, #d6e4ff 100%); */
  background-image: linear-gradient(to bottom, #ffffff 30%, #d6e4ff 100%);
}

span {
  font-weight: bold;
}
.main {
  padding: 0;
  /* background: transparent; */
  /* background: no-repeat url(https://p-pc-weboff.byteimg.com/tos-cn-i-9r5gewecjs/test.png) var(--color-bg-b0); */

  background-color: #ffffff;
  /* background-image: linear-gradient(133deg, #ffffff 12%, #b4ceff 100%); */
  background-image: linear-gradient(to bottom, #ffffff 30%, #d6e4ff 100%);
}
</style>
