<template>
  <el-row class="tac">
    <el-col>
      <el-menu
        :default-active="active"
        class="el-menu-vertical-demo"
        @open="handleOpen"
        @close="handleClose"
        router="true"
      >
        <el-menu-item index="rec">
          <div class="submenu-text">
            <el-icon><Star /></el-icon>
            <span>推荐</span>
          </div>
        </el-menu-item>
        <el-menu-item index="follow">
          <div class="submenu-text">
            <el-icon><Connection /></el-icon>
            <span>关注</span>
          </div>
        </el-menu-item>
        <el-menu-item index="user">
          <div class="submenu-text">
            <el-icon><House /></el-icon>
            <span>我的</span>
          </div>
        </el-menu-item>
        <el-divider content-position="center" >分类</el-divider>
        <el-menu-item index="hot">
          <div class="submenu-text">
            <el-icon><Medal /></el-icon>
            <span>热门</span>
          </div>
        </el-menu-item>
        <el-menu-item index="recreation">
          <div class="submenu-text">
            <el-icon><Goblet /></el-icon>
            <span>娱乐</span>
          </div>
        </el-menu-item>
        <el-menu-item index="sports">
          <div class="submenu-text">
            <el-icon><Basketball /></el-icon>
            <span>体育</span>
          </div>
        </el-menu-item>
        <el-menu-item index="food">
          <div class="submenu-text">
            <el-icon><KnifeFork /></el-icon>
            <span>美食</span>
          </div>
        </el-menu-item>
        <el-menu-item index="cartoon">
          <div class="submenu-text">
            <el-icon><Lollipop /></el-icon>
            <span>二次元</span>
          </div>
        </el-menu-item>
        <el-menu-item index="knowledge">
          <div class="submenu-text">
            <el-icon><Reading /></el-icon>
            <span>知识</span>
          </div>
        </el-menu-item>
      </el-menu>
    </el-col>
  </el-row>
</template>

<script lang="ts" setup>
import { routeStore } from '@/stores/route';
import { onMounted, ref } from 'vue';
import {
  KnifeFork,
  Goblet,
  Lollipop,
  Connection,
  Medal,
  Reading,
  House,
  Basketball,
  Star
} from '@element-plus/icons-vue'

const active = ref<String>('')

// 获取当前页面的末尾路由
function getLastSegmentFromRoute(route: string): string {
  const segments = route.split('/');
  return segments[segments.length - 1];
}

// 通过url确定menu的activekey
onMounted(() => {
  if(active.value == ''){
    const currentPath = window.location.pathname;
    let path = getLastSegmentFromRoute(currentPath)
    active.value = path
    routeStore().name = path
    console.log('path', path, active.value, routeStore().name)
  }
})

const handleOpen = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
  active.value = key
  routeStore().name = key
}
const handleClose = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
</script>

<style scoped>
.tac{
  height: 100%;
}
.el-menu-vertical-demo{
  height:100%;
}

.el-menu-item .submenu-text{
  margin-left: 20px;
}
</style>
  