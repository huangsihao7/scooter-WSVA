<!--
 * @Author: Xu Ning
 * @Date: 2023-10-27 22:00:03
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-01 12:53:10
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\components\comment\CommentCom.vue
-->
<script lang="ts" setup>
import { ref } from "vue";
import { Heart, ChatbubbleEllipses } from "@vicons/ionicons5";
import { NIcon, NButton, NSpace, NThing, NAvatar } from "naive-ui";
import { CommentType } from "@/apis/interface";

interface propsType {
  comment: CommentType;
}

const props = defineProps<propsType>();

const avatar = ref(true);
const header = ref(true);
const headerExtra = ref(true);
const action = ref(true);
</script>

<template>
  <NThing class="comment" content-indented>
    <template v-if="avatar" #avatar>
      <NAvatar round :src="props.comment.user.avatar"> </NAvatar>
    </template>
    <template v-if="header" #header>
      <span class="name"> {{ props.comment.user.name }} </span>
    </template>
    <template v-if="headerExtra" #header-extra >
      <span style="color:#666; font-size: 0.7rem;">{{ props.comment.create_date }}</span>
    </template>
    {{ props.comment.content }}
    <template v-if="action" #action>
      <NSpace>
        <NButton size="small">
          <template #icon>
            <NIcon>
              <Heart />
            </NIcon>
          </template>
          0
          <!-- {{ props.comment.likenum }} -->
        </NButton>
        <NButton size="small">
          <template #icon>
            <NIcon>
              <ChatbubbleEllipses />
            </NIcon>
          </template>
          评论
        </NButton>
      </NSpace>
    </template>
  </NThing>
</template>
<style lang="scss" scoped>
.n-thing {
  text-align: left;
  .name {
    font-weight: bold;
  }
}

.comment{
  background-color: #0000000a;
  border-radius: 25px;
  padding: 10px 15px;
}
</style>
