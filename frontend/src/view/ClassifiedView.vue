<script setup lang="ts">
import { computed } from 'vue';
import { routeStore } from '@/stores/route'


const routeName = computed(()=>routeStore().name)

</script>


<template>
  {{routeName}}
</template>
    
    

<style scoped>

</style>
