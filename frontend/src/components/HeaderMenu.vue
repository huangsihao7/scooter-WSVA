<!--
 * @Author: Xu Ning
 * @Date: 2023-10-25 16:22:40
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-31 19:53:23
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\components\HeaderMenu.vue
-->
<script lang="ts" setup>
import { watch, h, ref, Component, reactive, onBeforeMount } from 'vue'
import { NIcon, NMenu, NInput, NButton, NAvatar, NPopover, NTag, NImage } from 'naive-ui'
import type { MenuOption } from 'naive-ui'
import { Search,Add } from '@vicons/ionicons5'
import { userStore } from "@/stores/user";
import { login } from "@/apis/login";
import router from "@/router";
import { useMessage } from "naive-ui";
import PostVedio from "@/components/video/PostVideo.vue";


// 消息弹窗
const message = useMessage();
// 路由数据
const activeIndex = ref("");

// 搜索str
const inputstr = ref<string>('')
const historyList = ref<Array<string>>([])
const historyTabVisible = ref<boolean>(true)

// 登录表单数据
const loginFormVisible = ref<boolean>(false);
const formLabelWidth = "50px";
const form = reactive({
  phoneNum: "",
  pwd: "",
});

// TODO: history search 挂载渲染早，不能这么写
onBeforeMount(() => {
  historyList.value = ['1111','222222']
  console.log(historyList.value, historyList.value[0])
})

// 投稿表单数据
const isVideoFormVisible = ref<boolean>(false);

// 渲染图标
function renderIcon (icon: Component) {
  return () => h(NIcon, null, { default: () => h(icon) })
}

const doSearch = () =>{
  console.log('search')
}

// 渲染搜索
function renderSearch (searchStr:string, historyList:any) {
  console.log(historyList)
  return () => 
    h(NPopover, 
    {
      trigger:'focus',
      showArrow:false
    },
    {
      trigger: () => {
            return h(NInput, 
              {
                props:{
                  value:searchStr,
                },
                class:"h-input",
                round:true,
                placeholder:'搜索',
                onKeydown:(e)=>{ if(e.key == 'Enter'){doSearch()}},
                onFocus:()=>{historyTabVisible.value = true},
                onBlur:()=>{historyTabVisible.value = false}
              },
                {
                  suffix:renderIcon(Search)
                }
            )
      },
      default: renderTags('亚运会冠军')
    }
  )
}

//投稿渲染
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

// 头像渲染
function renderAvatar (avatarSrc:string) {
  return () => 
  h(NAvatar, 
  {
    src:avatarSrc,
    round:true,
  },'')
}

// 渲染history
function renderTags (historyStr:string) {
  return () => 
  h( NTag,{
   type:'info',
   round:true,
   closable:true
  },historyStr)
}

// 菜单数据
const loggedMenuOptions: MenuOption[] = [
{
    label: () =>
      h(
        NImage,
        {
          src:'public/logo-svg.svg',
          width: 100,
          previewDisabled: true
        },
        ''
      ),
    key: 'logo',
  },
  {
    // label: renderSearch(inputstr.value),renderSearch2
    label: renderSearch(inputstr.value, historyList),
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
        NImage,
        {
          src:'public/logo-svg.svg',
          width: 100,
          previewDisabled: true
        },
        ''
      ),
    key: 'logo'
  },
  {
    label: renderSearch(inputstr.value, historyList.value),
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

// 获取投稿dialogue
const getPostVideoForm = () => {
  isVideoFormVisible.value = true;
};

// 投稿完成后的回调
const updateVisible = (flag: boolean) => {
  isVideoFormVisible.value = flag;
  console.log("callback", flag);
};

// 路由选择index标志
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

//监听搜索回车

</script>

<template>
  <div>
    <n-menu class="header-menu" v-model:value="activeIndex" mode="horizontal" :options="userStore().isLoggedIn?loggedMenuOptions:notLogmenuOptions" />
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
  </div>
</template>

<style lang="scss">
.history-card{
  width: 50vw;
}
.header-menu{
  width: 100%;
  height: 60px;
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
