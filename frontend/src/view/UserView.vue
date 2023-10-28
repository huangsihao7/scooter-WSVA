<!--
 * @Author: Xu Ning
 * @Date: 2023-05-08 15:29:52
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-25 23:02:42
 * @Description: 
 * @FilePath: \scooter-wsva\frontend\src\view\User.vue
-->
<script setup lang="ts">
import MyHeaderCom from '@/components/myinfo/MyHeaderCom.vue';
import MyInteractCom from '@/components/myinfo/MyInteractCom.vue';
</script>


<template>
  <div class="user">
    <div class="header">
      <MyHeaderCom/>
    </div>
    <div class="interaction">
      <MyInteractCom/>
    </div>
  </div>
</template>
    
    

<style scoped lang="scss">
 .user{
    padding: 10vh 10vw;
    .interaction{
      margin-top: 10vh;
    }
 }
</style>
