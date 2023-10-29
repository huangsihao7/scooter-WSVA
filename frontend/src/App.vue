<!--
 * @Author: Xu Ning
 * @Date: 2023-10-22 19:33:20
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-27 20:26:10
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\App.vue
-->
<script setup lang="ts">
import { RouterView } from "vue-router";
import SubMenu from "./components/SubMenu.vue";
import HeaderMenu from "./components/HeaderMenu.vue";

// import HomeView from './view/HomeView.vue'
</script>

<template>
  <div class="common-layout">
    <ElContainer class="common-layout">
      <ElHeader><HeaderMenu /></ElHeader>
      <ElContainer>
        <ElAside><SubMenu /></ElAside>
        <ElMain><RouterView /></ElMain>
      </ElContainer>
    </ElContainer>
  </div>
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
