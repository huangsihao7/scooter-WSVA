<!--
 * @Author: Xu Ning
 * @Date: 2023-05-08 15:29:52
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-04 17:31:04
 * @Description: 关注页面
 * @FilePath: \scooter-WSVA\frontend\src\view\FollowView.vue
-->
<script setup lang="ts">
import UserCard from "@/components/cards/UserCard.vue";
import { computed } from "vue";
import { useRoute } from "vue-router";
import { userStore } from "@/stores/user";
const route = useRoute();
const userId = computed(() =>
  route.params.id ? parseInt(route.params.id.toString()) : userStore().user_id,
);
</script>

<template>
  <UserCard :user-id="userId" />
</template>

<style scoped></style>
