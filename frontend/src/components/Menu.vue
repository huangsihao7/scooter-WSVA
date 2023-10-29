<!--
 * @Author: Xu Ning
 * @Date: 2023-10-25 16:22:40
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-28 13:41:17
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\components\Menu.vue
-->
<script lang="ts" setup>
import { ref, reactive } from "vue";
import { userStore } from "@/stores/user";
import { login } from "@/apis/login";
import router from "@/router";
import { VideoCameraFilled, CirclePlus, Search } from "@element-plus/icons-vue";
import { ElMessage } from "element-plus";
import PostVedio from "@/components/video/PostVideo.vue";

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
  username: "",
  pwd: "",
});

// 投稿表单数据
const VideoFormVisible = ref<boolean>(false);

// 路由选择index标志
const handleSelect = (key: string, keyPath: string[]) => {
  console.log(key, keyPath);
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
  login(form.username, form.pwd).then((res: any) => {
    userStore().token = res.accessToken;
    userStore().user_id = res.user_id;
    userStore().isLoggedIn = true;
    userStore().avatar = res.avatar;
    ElMessage({
      message: "登录成功",
      type: "success",
    });
    router.push("/");
  });
  loginFormVisible.value = false;
};

// 用户登出
const doLogout = () => {
  userStore().isLoggedIn = false;
  userStore().token = "";
  userStore().avatar = "";
  userStore().username = "";
  userStore().user_id = -1;
  ElMessage({
    message: "已退出",
    type: "success",
  });
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
  VideoFormVisible.value = true;
};

// 投稿完成后的回调
const updateVisible = (flag: boolean) => {
  VideoFormVisible.value = flag;
  console.log("callback", flag);
};
</script>

<template>
  <ElAffix :offset="0">
    <ElMenu
      :default-active="activeIndex"
      class="el-menu-demo"
      mode="horizontal"
      :ellipsis="false"
      @select="handleSelect"
    >
      <ElMenuItem>
        <ElIcon><VideoCameraFilled /></ElIcon>
        See World
      </ElMenuItem>
      <div class="menu-search">
        <ElInput
          v-model="searchText"
          placeholder="搜索"
          class="search-input"
          @keyup.enter="SearchFunc"
          @focus="getFocus"
          @blur="getBlur"
        >
          <template #append>
            <ElButton :icon="Search" class="menu-search-btn" @click="SearchFunc"
            >搜索</ElButton
            >
          </template>
        </ElInput>
        <ElCard v-if="isSearch" id="search-tab" shadow="always">
          <div><ElText class="mx-1" size="small">历史记录</ElText></div>
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
      </div>

      <div class="flex-grow" />

      <ElMenuItem> </ElMenuItem>
      <div class="post-btn">
        <ElButton @click="getPostVideoForm">
          <ElIcon>
            <CirclePlus />
          </ElIcon>
          投稿
        </ElButton>
      </div>

      <ElSubMenu v-if="userStore().isLoggedIn" index="logout">
        <template #title>
          <ElAvatar :src="userStore().avatar" />
        </template>
        <ElMenuItem index="logout">退出登录</ElMenuItem>
      </ElSubMenu>
      <ElMenuItem v-else index="login"> 登录 </ElMenuItem>

      <ElDialog v-model="loginFormVisible" title="登录" width="30%">
        <ElForm :model="form">
          <ElFormItem label="账号" :label-width="formLabelWidth">
            <ElInput
              v-model="form.username"
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
        :VideoFormVisible="VideoFormVisible"
        @visible-update="updateVisible"
      />
    </ElMenu>
  </ElAffix>
</template>

<style scoped>
.flex-grow {
  flex-grow: 1;
}

.menu-search {
  margin: auto;
  margin-left: 30px;
  margin-top: 14px;
  width: 20vw;
}

.el-input-group__append button.el-button:hover {
  /* background-color: #ecf5ff !important; */
  color: #409eff !important;
  border: 1px solid #e1e4ea;
}

#search-tab {
  text-align: left;
  z-index: 3000;
}

.el-text {
  font-weight: bold;
  padding-bottom: 5px;
}

.mx-1 {
  margin: 3px;
  max-width: 99%;
}

.post-btn {
  margin: auto;
  margin-right: 5px;
}
</style>
