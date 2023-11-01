<!--
 * @Author: Xu Ning
 * @Date: 2023-10-25 16:22:40
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-01 17:02:39
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\components\HeaderMenu.vue
-->
<script lang="ts" setup>
import { watch, h, ref, Component, reactive, onBeforeMount, onMounted } from 'vue'
import { NIcon, NMenu, NInput, NButton, NAvatar, NPopover, NTag, NForm, NTabs, NImage, NTabPane, NModal, NFormItemRow } from 'naive-ui'
import type { MenuOption} from 'naive-ui'
import { Search,Add } from '@vicons/ionicons5'
import { userStore } from "@/stores/user";
import { videoStore } from '@/stores/video'
import { routeStore } from '@/stores/route'
import { login, register } from "@/apis/login";
import router from "@/router";
import { useMessage } from "naive-ui";
import PostVedio from "@/components/video/PostVideo.vue";

// 消息弹窗
const message = useMessage();
// 路由数据
const activeIndex = ref("");
// 登录登出样式
const menuClass = ref<string>('')
// 搜索str
const inputstr = ref<string>('')
const historyList = ref<Array<string>>([])
const historyTabVisible = ref<boolean>(true)

// 登录表单数据
const loginFormVisible = ref<boolean>(false);
const loginForm = reactive({
  phoneNum: "",
  pwd: "",
});

const registerForm = reactive({
  phoneNum: "",
  pwd: "",
  repeatPwd:""
});

// TODO: history search 挂载渲染早，不能这么写
onBeforeMount(() => {
  historyList.value = ['1111','222222']
  console.log(historyList.value, historyList.value[0])
})

onMounted(() => {
  menuClass.value =  userStore().token?'header-menu-login':'header-menu-logout'
})

// 投稿表单数据
const isVideoFormVisible = ref<boolean>(false);

// 渲染图标
function renderIcon (icon: Component) {
  return () => h(NIcon, null, { default: () => h(icon) })
}

const doSearch = () =>{
  let searchValue = inputstr.value
  videoStore().search_value = searchValue
  if(routeStore().name != 'search'){
    router.push({ name: "search"});
  }
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
          src:'/logo-svg.svg',
          width: 100,
          previewDisabled: true
        },
        ''
      ),
    key: 'logo',
  },
  {
    label: () => 
    h(NPopover, 
        {
          trigger:'focus',
          showArrow:false
        },
        {
          trigger: () => {
                return h(NInput, 
                  {
                    modelValue: inputstr.value,
                    onUpdateValue:(value:any)=>{
                      console.log('hhh', value)
                      inputstr.value = value;
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
      ),
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
    label: '登录',
    style:'margin-left:auto',
    key: 'login',
  }
]

// 用户登录
const doLogin = () => {
  login(loginForm.phoneNum, loginForm.pwd).then( async (res: any) => {
    if(res.status_code == 200){
      menuClass.value = 'header-menu-login'
      userStore().token = res.accessToken;
      userStore().user_id = res.user_id;
      userStore().isLoggedIn = true;
      userStore().phoneNum = loginForm.phoneNum;
      userStore().avatar = res.avatar;
      message.success("登录成功");
      loginFormVisible.value = false;
      window.location.reload()
      router.push("/");
    }
    else{
      message.error(res.status_msg)
    }
  });
};

// 用户注册
const doRegister = () => {
  //发请求
  if(registerForm.pwd == registerForm.repeatPwd){
    let defaultAvatar = 'http://s327crbzf.hn-bkt.clouddn.com/a74c50c16c528f561ecb3fc2f8e4856b9bd70c828616c558cad98be75a400c81.jpg'
    let defaultBg = 'http://s327crbzf.hn-bkt.clouddn.com/d689982a1aa68d1242a802fc10228e9a8f17901980bf4251c295ad619592ad98.jpg'
    register('匿名用户MOMO', 0, registerForm.phoneNum, registerForm.pwd,'该用户暂时还没有个人简介哦',defaultAvatar, defaultBg).then((res: any) => {
      if(res.status_code == 200){
        message.success("注册成功");
      }
      else{
        message.error(res.status_msg)
      }
    }); 
  }
  else{
    message.error('密码不一致，请重试')
  }
};

// 用户登出
const doLogout = () => {
  menuClass.value = 'header-menu-logout'
  userStore().isLoggedIn = false;
  userStore().token = "";
  userStore().avatar = "";
  userStore().name = "";
  userStore().gender = -1;
  userStore().signature = "";
  userStore().phoneNum = "";
  userStore().background_image = "";
  userStore().user_id = -1;
  message.success("已退出");
  router.push("/");
  window.location.reload()
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
</script>

<template>
  <div>
    <n-menu :class="menuClass" v-model:value="activeIndex" mode="horizontal" :options="userStore().isLoggedIn?loggedMenuOptions:notLogmenuOptions" />
    <!-- <ElDialog v-model="loginFormVisible" title="登录" width="30%">
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
    </ElDialog> -->
    <NModal
      v-model:show="loginFormVisible"
      class="custom-card"
      preset="card"
      :closable="false"
      size="huge"
      :bordered="false"
    >
    <n-tabs
      class="card-tabs"
      default-value="login"
      size="large"
      animated
      pane-wrapper-style="margin: 0 -4px"
      pane-style="padding-left: 4px; padding-right: 4px; box-sizing: border-box;"
    >
    
      <n-tab-pane name="login" tab="登录">
        <n-form :model="loginForm">
          <n-form-item-row label="手机号">
            <n-input  v-model:value="loginForm.phoneNum"
            placeholder="请输入"
              autocomplete="off"
              clearable/>
          </n-form-item-row>
          <n-form-item-row label="密码">
            <n-input v-model:value="loginForm.pwd"
              autocomplete="off"
              type="password"
              placeholder="请输入"
              show-password
              clearable/>
          </n-form-item-row>
        </n-form>
        <n-button type="primary" @click="doLogin" block secondary strong>
          登录
        </n-button>
      </n-tab-pane>
      <n-tab-pane name="register" tab="注册">
        <n-form :model="registerForm">
          <n-form-item-row label="手机号">
            <n-input v-model:value="registerForm.phoneNum"
              autocomplete="off"
              placeholder="请输入"
              clearable/>
          </n-form-item-row>
          <n-form-item-row label="密码">
            <n-input v-model:value="registerForm.pwd" 
            autocomplete="off"
              type="password"
              placeholder="请输入"
              show-password
              clearable/>
          </n-form-item-row>
          <n-form-item-row label="重复密码">
            <n-input v-model:value="registerForm.repeatPwd" 
            autocomplete="off"
              type="password"
              placeholder="请输入"
              show-password
              clearable/>
          </n-form-item-row>
        </n-form>
        <n-button type="primary" block secondary strong @click="doRegister">
          注册
        </n-button>
      </n-tab-pane>
    </n-tabs>
    </NModal>
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

.custom-card{
  width: 30vw;
}
.header-menu-login{
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

.header-menu-logout{
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
  .n-menu-item:nth-child(2){
    margin-left: auto;
    width: 20vw;
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
