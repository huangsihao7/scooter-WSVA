<!--
 * @Author: Xu Ning
 * @Date: 2023-10-25 16:22:40
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-25 23:09:36
 * @Description: 
 * @FilePath: \scooter-wsva\frontend\src\components\Menu.vue
-->
<script lang="ts" setup>
import { ref} from 'vue'
import { 
  VideoCameraFilled,
  CirclePlus,
  Search
 } from '@element-plus/icons-vue'
// import { userStore } from '@/stores/user';

const activeIndex = ref('')
const searchText = ref<string>('')
const isSearch = ref<boolean>(false)
const searchHistory = ref<Array<string>>(
  ['aaaaaa','bbbbbb','ccccc','dddddddddddddddddddddddddddd']
)

const handleSelect = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}

const SearchFunc = () => {
  console.log('Enter SearchFunc',searchText.value);
}

const handleClose = (tag: string) => {
  searchHistory.value.splice(searchHistory.value.indexOf(tag), 1)
}

const getFocus = () =>{
  isSearch.value = true
}

const getBlur = () =>{
  isSearch.value = false
}

var avatar : string = 'https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png'

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
      <el-card shadow="always" v-if="isSearch" >
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
    <el-menu-item index="1">
      <el-icon>
        <CirclePlus/>
      </el-icon>
      投稿
    </el-menu-item>
    <el-sub-menu index="2" v-if="avatar">
      <template #title>
        <el-avatar
          :src="avatar"
        />
      </template>
      <el-menu-item index="logout">退出登录</el-menu-item>
    </el-sub-menu>
    <el-menu-item index="login" v-else>
      登录
    </el-menu-item>
  </el-menu>
</template>



<style>
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

.el-card {
  text-align: left;
}

.el-text{
  font-weight: bold;
  padding-bottom: 5px;
}

.mx-1{
  margin: 3px;
  max-width: 99%;
}



</style>
