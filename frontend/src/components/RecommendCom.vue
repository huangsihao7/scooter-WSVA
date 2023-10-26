<!--
 * @Author: Xu Ning
 * @Date: 2023-10-26 18:39:00
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-26 19:19:09
 * @Description: 
 * @FilePath: \scooter-wsva\frontend\src\components\RecommendCom.vue
-->

<script lang="ts" setup>
import dplayer from '@/components/Video.vue';
import Hls from 'hls.js';
import { ref, reactive } from 'vue'
const dplayerObj = reactive({
    video: {
        url: 'http://127.0.0.1:8080/WeChat_20231026155918.mp4', //视频地址
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
    <dplayer :video="dplayerObj.video" :danmaku="dplayerObj.danmaku" :contextmenu="dplayerObj.contextmenu"
      :highlight="dplayerObj.highlight" />
</template>

<style scoped>
.dplayer{
    height: calc(100vh - 60px);
}
</style>