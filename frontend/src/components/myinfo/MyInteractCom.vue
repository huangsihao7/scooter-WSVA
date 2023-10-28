<!--
 * @Author: Xu Ning
 * @Date: 2023-10-28 12:30:41
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-28 15:57:32
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\components\myinfo\MyInteractCom.vue
-->
<template>
    <n-tabs default-value="work" justify-content="space-evenly" type="line">
        <n-tab-pane name="work" tab="作品">
            <VideoCard :isScroll="false"/>
        </n-tab-pane>
        <n-tab-pane name="favourite" tab="喜欢">
            <VideoCard :isScroll="false"/>
        </n-tab-pane>
        <n-tab-pane name="collect" tab="收藏">
            <VideoCard :isScroll="false"/>
        </n-tab-pane>
        <n-tab-pane name="history" tab="观看历史">
            <VideoCard :isScroll="false"/>
        </n-tab-pane>
    </n-tabs>
</template>
<script lang="ts" setup>
import {  NTabs, NTabPane } from 'naive-ui';
import VideoCard from '../VideoCard.vue';


  </script>
  <style>
  .demo-tabs > .el-tabs__content {
    padding: 32px;
    color: #6b778c;
    font-size: 32px;
    font-weight: 600;
  }
  </style>
  