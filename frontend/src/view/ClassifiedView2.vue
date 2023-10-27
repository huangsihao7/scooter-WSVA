<!--
 * @Author: Xu Ning
 * @Date: 2023-10-26 20:05:39
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-27 18:53:18
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\view\ClassifiedView.vue
-->
<template>
  <div class="container">
    
    <div class="item" v-for="(item, index) in videoLists">
      <img @load="imgLoaded" :src="`${item.src}`" alt="mock images">
    </div>
  </div>
</template>

<script lang="ts" setup>
import {ref, reactive} from 'vue'
  import Masonry from 'masonry-layout';
  import { onMounted } from 'vue';
  let masonry:any;

  const videoLists = reactive([
    {
    src:'http://127.0.0.1:8080/6.jpg',
    width: 5500,
    height: 3500,
  },{
    src:'http://127.0.0.1:8080/7.jpg',
    width: 400,
    height: 400,
  },{
    src:'http://127.0.0.1:8080/8.jpg',
    width: 300,
    height: 300,
  },{
    src:'http://127.0.0.1:8080/9.jpg',
    width: 250,
    height: 300,
  },{
    src:'http://127.0.0.1:8080/10.jpg',
    width: 300,
    height: 200,
  },{
    src:'http://127.0.0.1:8080/11.jpg',
    width: 1000,
    height: 200,
  },{
    src:'http://127.0.0.1:8080/12.jpg',
    width: 350,
    height: 300,
  },{
    src:'http://127.0.0.1:8080/13.jpg',
    width: 250,
    height: 350,
  }])

  onMounted(() => {
    document.body.setAttribute('style', '');
    masonry = new Masonry('.container', {
      itemSelector: '.item',//将以此选择器对应的元素作为瀑布流内容元素
      columnWidth: '.container .item',//将以此选择器对应的元素宽度作为瀑布流的列宽
      percentPosition: true,//支持百分比宽度
    });
  });


  function imgLoaded(e:any) {
    masonry.layout();//重新布局
  }

  //每2秒加载2张图片
  const imageList = ref(Array.from({ length: 5 }));

  setInterval(() => {
    imageList.value.push(...Array.from({ length: 2 }));
    masonry.reloadItems(); //https://masonry.desandro.com/methods.html
  }, 5000);

</script>