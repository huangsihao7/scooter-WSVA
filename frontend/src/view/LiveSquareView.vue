<!--
 * @Author: Xu Ning
 * @Date: 2023-11-05 18:42:02
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-05 22:34:08
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\view\LiveSquareView.vue
-->
<script setup lang="ts">
import { ref, onMounted } from  'vue'
import LiveCard from "@/components/cards/LiveCard.vue";
const videos = ref<any>([])

onMounted(() => {
  //TODO getliveinfo
})
</script>
<template>
  <LiveCard :videos="videos"/>
</template>

<style scoped></style>
