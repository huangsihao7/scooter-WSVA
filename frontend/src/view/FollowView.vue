<!--
 * @Author: Xu Ning
 * @Date: 2023-05-08 15:29:52
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-02 16:48:32
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\view\FollowView.vue
-->
<script setup lang="ts">
import UserCard from "@/components/UserCard.vue";
import { computed } from "vue";
import { useRoute } from "vue-router";
import { routeStore } from "@/stores/route";
import { userStore } from "@/stores/user";
const route = useRoute();
const userId = computed(() => parseInt(route.params.id.toString()));
const cardType = computed(()=> routeStore().name)

</script>

<template>
  <UserCard v-if="cardType!='follow'" :cardType="cardType" :userId="userId"/>
  <UserCard v-else cardType='follow' :userId="userStore().user_id"/>
</template>

<style scoped></style>
