<script setup lang="ts">
import { NCarousel, NCarouselItem } from 'naive-ui';
import dplayer from './Video.vue'
import { reactive,ref } from 'vue'

const videoUrls = ref<any>([
    {
        url:'http://127.0.0.1:8080/3.mp4',
        cover:'',
        username:'我是一个粉刷匠',
        createTime:'3天前',
        likes:'12',
        isLike:true,
        collect:'22',
        isCollect: true,
        isFollowed: false,
        title: '你好哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈',
        content:'在你拿着一堆早餐摆着慢慢吃，武汉人发出尖锐的爆鸣[捂脸]主要是那些炸物刚出炉最好吃了，你还等豆皮等半天[泪奔]来武汉过早，这种炸物建议到手就吃哦，热干面也不能放，拿到就要拌开'
    },
    {
        url:'http://127.0.0.1:8080/3.mp4',
        cover:'',
        username:'我是一个粉刷匠',
        createTime:'3天前',
        likes:'12',
        collect:'22',
        isFollowed: true,
        title: '你好哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈',
        content:'离开家睡觉了还能考古发掘按时的可见光和'
    },
    {
        url:'http://127.0.0.1:8080/WeChat_20231026155918.mp4',
        cover:'',
        username:'cccccccc',
        createTime:'3天前',
        likes:'12',
        collect:'22',
        isFollowed: true,
        title: '你好哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈',
        content:'离开家睡觉了还能考古发掘按时的可见光和'
    }
])

const dplayerObj = reactive({
    autoplay:true,
    video: {
        url: videoUrls.value[0].url, //视频地址
        type: 'mp4',
        customType: {
            customHls: function (video:any, player:any) {
                const hls = new Hls(); //实例化Hls  用于解析m3u8
                hls.loadSource(video.src);
                hls.attachMedia(video);
            }
        }
    },
    danmaku: {
        id: '9E2E3368B56CDBB4',
        api: 'https://api.prprpr.me/dplayer/',
        token: 'tokendemo',
        maximum: 1000,
        addition: ['https://api.prprpr.me/dplayer/v3/bilibili?aid=4157142'],
        user: 'DIYgod',
        bottom: '15%',
        unlimited: true,
        speedRate: 0.5,
    },
    contextmenu: [
        {
            text: 'custom1',
            link: 'https://github.com/DIYgod/DPlayer',
        },
        {
            text: 'custom2',
            click: (player:any) => {
                console.log(player);
            },
        },
    ],
    highlight: [
        {
        time: 20,
        text: '这是第 20 秒',
        },
        {
        time: 120,
        text: '这是 2 分钟',
        },
    ],
})


</script>

<template>
    <n-carousel
    class="vedeolists"
        direction="vertical"
        dot-placement="right"
        mousewheel
        show-arrow
        :show-dots="false"
        style="width: 100%; height: calc(100vh - 60px);"
    >
        <dplayer class="video-com" :video="dplayerObj.video" :danmaku="dplayerObj.danmaku" :contextmenu="dplayerObj.contextmenu"
            :highlight="dplayerObj.highlight" />
        <dplayer class="video-com" :video="dplayerObj.video" :danmaku="dplayerObj.danmaku" :contextmenu="dplayerObj.contextmenu"
            :highlight="dplayerObj.highlight" />
        <img
        class="carousel-img"
        src="https://naive-ui.oss-cn-beijing.aliyuncs.com/carousel-img/carousel1.jpeg"
        >
        <img
        class="carousel-img"
        src="https://naive-ui.oss-cn-beijing.aliyuncs.com/carousel-img/carousel2.jpeg"
        >
        <img
        class="carousel-img"
        src="https://naive-ui.oss-cn-beijing.aliyuncs.com/carousel-img/carousel3.jpeg"
        >
        <img
        class="carousel-img"
        src="https://naive-ui.oss-cn-beijing.aliyuncs.com/carousel-img/carousel4.jpeg"
        >
    </n-carousel>
</template>

<style scoped>
.carousel-img {
width: 100%;
height: 240px;
object-fit: cover;
}

.vedeolists{
    .video-com{
        width: auto;
        height: calc(100vh - 60px);
    }
}
</style>