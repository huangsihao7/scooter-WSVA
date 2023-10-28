<!--
 * @Author: Xu Ning
 * @Date: 2023-10-28 12:30:12
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-28 13:03:47
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\components\myinfo\myHeaderCom.vue
-->
<script setup lang="ts">
import { userStore } from '@/stores/user'
const avatar = userStore().avatar
const username = userStore().username


</script>

<template>
    <div class="demo-type">
      <div>
        <el-avatar :size="150"
          :src="avatar"
        />
      </div>
      {{ username }}
    </div>
</template>
  

  <style scoped>
  .demo-type {
    display: flex;
  }
  .demo-type > div {
    flex: 1;
    text-align: center;
  }
  
  .demo-type > div:not(:last-child) {
    border-right: 1px solid var(--el-border-color);
  }

  .el-avatar{
    float: left;
  }
  </style>
  