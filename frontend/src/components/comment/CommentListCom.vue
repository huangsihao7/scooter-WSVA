<script lang="ts" setup>
import CommentCom from "./CommentCom.vue";
import { CommentType } from "@/apis/interface";

interface propsType {
  commentlists: Array<CommentType>;
}

const props = defineProps<propsType>();
const emit = defineEmits(["delete-comment"]);
const deleteFunc = (comment_id: number) => {
  console.log('22222',comment_id)
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
