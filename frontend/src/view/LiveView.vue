<!--
 * @Author: Xu Ning
 * @Date: 2023-11-05 18:42:02
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-05 21:14:21
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\view\LiveView.vue
-->
<script setup lang="ts">
import { ref, onMounted, computed } from  'vue'
import LiveCarouselCom from "@/components/live/LiveCarouselCom.vue";
import { useRoute } from 'vue-router';
const route = useRoute()
const userId = computed(() => route.params.id);
const url = computed(() => route.params.url);

onMounted(() => {
  //TODO getliveinfo
})
</script>
<template>
  <LiveCarouselCom :liveId = "userId" :url="url"/>
</template>

<style scoped></style>
