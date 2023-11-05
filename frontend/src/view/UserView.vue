<!--
 * @Author: Xu Ning
 * @Date: 2023-05-08 15:29:52
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-03 22:46:17
 * @Description: 我的页面
 * @FilePath: \scooter-WSVA\frontend\src\view\UserView.vue
-->
<script setup lang="ts">
import MyHeaderCom from "@/components/myinfo/MyHeaderCom.vue";
import MyInteractCom from "@/components/myinfo/MyInteractCom.vue";
import { useRoute } from "vue-router";
import { computed, onMounted, ref } from "vue";
import { userStore } from "@/stores/user";
import { NScrollbar } from "naive-ui";

const route = useRoute();
const userId = computed(() => route.params.id);
const passUserId = ref<any>();

onMounted(() => {
  if (userId.value) {
    passUserId.value = userId.value;
  } else {
    passUserId.value = userStore().user_id;
  }
});
</script>

<template>
  <NScrollbar style="max-height: calc(100vh- 60px)">
    <div class="user">
      <div class="header">
        <MyHeaderCom v-if="passUserId" :user-id="passUserId" />
      </div>
      <div class="interaction">
        <MyInteractCom v-if="passUserId" :user-id="passUserId" />
      </div>
    </div>
  </NScrollbar>
</template>

<style scoped lang="scss">
.user {

  .interaction {
    margin-top: 10px;
  }
}
</style>
