<!--
 * @Author: Xu Ning
 * @Date: 2023-10-28 12:30:12
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-28 12:57:31
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\components\myinfo\myHeaderCom.vue
-->
<template>
    <div class="demo-type">
      <div>
        <el-avatar
          src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
        />
      </div>
      
    </div>
  </template>
  
  <script setup lang="ts">
  </script>
  
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
  </style>
  