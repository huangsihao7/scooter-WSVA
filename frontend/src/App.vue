<!--
 * @Author: Xu Ning
 * @Date: 2023-10-22 19:33:20
 * @LastEditors: huangsihao7 1057434651@qq.com
 * @LastEditTime: 2023-10-30 10:48:50
 * @Description: 
 * @FilePath: /scooter-WSVA/frontend/src/App.vue
-->
<script setup lang="ts">
import { RouterView } from "vue-router";
import SubMenu from "./components/SubMenu.vue";
import HeaderMenu from "./components/HeaderMenu.vue";
import {
  NLayout,
  NLayoutHeader,
  NLayoutSider,
  NMessageProvider,
} from "naive-ui";
</script>

<template>
  <!-- App.vue -->
  <NMessageProvider>
    <div class="common-layout">
      <NLayout class="common-layout">
        <NLayoutHeader>
          <HeaderMenu />
        </NLayoutHeader>
        <NLayout has-sider>
          <NLayoutSider>
            <SubMenu />
          </NLayoutSider>
          <NLayout>
            <RouterView />
          </NLayout>
        </NLayout>
      </NLayout>
    </div>
  </NMessageProvider>
</template>

<style scoped>
.common-layout {
  width: 100vw;
  height: 100vh;
}

.el-header {
  padding: 0;
  z-index: 10;
}

.el-aside {
  height: 100%;
  width: 160px;
}

.el-aside {
  z-index: 2;
}

.el-main {
  padding: 0;
  background-color: #ffffff;
  background-image: linear-gradient(133deg, #ffffff 12%, #d6e4ff 100%);
}
</style>
