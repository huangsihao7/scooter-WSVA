<!--
 * @Author: Xu Ning
 * @Date: 2023-05-08 15:29:52
 * @LastEditors: huangsihao7 1057434651@qq.com
 * @LastEditTime: 2023-10-30 13:26:31
 * @Description: 
 * @FilePath: /scooter-WSVA/frontend/src/view/UserView.vue
-->
<script setup lang="ts">
import MyHeaderCom from "@/components/myinfo/MyHeaderCom.vue";
import MyInteractCom from "@/components/myinfo/MyInteractCom.vue";
</script>

<template>
  <div class="user">
    <div class="header">
      <MyHeaderCom />
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
