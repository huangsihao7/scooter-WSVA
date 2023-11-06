<!--
 * @Author: Xu Ning
 * @Date: 2023-10-25 16:22:40
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-06 11:29:39
 * @Description: 顶部导航栏组件
 * @FilePath: \scooter-WSVA\frontend\src\components\menu\HeaderMenu.vue
-->
<script lang="ts" setup>
import { watch, h, ref, Component, reactive, onMounted, VNodeChild } from "vue";
import {
  DropdownOption,
  NIcon,
  NMenu,
  NInput,
  NButton,
  NAvatar,
  NForm,
  NTabs,
  NImage,
  NTabPane,
  NModal,
  NFormItemRow,
  NDropdown,
  useMessage,
} from "naive-ui";
import type { MenuOption } from "naive-ui";
import { Search, Add, Trash } from "@vicons/ionicons5";
import { userStore } from "@/stores/user";
import { videoStore } from "@/stores/video";
import { routeStore } from "@/stores/route";
import { login, register } from "@/apis/login";
import router from "@/router";
import PostVedioCom from "@/components/video/PostVideoCom.vue";
import { historyStore } from "@/stores/historySearch";

// 消息弹窗message
const message = useMessage();
// 路由数据
const activeIndex = ref("");
// 登录登出样式
const menuClass = ref<string>("");
// 搜索内容
const searchContent = ref<string>("");
// 历史列表是否可见
const historyTabVisible = ref<boolean>(true);
// 搜索的历史记录数据
const dropOptions = ref<DropdownOption[]>();
// 投稿表单数据
const isVideoFormVisible = ref<boolean>(false);
// 直播表单数据
const isLiveFormVisible = ref<boolean>(false);
// 登录表单数据
const loginFormVisible = ref<boolean>(false);

const loginForm = reactive({
  phoneNum: "",
  pwd: "",
});
// 注册表单数据
const registerForm = reactive({
  phoneNum: "",
  pwd: "",
  repeatPwd: "",
});

onMounted(() => {
  renderHistory();
  menuClass.value = userStore().token
    ? "header-menu-login"
    : "header-menu-logout";
});

// 渲染图标
function renderIcon(icon: Component) {
  return () => h(NIcon, null, { default: () => h(icon) });
}

const searchHistoryNum = ref<number>(0);

// 搜索时把搜索记录保存
const doSearch = (isHistory: boolean, historyValue?: string) => {
  let searchValue = "";
  if (isHistory && historyValue) {
    searchValue = historyValue;
    searchContent.value = historyValue;
    console.log(searchContent.value);

    videoStore().search_value = searchValue;
  } else {
    searchValue = searchContent.value;
    videoStore().search_value = searchValue;
  }
  let child = {
    label: searchValue,
    key: searchHistoryNum.value.toString(),
  };
  searchHistoryNum.value++;
  historyStore().historyData.push(child);
  if (historyStore().historyData.length > 5) {
    historyStore().historyData.shift(); // 移除第一条数据
  }
  if (routeStore().name != "search") {
    router.push({ name: "search" });
  }
  renderHistory();
};

//投稿渲染
function renderPostBtn() {
  return () =>
    h(
      NButton,
      {
        attr: {
          id: "post-btn",
        },
        round: true,
        renderIcon: renderIcon(Add),
        onClick: () => {
          getPostVideoForm();
        },
      },
      "投稿",
    );
}

//直播渲染
function renderLiveBtn() {
  return () =>
    h(
      NButton,
      {
        attr: {
          id: "post-btn",
        },
        round: true,
        renderIcon: renderIcon(Add),
        onClick: () => {
          getLiveForm();
        },
      },
      "我要直播",
    );
}

// 头像渲染
function renderAvatar(avatarSrc: string) {
  return () =>
    h(
      NAvatar,
      {
        src: avatarSrc,
        round: true,
      },
      "",
    );
}

// 渲染历史记录数据
const renderHistory = () => {
  let hisChild = historyStore().historyData;
  let deleteChild = {
    label: "删除历史记录",
    key: "delete",
    icon: renderIcon(Trash),
  };
  let children = [];
  children[0] = deleteChild;
  hisChild.forEach((item: any) => {
    children.push(item);
  });
  dropOptions.value = [
    {
      type: "group",
      label: "历史记录",
      key: "main",
      children: children,
    },
  ];
};

// 删除历史记录
const doDelete = () => {
  historyStore().historyData = [];
  dropOptions.value = [
    {
      type: "group",
      label: "历史记录",
      key: "main",
      children: [],
    },
  ];
};

// 渲染历史记录标签
const renderDropdownLabel = (option: DropdownOption) => {
  if (option.type === "group") {
    return option.label as VNodeChild;
  }
  return h(
    "a",
    {
      onClick: () => {
        option.key == "delete"
          ? doDelete()
          : doSearch(true, option.label?.toString());
      },
    },
    {
      default: () => option.label as VNodeChild,
    },
  );
};

// 菜单数据
const loggedMenuOptions: MenuOption[] = [
  {
    label: () =>
      h(
        NImage,
        {
          src: "/logo-svg.svg",
          width: 100,
          previewDisabled: true,
        },
        "",
      ),
    key: "logo",
  },
  {
    label: () =>
      h(
        NDropdown,
        {
          trigger: "focus",
          options: dropOptions.value,
          placement: "bottom-start",
          showArrow: false,
          renderLabel: renderDropdownLabel,
          style: {
            width: "calc(20vw - 40px)",
          },
        },
        [
          h(
            NInput,
            {
              modelValue: searchContent.value,
              onUpdateValue: (value: any) => {
                searchContent.value = value;
              },
              class: "h-input",
              round: true,
              placeholder: "世界那么大，我想去看看🕶️",
              onKeydown: (e) => {
                if (e.key == "Enter") {
                  doSearch(false);
                }
              },
              onFocus: () => {
                historyTabVisible.value = true;
              },
              onBlur: () => {
                historyTabVisible.value = false;
              },
            },
            {
              suffix: renderIcon(Search),
            },
          ),
        ],
      ),
    key: "search",
  },
  {
    label: renderPostBtn(),
    key: "post-btn",
  },
  {
    label: renderLiveBtn(),
    key: "live-btn",
  },
  {
    label: renderAvatar(userStore().avatar),
    key: "avatar",
    children: [
      {
        label: "退出登录",
        key: "logout",
      },
    ],
  },
];
const notLogmenuOptions: MenuOption[] = [
  {
    label: () =>
      h(
        NImage,
        {
          src: "/logo-svg.svg",
          width: 100,
          previewDisabled: true,
        },
        "",
      ),
    key: "logo",
  },
  {
    label: "登录",
    style: "margin-left:auto",
    key: "login",
  },
];

// 用户登录
const doLogin = () => {
  login(loginForm.phoneNum, loginForm.pwd).then(async (res: any) => {
    if (res.status_code == 200) {
      menuClass.value = "header-menu-login";
      userStore().token = res.accessToken;
      userStore().user_id = res.user_id;
      userStore().isLoggedIn = true;
      userStore().phoneNum = loginForm.phoneNum;
      userStore().name = res.name;
      userStore().avatar = res.avatar;
      userStore().gender = res.gender;
      userStore().signature = res.signature;
      userStore().background_image = res.background_image;
      message.success("登录成功");
      loginFormVisible.value = false;
      window.location.reload();
      router.push("/");
    } else {
      message.error(res.status_message);
    }
  });
};

// 用户注册
const doRegister = () => {
  //发请求
  if (registerForm.pwd == registerForm.repeatPwd) {
    let defaultAvatar =
      "http://s327crbzf.hn-bkt.clouddn.com/a74c50c16c528f561ecb3fc2f8e4856b9bd70c828616c558cad98be75a400c81.jpg";
    let defaultBg =
      "http://s327crbzf.hn-bkt.clouddn.com/d689982a1aa68d1242a802fc10228e9a8f17901980bf4251c295ad619592ad98.jpg";
    register(
      "匿名用户MOMO",
      0,
      registerForm.phoneNum,
      registerForm.pwd,
      "该用户暂时还没有个人简介哦",
      defaultAvatar,
      defaultBg,
    ).then(() => {
      message.success("注册成功");
    });
  } else {
    message.error("密码不一致，请重试");
  }
};

// 用户登出
const doLogout = () => {
  menuClass.value = "header-menu-logout";
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
  routeStore().name = "rec";
  router.push("/rec");
  // window.location.reload();
};

// 获取投稿dialogue
const getPostVideoForm = () => {
  isVideoFormVisible.value = true;
};

// 获取我要直播dialogue
const getLiveForm = () =>{
  isLiveFormVisible.value = true
}

// 投稿完成后的回调
const updatePostVideoVisible = (flag: boolean) => {
  isVideoFormVisible.value = flag;
};

// 数字判断
const onlyAllowNumber = (value: string) => !value || /^\d+$/.test(value);

// 路由选择index标志
watch(
  () => activeIndex.value,
  (newValue: any) => {
    switch (newValue) {
      case "login":
        loginFormVisible.value = true;
        break;
      case "logout":
        doLogout();
        break;
    }
  },
);
watch(
  () => loginFormVisible.value,
  (newValue: any) => {
    if (!newValue) {
      activeIndex.value = "";
    }
  },
);

import { useClipboard } from "@vueuse/core";
const { copy, copied } = useClipboard();
const url = ref<string>('')
// 直播链接复制成功回调
const handleCopy = () => {
  copy(url.value);
  message.success("复制成功");
  isLiveFormVisible.value = false;
};

</script>

<template>
  <div>
    <NMenu
      v-model:value="activeIndex"
      :class="menuClass"
      mode="horizontal"
      :options="userStore().isLoggedIn ? loggedMenuOptions : notLogmenuOptions"
    />
    <NModal
      v-model:show="loginFormVisible"
      class="custom-card"
      preset="card"
      :closable="false"
      size="huge"
      :bordered="false"
    >
      <NTabs
        class="card-tabs"
        default-value="login"
        size="large"
        animated
        pane-wrapper-style="margin: 0 -4px"
        pane-style="padding-left: 4px; padding-right: 4px; box-sizing: border-box;"
      >
        <NTabPane name="login" tab="登录">
          <NForm :model="loginForm">
            <NFormItemRow label="手机号">
              <NInput
                v-model:value="loginForm.phoneNum"
                maxlength="11"
                :allow-input="onlyAllowNumber"
                placeholder="请输入数字"
                autocomplete="off"
                clearable
              />
            </NFormItemRow>
            <NFormItemRow label="密码">
              <NInput
                v-model:value="loginForm.pwd"
                autocomplete="off"
                maxlength="30"
                type="password"
                placeholder="请输入"
                show-password
                clearable
              />
            </NFormItemRow>
          </NForm>
          <NButton type="primary" block secondary strong @click="doLogin">
            登录
          </NButton>
        </NTabPane>
        <NTabPane name="register" tab="注册">
          <NForm :model="registerForm">
            <NFormItemRow label="手机号">
              <NInput
                v-model:value="registerForm.phoneNum"
                autocomplete="off"
                maxlength="11"
                :allow-input="onlyAllowNumber"
                placeholder="请输入数字"
                clearable
              />
            </NFormItemRow>
            <NFormItemRow label="密码">
              <NInput
                v-model:value="registerForm.pwd"
                maxlength="30"
                autocomplete="off"
                type="password"
                placeholder="请输入"
                show-password
                clearable
              />
            </NFormItemRow>
            <NFormItemRow label="重复密码">
              <NInput
                v-model:value="registerForm.repeatPwd"
                maxlength="30"
                autocomplete="off"
                type="password"
                placeholder="请输入"
                show-password
                clearable
              />
            </NFormItemRow>
          </NForm>
          <NButton type="primary" block secondary strong @click="doRegister">
            注册
          </NButton>
        </NTabPane>
      </NTabs>
    </NModal>
    <PostVedioCom
      :video-form-visible="isVideoFormVisible"
      @visible-update="updatePostVideoVisible"
    />
    <PostVedioCom
      :video-form-visible="isVideoFormVisible"
      @visible-update="updatePostVideoVisible"
    />
    <NModal :closable="false" v-model:show="isLiveFormVisible" preset="dialog" title="我要直播">
      您的直播推流地址为: {{ url }} 
      <template #action>
        <NSpin v-if="copied">
          <NButton round> 复制中 </NButton>
        </NSpin>
        <NButton v-else round @click="handleCopy"> 复制 </NButton>
      </template>
    </NModal>
  </div>
</template>

<style lang="scss">
.history-card {
  width: 50vw;
}

.custom-card {
  width: 30vw;
}
.header-menu-login {
  width: 100%;
  height: 60px;
  .n-menu-item {
    height: 100%;
    .n-menu-item-content {
      height: 100%;
    }
  }
  .n-menu-item:nth-child(1) .n-menu-item-content {
    width: 160px;
  }

  .n-menu-item:nth-child(2) .n-menu-item-content {
    width: 20vw;
  }
  .n-menu-item:nth-child(3) {
    margin-left: auto;
  }
}

.header-menu-logout {
  width: 100%;
  height: 60px;
  .n-menu-item {
    height: 100%;
    .n-menu-item-content {
      height: 100%;
    }
  }
  .n-menu-item:nth-child(1) .n-menu-item-content {
    width: 160px;
  }
  .n-menu-item:nth-child(2) {
    margin-left: auto;
    width: 20vw;
  }
}

.flex-grow {
  flex-grow: 1;
}

</style>
