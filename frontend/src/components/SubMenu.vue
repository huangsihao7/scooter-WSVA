<template>
  <ElRow class="tac">
    <ElCol>
      <ElMenu
        :default-active="active"
        class="el-menu-vertical-demo"
        :router="true"
        @open="handleOpen"
        @close="handleClose"
      >
        <ElMenuItem index="rec">
          <div class="submenu-text">
            <ElIcon>
              <Star />
            </ElIcon>
            <span>推荐</span>
          </div>
        </ElMenuItem>
        <ElMenuItem index="follow">
          <div class="submenu-text">
            <ElIcon>
              <Connection />
            </ElIcon>
            <span>关注</span>
          </div>
        </ElMenuItem>
        <ElMenuItem index="user">
          <div class="submenu-text">
            <ElIcon>
              <House />
            </ElIcon>
            <span>我的</span>
          </div>
        </ElMenuItem>
        <ElDivider content-position="center">分类</ElDivider>
        <ElMenuItem index="hot">
          <div class="submenu-text">
            <ElIcon>
              <Medal />
            </ElIcon>
            <span>热门</span>
          </div>
        </ElMenuItem>
        <ElMenuItem index="recreation">
          <div class="submenu-text">
            <ElIcon>
              <Goblet />
            </ElIcon>
            <span>娱乐</span>
          </div>
        </ElMenuItem>
        <ElMenuItem index="sports">
          <div class="submenu-text">
            <ElIcon>
              <Basketball />
            </ElIcon>
            <span>体育</span>
          </div>
        </ElMenuItem>
        <ElMenuItem index="food">
          <div class="submenu-text">
            <ElIcon>
              <KnifeFork />
            </ElIcon>
            <span>美食</span>
          </div>
        </ElMenuItem>
        <ElMenuItem index="cartoon">
          <div class="submenu-text">
            <ElIcon>
              <Lollipop />
            </ElIcon>
            <span>二次元</span>
          </div>
        </ElMenuItem>
        <ElMenuItem index="knowledge">
          <div class="submenu-text">
            <ElIcon>
              <Reading />
            </ElIcon>
            <span>知识</span>
          </div>
        </ElMenuItem>
      </ElMenu>
    </ElCol>
  </ElRow>
</template>

<script lang="ts" setup>
import { routeStore } from "@/stores/route";
import { onMounted, ref } from "vue";
import {
  KnifeFork,
  Goblet,
  Lollipop,
  Connection,
  Medal,
  Reading,
  House,
  Basketball,
  Star,
} from "@element-plus/icons-vue";

const active = ref<String>("");

// 获取当前页面的末尾路由
function getLastSegmentFromRoute(route: string): string {
  const segments = route.split("/");
  return segments[segments.length - 1];
}

// 通过url确定menu的activekey
onMounted(() => {
  if (active.value == "") {
    const currentPath = window.location.pathname;
    let path = getLastSegmentFromRoute(currentPath);
    active.value = path;
    routeStore().name = path;
    console.log("path", path, active.value, routeStore().name);
  }
});

const handleOpen = (key: string, keyPath: string[]) => {
  console.log(key, keyPath);
  active.value = key;
  routeStore().name = key;
};
const handleClose = (key: string, keyPath: string[]) => {
  console.log(key, keyPath);
};
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
