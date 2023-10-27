<script lang="ts" setup>
import { ref } from 'vue';
import Comment from './Comment.vue'

const commentlist = ref<Array<any>>([
    {
  avatar: '',
  name:'kkxx',
  content:'货币是为了提高交易效率而用于交换的中介商品。货币有多种形式，如贝壳粮食等自然物、金属纸张等加工品、银行卡信用卡等磁条卡、移动支付加密货币等APP。',
  likenum:2,
  date:'3天前'
},
{
  avatar: '',
  name:'kkxx',
  content:'货币是为了提高交易效率而用于交换的中介商品。货币有多种形式，如贝壳粮食等自然物、金属纸张等加工品、银行卡信用卡等磁条卡、移动支付加密货币等APP。',
  likenum:2,
  date:'3天前'
},
{
  avatar: '',
  name:'kkxx',
  content:'货币是为了提高交易效率而用于交换的中介商品。货币有多种形式，如贝壳粮食等自然物、金属纸张等加工品、银行卡信用卡等磁条卡、移动支付加密货币等APP。',
  likenum:2,
  date:'3天前'
},
{
  avatar: '',
  name:'kkxx',
  content:'货币是为了提高交易效率而用于交换的中介商品。货币有多种形式，如贝壳粮食等自然物、金属纸张等加工品、银行卡信用卡等磁条卡、移动支付加密货币等APP。',
  likenum:2,
  date:'3天前'
}
]);

</script>

<template>
    <div class="comment-lists" v-for="(comment, index) in commentlist" :key="index">
        <Comment class="comment" :comment="comment" />
    </div>

</template>
<style lang="scss" scoped>
.comment-lists{
    .comment{
        margin-top: 10px;
    }
}
</style>