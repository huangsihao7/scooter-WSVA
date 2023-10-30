<!--
 * @Author: Xu Ning
 * @Date: 2023-05-08 15:29:52
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-30 20:36:04
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\view\UserView.vue
-->
<script setup lang="ts">
import MyHeaderCom from "@/components/myinfo/MyHeaderCom.vue";
import MyInteractCom from "@/components/myinfo/MyInteractCom.vue";
import { useRoute } from "vue-router";
import { computed, onMounted, ref } from "vue";
import { userStore } from "@/stores/user";

const route = useRoute();
const userId = computed(() => route.params.id);
const passUserId = ref<any>();
onMounted(() => {
  console.log("fffffffff", userId, userId.value);
  if (userId.value) {
    passUserId.value = userId.value;
  } else {
    passUserId.value = userStore().user_id;
  }
});
</script>

<template>
  <div class="user">
    <div class="header">
      <MyHeaderCom v-if="passUserId" :user-id="passUserId" />
    </div>
    <div class="interaction">
      <MyInteractCom />
    </div>
  </div>
</template>

<style scoped lang="scss">
.user {
  padding: 5vh 10vw;

  .interaction {
    margin-top: 10px;
  }
}
</style>
