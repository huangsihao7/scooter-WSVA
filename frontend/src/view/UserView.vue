<!--
 * @Author: Xu Ning
 * @Date: 2023-05-08 15:29:52
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-28 14:28:55
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\view\UserView.vue
-->
<script setup lang="ts">
import MyHeaderCom from '@/components/myinfo/MyHeaderCom.vue';
import MyInteractCom from '@/components/myinfo/MyInteractCom.vue';
</script>


<template>
  <el-scrollbar>
    <div class="user">
      <div class="header">
        <MyHeaderCom/>
      </div>
      <div class="interaction">
        <MyInteractCom/>
      </div>
    </div>
  </el-scrollbar>
</template>
    
    

<style scoped lang="scss">
 .user{
    padding: 5vh 10vw;
    .interaction{
      margin-top: 10px;
    }
 }
</style>
