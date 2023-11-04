<!--
 * @Author: Xu Ning
 * @Date: 2023-10-27 22:00:03
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-04 17:27:33
 * @Description: 评论组件
 * @FilePath: \scooter-WSVA\frontend\src\components\comment\CommentCom.vue
-->
<script lang="ts" setup>
import { ref } from "vue";
import { Heart, ChatbubbleEllipses, Trash } from "@vicons/ionicons5";
import {
  NIcon,
  NButton,
  NSpace,
  NThing,
  NAvatar,
  useMessage,
  NMessageProvider,
} from "naive-ui";
import { CommentType } from "@/apis/interface";
import { userStore } from "@/stores/user";
import { doComment } from "@/apis/comment";
import { videoStore } from "@/stores/video";

interface propsType {
  comment: CommentType;
}

const props = defineProps<propsType>();

const message = useMessage();
const avatar = ref(true);
const header = ref(true);
const headerExtra = ref(true);
const action = ref(true);
const emit = defineEmits(["delete-comment"]);

// 删除我的评论
const deleteMyComment = () => {
  let comment_id = props.comment.comment_id;
  doComment(videoStore().video_id, 2, "", comment_id).then(() => {
    message.success("删除成功");
    emit("delete-comment", comment_id);
  });
};

const fake = () => {
  console.log("hiiiiiiiiiiiii");
  message.error("aaa");
};
</script>

<template>
  <NThing class="comment" content-indented>
    <template v-if="avatar" #avatar>
      <NAvatar round :src="props.comment.user.avatar"> </NAvatar>
    </template>
    <template v-if="header" #header>
      <span class="name"> {{ props.comment.user.name }} </span>
    </template>
    <template v-if="headerExtra" #header-extra>
      <span style="color: #666; font-size: 0.7rem">{{
        props.comment.create_date
      }}</span>
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
        <NMessageProvider placement="top-right">
          <NButton size="small" @click="fake">
            <template #icon>
              <NIcon>
                <ChatbubbleEllipses />
              </NIcon>
            </template>
            评论
          </NButton>
        </NMessageProvider>
        <NButton
          v-if="props.comment.user.id == userStore().user_id"
          size="small"
          @click="deleteMyComment"
        >
          <template #icon>
            <NIcon>
              <Trash />
            </NIcon>
          </template>
          删除
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

.comment {
  background-color: #0000000a;
  border-radius: 25px;
  padding: 10px 15px;
}
</style>
