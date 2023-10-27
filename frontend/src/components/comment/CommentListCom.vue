<script lang="ts" setup>
import { ref } from 'vue';
import Comment from './Comment.vue'

const commentlist = ref<Array<any>>();

</script>

<template>
    <div v-for="(item, index) in commentlist">
        <Comment :comment = item/>
    </div>
</template>