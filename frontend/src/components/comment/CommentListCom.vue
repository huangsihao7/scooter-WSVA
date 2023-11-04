<!--
 * @Author: Xu Ning
 * @Date: 2023-10-27 22:12:03
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-04 17:27:49
 * @Description: 评论卡片列表
 * @FilePath: \scooter-WSVA\frontend\src\components\comment\CommentListCom.vue
-->
<script lang="ts" setup>
import CommentCom from "./CommentCom.vue";
import { CommentType } from "@/apis/interface";

interface propsType {
  commentlists: Array<CommentType>;
}

const props = defineProps<propsType>();
const emit = defineEmits(["delete-comment"]);
const deleteFunc = (comment_id: number) => {
  emit("delete-comment", comment_id);
};
</script>

<template>
  <div class="comment-space">
    <div
      v-for="(comment, index) in props.commentlists"
      :key="index"
      class="comment-lists"
    >
      <CommentCom
        class="comment"
        :comment="comment"
        @delete-comment="deleteFunc"
      />
    </div>
  </div>
</template>
<style lang="scss" scoped>
.comment-space {
  height: 100%;
  .comment-lists {
    .comment {
      margin-top: 10px;
    }
  }

  .post-comment {
    position: fixed;
    width: 360px;
    bottom: 20px;
    background-color: antiquewhite;
  }
}
</style>
