<script lang="ts" setup>
import CommentCom from "./CommentCom.vue";
import { CommentType } from '@/apis/interface'

interface propsType{
  commentlists: Array<CommentType>
}

const props = defineProps<propsType>();

</script>

<template>
  <div
    v-for="(comment, index) in props.commentlists"
    :key="index"
    class="comment-lists"
  >
    <CommentCom class="comment" :comment="comment" />
  </div>
</template>
<style lang="scss" scoped>
.comment-lists {
  .comment {
    margin-top: 10px;
  }
}
</style>
