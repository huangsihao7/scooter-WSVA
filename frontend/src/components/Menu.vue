<!--
 * @Author: Xu Ning
 * @Date: 2023-10-25 16:22:40
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-28 12:32:33
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\components\Menu.vue
-->
<script lang="ts" setup>
import { ref, reactive } from 'vue'
import { userStore } from '@/stores/user'
import { login } from '@/apis/login'
import router from '@/router'
import { 
  VideoCameraFilled,
  CirclePlus,
  Search
 } from '@element-plus/icons-vue'
 import { ElMessage } from 'element-plus'
import  PostVedio from '@/components/video/PostVideo.vue'

// 路由数据
const activeIndex = ref('')

// 搜索数据
const searchText = ref<string>('')
const isSearch = ref<boolean>(false)
const isTagClose = ref<boolean>(false)
const searchHistory = ref<Array<string>>(
  ['aaaaaa','bbbbbb','ccccc','dddddddddd']
)

// 登录表单数据
const loginFormVisible = ref<boolean>(false)
const formLabelWidth = '50px'
const form = reactive({
  username: '',
  pwd: ''
})

// 投稿表单数据
const VideoFormVisible = ref<boolean>(false)


// 路由选择index标志
const handleSelect = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
  switch (key) {
    case 'login':
      loginFormVisible.value = true
      break;
    case 'logout':
      doLogout()
      break;
  }
 
}

// 用户登录
const doLogin = () =>{
  //发请求
  login(form.username, form.pwd).then((res: any) => {
    userStore().token = res.accessToken
    userStore().user_id = res.user_id
    userStore().isLoggedIn = true
    userStore().avatar ='http://' + res.avatar
    userStore().username = form.username
    ElMessage({
      message: '登录成功',
      type: 'success'
    })
    router.push('/')
  })
  loginFormVisible.value = false
}

// 用户登出
const doLogout = () =>{
  userStore().isLoggedIn = false
  userStore().token = ''
  userStore().avatar = ''
  userStore().username = ''
  userStore().user_id = -1
  ElMessage({
    message: '已退出',
    type: 'success'
  })
}

// 搜索
const SearchFunc = () => {
  console.log('Enter SearchFunc',searchText.value);
}

const handleClose = (tag: string) => {
  isTagClose.value = true
  searchHistory.value.splice(searchHistory.value.indexOf(tag), 1)
}

// 搜索框是否被点击
const getFocus = () =>{
  isSearch.value = true
}

// 搜索框是否失去焦点
const getBlur = () =>{
  if(!isTagClose.value){
    isSearch.value = false
  }
}

// 获取投稿dialogue
const getPostVideoForm = () =>{
  VideoFormVisible.value = true
}

// 投稿完成后的回调
const updateVisible = (flag : boolean) =>{
  VideoFormVisible.value = flag
  console.log('callback',flag)
}

</script>

<template>
  <el-menu
    :default-active="activeIndex"
    class="el-menu-demo"
    mode="horizontal"
    :ellipsis="false"
    @select="handleSelect"
  >
    <el-menu-item >
      <el-icon><VideoCameraFilled /></el-icon>
      See World
    </el-menu-item>
    <div class="menu-search">
      <el-input
        v-model="searchText"
        placeholder="搜索"
        class="search-input"
        @keyup.enter="SearchFunc"
        @focus = "getFocus"
        @blur = "getBlur"
      >
        <template #append>
          <el-button :icon="Search" class="menu-search-btn" @click="SearchFunc">搜索</el-button>
        </template>
      </el-input>
      <el-card id="search-tab" shadow="always" v-if="isSearch" >
        <div><el-text class="mx-1" size="small">历史记录</el-text></div>
        <el-tag
          class="mx-1"
          v-for="history in searchHistory"
          :key="history"
          effect="light"
          type="info"
          closable
          @close="handleClose(history)"
        >
          {{ history }}
        </el-tag>
      </el-card>
    </div>

    <div class="flex-grow" />
    
    <el-menu-item>
    </el-menu-item>
    <div class="post-btn">
      <el-button @click="getPostVideoForm">
        <el-icon>
          <CirclePlus/>
        </el-icon>
        投稿
      </el-button>
    </div>
      
    <el-sub-menu index="logout" v-if="userStore().isLoggedIn">
      <template #title>
        <el-avatar :src=userStore().avatar />
      </template>
      <el-menu-item index="logout">退出登录</el-menu-item>
    </el-sub-menu>
    <el-menu-item index="login" v-else>
      登录
    </el-menu-item>

    <el-dialog v-model="loginFormVisible" title="登录" width="30%">
      <el-form :model="form">
        <el-form-item label="账号" :label-width="formLabelWidth">
          <el-input v-model="form.username" autocomplete="off" placeholder="输入账号" clearable />
        </el-form-item>
        <el-form-item label="密码" :label-width="formLabelWidth">
          <el-input v-model="form.pwd" autocomplete="off" type="password" placeholder="输入密码" show-password clearable />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button type="primary" @click="doLogin">
            登录
          </el-button>
        </span>
      </template>
    </el-dialog>

    <PostVedio :VideoFormVisible="VideoFormVisible"  @visible-update="updateVisible"/>
    
  </el-menu>
</template>



<style scoped>
.flex-grow {
  flex-grow: 1;
}

.menu-search{
  margin: auto;
  margin-left: 30px;
  margin-top: 14px;
  width:20vw;
}

.el-input-group__append button.el-button:hover {
  /* background-color: #ecf5ff !important; */
  color: #409EFF!important;
  border: 1px solid #e1e4ea;
}

#search-tab {
  text-align: left;
  z-index: 3000;
}

.el-text{
  font-weight: bold;
  padding-bottom: 5px;
}

.mx-1{
  margin: 3px;
  max-width: 99%;
}

.post-btn{
  margin: auto;
  margin-right: 5px;
}


</style>
