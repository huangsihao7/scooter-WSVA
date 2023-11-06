<!--
 * @Author: Xu Ning
 * @Date: 2023-10-25 16:22:40
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-06 13:40:45
 * @Description: 侧边导航栏组件
 * @FilePath: \scooter-WSVA\frontend\src\components\menu\SubMenu.vue
-->
<template>
  <NMenu
    v-model:value="routeValue"
    class="tac"
    :options="menuOptions"
    :default-value="defaultActive"
    @update:value="handleUpdateValue"
  />
</template>

<script lang="ts" setup>
import { h, onBeforeMount, ref, watch } from "vue";
import type { MenuOption } from "naive-ui";
import { NMenu } from "naive-ui";
import { RouterLink } from "vue-router";
import { routeStore } from "@/stores/route";

const defaultActive = ref<any>("");
const routeValue = ref<any>();
//菜单
const menuOptions: MenuOption[] = [
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: "rec",
            params: {
              lang: "zh-CN",
            },
          },
        },
        { default: () => "推荐" },
      ),
    key: "rec",
    icon: renderIcon('icon-faxian'),
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: "follow",
            params: {
              lang: "zh-CN",
            },
          },
        },
        { default: () => "关注" },
      ),
    key: "follow",
    icon: renderIcon('icon-dianzan2'),
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: "user",
            params: {
              lang: "zh-CN",
            },
          },
        },
        { default: () => "我的" },
      ),
    key: "user",
    icon: renderIcon('icon-wode'),
  },
  {
    key: "divider-1",
    type: "divider",
    props: {
      style: {
        marginLeft: "32px",
      },
    },
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: "hot",
            params: {
              lang: "zh-CN",
            },
          },
        },
        { default: () => "热门" },
      ),
    key: "hot",
    icon: renderIcon('icon-redu'),
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: "liveSquare",
            params: {
              lang: "zh-CN",
            },
          },
        },
        { default: () => "直播" },
      ),
    key: "liveSquare",
    icon: renderIcon('icon-zhibo'),
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: "recreation",
            params: {
              lang: "zh-CN",
            },
          },
        },
        { default: () => "娱乐" },
      ),
    key: "recreation",
    icon: renderIcon('icon-youxi'),
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: "sports",
            params: {
              lang: "zh-CN",
            },
          },
        },
        { default: () => "体育" },
      ),
    key: "sports",
    icon: renderIcon('icon-jiangpai'),
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: "food",
            params: {
              lang: "zh-CN",
            },
          },
        },
        { default: () => "食物" },
      ),
    key: "food",
    icon: renderIcon('icon-kafei'),
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: "cartoon",
            params: {
              lang: "zh-CN",
            },
          },
        },
        { default: () => "二次元" },
      ),
    key: "cartoon",
    icon: renderIcon('icon-xingqiu'),
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: "knowledge",
            params: {
              lang: "zh-CN",
            },
          },
        },
        { default: () => "知识" },
      ),
    key: "knowledge",
    icon: renderIcon('icon-tushu'),
  },
];

// 渲染图标
function renderIcon(iconfontName: string) {
  return () => h('span', {
    class: ['iconfont', iconfontName]
  });
}

// 获取当前页面的末尾路由
function getLastSegmentFromRoute(route: string): string {
  const segments = route.split("/");
  return segments[segments.length - 1];
}

// 更新值
const handleUpdateValue = (key: string) => {
  let keyStr = JSON.stringify(key);
  routeStore().name = keyStr.replace(/"/g, "");
};

watch(
  () => routeStore().name,
  (newValue: any) => (routeValue.value = newValue),
);

// 挂载之前获取路由
onBeforeMount(() => {
  let path = window.location.href;
  routeValue.value = getLastSegmentFromRoute(path);
  routeStore().name = routeValue.value;
});
</script>

<style>
.tac {
  height: 100%;
}

.n-menu-item-content__icon{
  margin-right: 0px !important;
}
</style>
