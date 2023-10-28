<!--
 * @Author: Xu Ning
 * @Date: 2023-10-27 22:00:03
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-27 23:02:56
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\components\comment\Comment.vue
-->
<script lang="ts" setup>
import { ref } from 'vue'
import { CashOutline as CashIcon, Heart, ChatbubbleEllipses } from '@vicons/ionicons5'
import { NIcon, NButton, NSpace,NThing,NAvatar } from 'naive-ui'

interface propsType {
  comment: CommentType
}
interface CommentType{
  avatar: string,
  name:string,
  content:string,
  likenum:number,
  date:string
}

const props = defineProps<propsType>()

const avatar = ref(true)
const header = ref(true)
const headerExtra = ref(true)
const action = ref(true)

</script>

<template>
    <n-thing content-indented>
      <template v-if="avatar" #avatar>
        <n-avatar :src="props.comment.avatar">
          <n-icon>
            <cash-icon />
          </n-icon>
        </n-avatar>
      </template>
      <template v-if="header" #header >
        <span class="name"> {{props.comment.name}} </span>
      </template>
      <template v-if="headerExtra" #header-extra>
        {{props.comment.date}}
      </template>
      {{ props.comment.content }}
      <template v-if="action" #action>
        <n-space>
          <n-button size="small">
            <template #icon>
              <n-icon>
                <Heart />
              </n-icon>
            </template>
            {{ props.comment.likenum }}
          </n-button>
          <n-button size="small">
            <template #icon>
              <n-icon>
                <ChatbubbleEllipses />
              </n-icon>
            </template>
            评论
          </n-button>
        </n-space>
      </template>
    </n-thing>
  </template>
<style lang="scss" scoped>
.n-thing{
  text-align: left;
  .name{
    font-weight: bold;
  }
}
</style>
  
