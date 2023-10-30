<template>
  <n-menu class="tac" :options="menuOptions" :default-value="defaultActive" @update:value="handleUpdateValue" />
</template>

<script lang="ts" setup>
import { routeStore } from "@/stores/route";
import {  ref, onBeforeMount,  h, Component } from "vue";

import { NMenu } from 'naive-ui'
import { NIcon } from 'naive-ui'
import type { MenuOption } from 'naive-ui'
import { RouterLink } from 'vue-router'
import {
  Diamond as RecIcon,
  Person as FollowIcon,
  Home as UserIcon,
  Rocket as HotIcon,
  GameController as RecreationIcon,
  Basketball as SportsIcon,
  FastFood as FoodIcon,
  PlanetSharp as CartoonIcon,
  Book as KnowledgeIcon,
} from '@vicons/ionicons5'

function renderIcon (icon: Component) {
  return () => h(NIcon, null, { default: () => h(icon) })
}

const menuOptions: MenuOption[] = [
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: 'rec',
            params: {
              lang: 'zh-CN'
            }
          }
        },
        { default: () => '推荐' }
      ),
    key: 'rec',
    icon: renderIcon(RecIcon)
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: 'follow',
            params: {
              lang: 'zh-CN'
            }
          }
        },
        { default: () => '关注' }
      ),
    key: 'follow',
    icon: renderIcon(FollowIcon)
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: 'user',
            params: {
              lang: 'zh-CN'
            }
          }
        },
        { default: () => '我的' }
      ),
    key: 'user',
    icon: renderIcon(UserIcon)
  },
  {
    key: 'divider-1',
    type: 'divider',
    props: {
      style: {
        marginLeft: '32px'
      }
    }
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: 'hot',
            params: {
              lang: 'zh-CN'
            }
          }
        },
        { default: () => '热门' }
      ),
    key: 'hot',
    icon: renderIcon(HotIcon)
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: 'recreation',
            params: {
              lang: 'zh-CN'
            }
          }
        },
        { default: () => '娱乐' }
      ),
    key: 'recreation',
    icon: renderIcon(RecreationIcon)
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: 'sports',
            params: {
              lang: 'zh-CN'
            }
          }
        },
        { default: () => '体育' }
      ),
    key: 'sports',
    icon: renderIcon(SportsIcon)
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: 'food',
            params: {
              lang: 'zh-CN'
            }
          }
        },
        { default: () => '食物' }
      ),
    key: 'food',
    icon: renderIcon(FoodIcon)
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: 'cartoon',
            params: {
              lang: 'zh-CN'
            }
          }
        },
        { default: () => '二次元' }
      ),
    key: 'cartoon',
    icon: renderIcon(CartoonIcon)
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: 'knowledge',
            params: {
              lang: 'zh-CN'
            }
          }
        },
        { default: () => '知识' }
      ),
    key: 'knowledge',
    icon: renderIcon(KnowledgeIcon)
  },
]

const defaultActive = ref<any>("");

// 获取当前页面的末尾路由
function getLastSegmentFromRoute(route: string): string {
  const segments = route.split("/");
  return segments[segments.length - 1];
}


const handleUpdateValue = (key: string) => {
    let keyStr = JSON.stringify(key)
    const withoutQuotes = keyStr.replace(/"/g, '');
    routeStore().name = withoutQuotes
  }


onBeforeMount(() => {
  let path = window.location.href
  defaultActive.value = getLastSegmentFromRoute(path)
  routeStore().name = defaultActive.value
})

</script>

<style scoped>
.tac {
  height: 100%;
}

.el-menu-vertical-demo {
  height: 100%;
}

.el-menu-item .submenu-text {
  margin-left: 20px;
}
</style>
