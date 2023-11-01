<!--
 * @Author: Xu Ning
 * @Date: 2023-10-25 16:22:40
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-01 14:09:58
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\components\HeaderMenu.vue
-->
<script lang="ts" setup>
import { watch, h, ref, Component, reactive, onBeforeMount, onMounted } from 'vue'
import { NIcon, NMenu, NInput, NButton, NAvatar, NPopover, NTag, NImage } from 'naive-ui'
import type { MenuOption} from 'naive-ui'
import { Search,Add } from '@vicons/ionicons5'
import { userStore } from "@/stores/user";
import { videoStore } from '@/stores/video'
import { routeStore } from '@/stores/route'
import { login } from "@/apis/login";
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
  //发请求
  login(form.phoneNum, form.pwd).then((res: any) => {
    menuClass.value = 'header-menu-login'
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
  menuClass.value = 'header-menu-logout'
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

</script>

<template>
  <div>
    <n-menu :class="menuClass" v-model:value="activeIndex" mode="horizontal" :options="userStore().isLoggedIn?loggedMenuOptions:notLogmenuOptions" />
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
    <!-- <NDialogProvider v-model="loginFormVisible" title="登录" width="30%">
      <NForm :model="form">
        <NFormItem label="账号" :label-width="formLabelWidth">
          <NInput 
              v-model="form.phoneNum"
              autocomplete="off"
              placeholder="输入账号"
              clearable></NInput>
        </NFormItem>
        <NFormItem label="密码" :label-width="formLabelWidth">
          <NInput 
              v-model="form.pwd"
              autocomplete="off"
              type="password"
              placeholder="输入密码"
              show-password
              clearable></NInput>
        </NFormItem>
        <div style="display: flex; justify-content: flex-end">
          <n-button round type="primary" @click="doLogin">
            登录
          </n-button>
        </div>
      </NForm>
    </NDialogProvider> -->
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
