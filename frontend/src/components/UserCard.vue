<!--
 * @Author: Xu Ning
 * @Date: 2023-10-27 14:13:32
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-28 12:57:44
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\components\UserCard.vue
-->
<script lang="ts" setup>
import { reactive } from 'vue'

const res:any = reactive([
    {
        avatar:'http://127.0.0.1:8080/OIP-C.jpg',
        url:'http://127.0.0.1:8080/1.jpg',
        name:'我今天好帅',
        content:'是一个热爱学习的人',
        isfollowed: true
    },
    {
        avatar:'http://127.0.0.1:8080/OIP-C.jpg',
        url:'http://127.0.0.1:8080/2.jpg',
        name:'吃大餐饱了',
        content:'明日复明日，ooooo',
        isfollowed: true
    },
    {
        avatar:'http://127.0.0.1:8080/OIP-C.jpg',
        url:'http://127.0.0.1:8080/3.jpg',
        name:'吃大餐饱了',
        content:'明日复明日，ooooo',
        isfollowed: true
    },
    {
        avatar:'http://127.0.0.1:8080/OIP-C.jpg',
        url:'http://127.0.0.1:8080/4.jpg',
        name:'原皮',
        content:'明日复明日，ooooo',
        isfollowed: true
    },
    {
        avatar:'http://127.0.0.1:8080/OIP-C.jpg',
        url:'http://127.0.0.1:8080/5.jpg',
        name:'大师傅',
        content:'学习是终身事业！！！！！！！！！',
        isfollowed: true
    },
    {
        avatar:'http://127.0.0.1:8080/OIP-C.jpg',
        url:'http://127.0.0.1:8080/6.jpg',
        name:'大发噶是否',
        content:'学习是终身事业！！！！！！！！！',
        isfollowed: true
    },
    {
        avatar:'http://127.0.0.1:8080/OIP-C.jpg',
        url:'http://127.0.0.1:8080/7.jpg',
        name:'爱上的人服务',
        content:'学习是终身事业！！！！！！！！！',
        isfollowed: true
    },
    {
        avatar:'http://127.0.0.1:8080/OIP-C.jpg',
        url:'http://127.0.0.1:8080/8.jpg',
        name:'暴怒改变',
        content:'学习是终身事业！！！！！！！！！',
        isfollowed: true
    },
])

const handleShowVedio = () =>{
    console.log('show')
}

const cancleFollow = (item:any, index:any) =>{
    if(item.isfollowed){
        res[index].isfollowed = false
    }
    else{
        res[index].isfollowed = true
    }
    
    // item.isfollowed = false
    // TODO 取消关注接口
}

</script>

<template>
  <el-space wrap>
    <el-card v-for="(info, index) in res" :key="index" class="box-card" style="width: calc((100vw - 260px) / 4)">
      <template #header>
        <div class="card-header">
            <div class="avatar"><el-avatar :size="60" :src="info.avatar"/></div>
            <div>
                <span>{{ info.name }}</span>
                <el-button @click="cancleFollow(info, index)">{{ info.isfollowed == true?'已关注':'关注' }}</el-button>
                <p>{{ info.content }}</p>
            </div>
        </div>
        
      </template>
      <div>
        <el-image @click="handleShowVedio" :src="info.url" fit="fill"/>
      </div>
    </el-card>
  </el-space>
 </template>
  

  
<style lang="scss" scoped>

.el-card{
    --el-card-padding:0
}
.box-card{
    margin-top: 15px;
    margin-left: 10px;
    .el-card__body{
        
        .el-image{
            width: 100%;
            height: 30vh;
        }
        .play-btn{
            width: 100%;
            text-align: center;
            position: relative;
            bottom: 125px;
            top: 35.5%;
            z-index: 2;
            font-size: 50px;
        }
    }
    
    .card-header{
        margin-top: 15px;
        display: flex;
        justify-content: left;
        .avatar, span, p, .el-button{
            margin-left: 10px;
        }

        span{
            font-weight: bold;
            font-size: 1.2rem;
        }
        p{
            font-size:0.8rem;
        }
    }
}
</style>
