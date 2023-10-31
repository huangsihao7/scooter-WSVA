<!--
 * @Author: Xu Ning
 * @Date: 2023-10-25 16:22:40
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-31 16:05:38
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\components\HeaderMenu.vue
-->
<script lang="ts" setup>
import { watch, h, ref, Component, reactive } from 'vue'
import { NIcon, NMenu, NInput, NButton, NAvatar } from 'naive-ui'
import type { MenuOption } from 'naive-ui'
import {
  Book as BookIcon,
  Search,
  Add,
  Bicycle,
  PersonOutline as PersonIcon,
  WineOutline as WineIcon
} from '@vicons/ionicons5'
import { userStore } from "@/stores/user";
import { login } from "@/apis/login";
import router from "@/router";
import { VideoCameraFilled, CirclePlus } from "@element-plus/icons-vue";
import { useMessage } from "naive-ui";
import PostVedio from "@/components/video/PostVideo.vue";
import { number } from 'echarts';

// 消息弹窗
const message = useMessage();
// 路由数据
const activeIndex = ref("");

// 搜索数据
const searchText = ref<string>("");
const isSearch = ref<boolean>(false);
const isTagClose = ref<boolean>(false);
const searchHistory = ref<Array<string>>([
  "aaaaaa",
  "bbbbbb",
  "ccccc",
  "dddddddddd",
]);

// 登录表单数据
const loginFormVisible = ref<boolean>(false);
const formLabelWidth = "50px";
const form = reactive({
  phoneNum: "",
  pwd: "",
});

// 投稿表单数据
const isVideoFormVisible = ref<boolean>(false);

function renderIcon (icon: Component) {
  return () => h(NIcon, null, { default: () => h(icon) })
}


function renderSearch (num:string) {
  return () => 
    h(NInput, 
    {
      props:{
        value:num,
      },
      class:"h-input",
      round:true,
      placeholder:'搜索',
      onInput:(value)=>{console.log('input', value)},
      onFocus:()=>{historyTabVisible.value = true},
      onBlur:()=>{historyTabVisible.value = false}
    },
      {
        suffix:renderIcon(Search)
      }
    )
}

function renderBtn () {
  return () => 
  h(NButton, 
  {
    attr:{
      id:'post-btn'
    },
    round:true,
    renderIcon:renderIcon(Add),
    onClick:()=>{getPostVideoForm()}
  },'投稿')
}

function renderAvatar (avatarSrc:string) {
  return () => 
  h(NAvatar, 
  {
    src:avatarSrc,
    round:true,
  },'')
}


const inputstr = ref<string>('')
const historyTabVisible = ref<boolean>(false)

// 菜单数据
const loggedMenuOptions: MenuOption[] = [
  {
    label: () =>
      h(
        'div',
        {
          href: 'https://baike.baidu.com/item/%E4%B8%94%E5%90%AC%E9%A3%8E%E5%90%9F',
          target: '_blank',
          rel: 'noopenner noreferrer'
        },
        'Scooter'
      ),
    key: 'logo',
    icon: renderIcon(Bicycle)
  },
  {
    label: renderSearch(inputstr.value),
    key: 'search'
  },
  {
    label: renderBtn(),
    key: 'post-btn',
  },
  {
    label: renderAvatar(userStore().avatar),
    key: 'avatar',
    children: [
      {
        label: '退出登录',
        key: 'logout',
      }
    ]
  }
]

const notLogmenuOptions: MenuOption[] = [
  {
    label: () =>
      h(
        'div',
        {
          href: 'https://baike.baidu.com/item/%E4%B8%94%E5%90%AC%E9%A3%8E%E5%90%9F',
          target: '_blank',
          rel: 'noopenner noreferrer'
        },
        'Scooter'
      ),
    key: 'logo',
    icon: renderIcon(Bicycle)
  },
  {
    label: renderSearch(inputstr.value),
    key: 'search'
  },
  {
    label: renderBtn(),
    key: 'post-btn',
  },
  {
    label: '登录',
    key: 'login',
  }
]

// 路由选择index标志
const handleSelect = (key: string, _keyPath: string[]) => {
  // console.log(key, keyPath);
  switch (key) {
    case "login":
      loginFormVisible.value = true;
      break;
    case "logout":
      doLogout();
      break;
  }
};

// 用户登录
const doLogin = () => {
  //发请求
  login(form.phoneNum, form.pwd).then((res: any) => {
    userStore().token = res.accessToken;
    userStore().user_id = res.user_id;
    userStore().isLoggedIn = true;
    userStore().phoneNum = form.phoneNum;
    userStore().avatar = res.avatar;
    message.success("登录成功");
    router.push("/");
  });
  loginFormVisible.value = false;
};

// 用户登出
const doLogout = () => {
  userStore().isLoggedIn = false;
  userStore().token = "";
  userStore().avatar = "";
  userStore().phoneNum = "";
  userStore().user_id = -1;
  message.success("已退出");
};

// 搜索
const SearchFunc = () => {
  console.log("Enter SearchFunc", searchText.value);
};

const handleClose = (tag: string) => {
  isTagClose.value = true;
  searchHistory.value.splice(searchHistory.value.indexOf(tag), 1);
};

// 搜索框是否被点击
const getFocus = () => {
  isSearch.value = true;
};

// 搜索框是否失去焦点
const getBlur = () => {
  if (!isTagClose.value) {
    isSearch.value = false;
  }
};

// 获取投稿dialogue
const getPostVideoForm = () => {
  isVideoFormVisible.value = true;
};

// 投稿完成后的回调
const updateVisible = (flag: boolean) => {
  isVideoFormVisible.value = flag;
  console.log("callback", flag);
};

watch(()=>activeIndex.value,
(newValue:any)=>{
  switch (newValue) {
    case "login":
      loginFormVisible.value = true;
      break;
    case "logout":
      doLogout();
      break;
  }
})
watch(()=>loginFormVisible.value,
(newValue:any)=>{
  if(!newValue){
    activeIndex.value = ''
  }
  
})


</script>

<template>
  <n-menu class="header-menu" v-model:value="activeIndex" mode="horizontal" :options="userStore().isLoggedIn?loggedMenuOptions:notLogmenuOptions" />
  <ElCard v-if="isSearch" id="search-tab" shadow="always">
        <div>
          <ElText class="mx-1" size="small">历史记录</ElText>
        </div>
        <ElTag
          v-for="history in searchHistory"
          :key="history"
          class="mx-1"
          effect="light"
          type="info"
          closable
          @close="handleClose(history)"
        >
          {{ history }}
        </ElTag>
  </ElCard>
  <ElDialog v-model="loginFormVisible" title="登录" width="30%">
      <ElForm :model="form">
        <ElFormItem label="账号" :label-width="formLabelWidth">
          <ElInput
            v-model="form.phoneNum"
            autocomplete="off"
            placeholder="输入账号"
            clearable
          />
        </ElFormItem>
        <ElFormItem label="密码" :label-width="formLabelWidth">
          <ElInput
            v-model="form.pwd"
            autocomplete="off"
            type="password"
            placeholder="输入密码"
            show-password
            clearable
          />
        </ElFormItem>
      </ElForm>
      <template #footer>
        <span class="dialog-footer">
          <ElButton type="primary" @click="doLogin"> 登录 </ElButton>
        </span>
      </template>
  </ElDialog>
  <PostVedio
    :video-form-visible="isVideoFormVisible"
    @visible-update="updateVisible"
  />
</template>

<style lang="scss">

.header-menu{
  width: 100%;
  height: 60px;
  // flex-direction: row;
  // justify-content: space-between;
  .n-menu-item{
    height: 100%;
    .n-menu-item-content{
      height: 100%;
    }
    
  }
  .n-menu-item:nth-child(1) .n-menu-item-content{
    width: 160px;
  }
  .n-menu-item:nth-child(2) .n-menu-item-content{
    width: 20vw;
  }
  .n-menu-item:nth-child(3){
    margin-left: auto;
  }
  .n-menu-item:nth-child(3) .n-menu-item-content{
    
  }
  .n-menu-item:nth-child(4) .n-menu-item-content{
    
  }
}

.flex-grow {
  flex-grow: 1;
}

#post-btn {
  margin-left: 20vw;
}

.aaa{
  background-color: #409eff;
}

.post-btn {
  margin: auto;
  margin-right: 5px;
}
</style>
