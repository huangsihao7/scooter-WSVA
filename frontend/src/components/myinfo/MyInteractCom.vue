<!--
 * @Author: Xu Ning
 * @Date: 2023-10-28 12:30:41
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-28 15:57:32
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\components\myinfo\MyInteractCom.vue
-->
<template>
  <NTabs default-value="work" justify-content="space-evenly" type="line">
    <NTabPane name="work" tab="作品">
      <VideoCard :is-scroll="false" />
    </NTabPane>
    <NTabPane name="favourite" tab="喜欢">
      <VideoCard :is-scroll="false" />
    </NTabPane>
    <NTabPane name="collect" tab="收藏">
      <VideoCard :is-scroll="false" />
    </NTabPane>
    <NTabPane name="history" tab="观看历史">
      <VideoCard :is-scroll="false" />
    </NTabPane>
  </NTabs>
</template>
<script lang="ts" setup>
import { NTabs, NTabPane } from "naive-ui";
import VideoCard from "../VideoCard.vue";
</script>
<style>
.demo-tabs > .el-tabs__content {
  padding: 32px;
  color: #6b778c;
  font-size: 32px;
  font-weight: 600;
}
</style>
