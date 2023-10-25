<!--
 * @Author: Xu Ning
 * @Date: 2023-10-25 16:22:40
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-25 20:08:52
 * @Description: 
 * @FilePath: \scooter-wsva\frontend\src\components\Menu.vue
-->
<script lang="ts" setup>
import { ref } from 'vue'
import { 
  VideoCameraFilled,
  CirclePlus
 } from '@element-plus/icons-vue'
import { userStore } from '@/stores/user';

const activeIndex = ref('1')
const handleSelect = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
const store = userStore()
var avatar : string = 'https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png'

</script>

<template>
  <el-menu
    :default-active="activeIndex"
    class="el-menu-demo"
    mode="horizontal"
    :ellipsis="false"
    @select="handleSelect"
  >
    <el-menu-item index="0">
      <el-icon><VideoCameraFilled /></el-icon>
      See World
    </el-menu-item>
    <div class="flex-grow" />
    <el-menu-item index="1">
      <el-icon>
        <CirclePlus/>
      </el-icon>
      投稿
    </el-menu-item>
    <el-sub-menu index="2" v-if="avatar">
      <template #title>
        <el-avatar
          :src="avatar"
        />
      </template>
      <el-menu-item index="logout">退出登录</el-menu-item>
    </el-sub-menu>
    <el-menu-item index="login" v-else>
      登录
    </el-menu-item>
  </el-menu>
</template>



<style>
.flex-grow {
  flex-grow: 1;
}
</style>
