<!--
 * @Author: Xu Ning
 * @Date: 2023-10-28 12:30:12
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-28 14:56:12
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\components\myinfo\myHeaderCom.vue
-->
<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { userStore } from '@/stores/user'
import  {getUserInfo}  from '@/apis/user'
import { UserType } from '@/apis/interface'
import { NButton, NIcon } from 'naive-ui';
import { CashOutline as CashIcon } from '@vicons/ionicons5'
const avatar = userStore().avatar
const user_id = userStore().user_id
const userInfo = ref<UserType>()

const getUserInfoFunc = () =>{
    getUserInfo(user_id).then((res:any)=>{
        userInfo.value = res
    })
}

onMounted(() => {
    getUserInfoFunc()
})

</script>

<template>
    <div class="header">
        <el-row>
            <el-col :span="4">
                <el-avatar 
                    :src="avatar"
                    />
            </el-col>
            <el-col :span="20" v-if="userInfo" class="info-tab">
                <el-text tag="b" >{{ userInfo.name }}</el-text>
                <el-text tag="p">{{ userInfo.signature }}</el-text>
                <div class="follow">
                    <n-button color="#606266" text>关注 {{ userInfo.follow_count }}</n-button>
                    <el-divider direction="vertical" />
                    <n-button color="#606266" text>粉丝 {{ userInfo.follower_count }}</n-button>
                </div>
                <n-button strong round class="edit-info" color="#409eff85">
                    <template #icon>
                        <n-icon><cash-icon /></n-icon>
                    </template>
                    编辑资料
                </n-button>
            </el-col>
        </el-row>
    </div>
</template>
  

  <style scoped lang="scss">
  .header {
    background-color: aquamarine;
    text-align: left;
    display: flex;
    padding: 3vh 0;
    border-radius: 25px;
    background: no-repeat center top / 100% 100%;
	background-image: linear-gradient(to right, rgba(255, 255, 255, 1), rgba(255, 255, 255, 0)), url(http://127.0.0.1:8080/17.jpg);
    .el-avatar{
        float: right;
        font-size: 5rem;
        width: calc((80vw - 260px) / 6);
        // height: calc((80vw - 260px) / 6);
        height: calc((90vh - 70px) / 5);
    }
    .info-tab{
        // flex-direction: column;
        // justify-content: space-evenly;
        .edit-info{
            position: absolute;
            top: 0;
            right: 20px;
        }
        .el-text, .follow{
            margin: 10px 20px;
            display: block;
        }
        .follow{
            .n-button{
                padding-right: 10px;
            }
        }
        .el-text:nth-child(1){
            color: black;
            font-size: 2.5rem;
            font-weight: bold;
        }
    }
    
  }
  .header > div {
    flex: 1;
  }
  
  .header > div:not(:last-child) {
    border-right: 1px solid var(--el-border-color);
  }
  </style>
  