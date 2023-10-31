<!--
 * @Author: Xu Ning
 * @Date: 2023-10-22 19:33:20
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-31 19:56:57
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\App.vue
-->
<script setup lang="ts">
import { RouterView } from "vue-router";
import SubMenu from "./components/SubMenu.vue";
import HeaderMenu from "./components/HeaderMenu.vue";
import {
NConfigProvider,
  NLayout,
  NLayoutHeader,
  NLayoutSider,
  NMessageProvider,
} from "naive-ui";
import themeOverrides from './theme'
</script>

<template>
  <!-- App.vue -->
  <NConfigProvider :theme-overrides="themeOverrides">
    <NMessageProvider>
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

.main {
  padding: 0;
  background-color: #ffffff;
  background-image: linear-gradient(133deg, #ffffff 12%, #d6e4ff 100%);
}
</style>
