<!--
 * @Author: Xu Ning
 * @Date: 2023-10-26 18:39:00
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-27 12:37:01
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\components\RecommendCom.vue
-->

<script lang="ts" setup>
import dplayer from '@/components/Video.vue';
import Hls from 'hls.js';
import { ref, reactive, onMounted } from 'vue'
import { NIcon, NButton } from 'naive-ui'
import { Heart  } from '@vicons/ionicons5'


const videoUrls = ref<any>([
    {
        url:'http://127.0.0.1:8080/3.mp4',
        cover:'',
        username:'我是一个粉刷匠',
        createTime:'3天前',
        likes:'12',
        isLike:false,
        collect:'22',
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

const likeState = ref<boolean>(false)

const handleBtn = () =>{
    console.log('btnnnnn')
    videoUrls[0].isLike = !videoUrls[0].isLike
}

onMounted(() => {
})
</script>

<template>
    <div>
        <div class="video-container">
            <dplayer :video="dplayerObj.video" :danmaku="dplayerObj.danmaku" :contextmenu="dplayerObj.contextmenu"
            :highlight="dplayerObj.highlight" />
            <div class="video-info-box">
                <div class="header">
                    <p class="title">@{{videoUrls[0].username}}</p>
                    <p class="time">{{ videoUrls[0].createTime }}</p>
                </div>
                <div class="content">{{ videoUrls[0].content }}</div>
            </div>
            <div class="video-interaction-box">
                <div class="like">
                    <n-icon :component="Heart" size="40" color="rgb(254, 44, 85)"  :depth="1" />
                    <div class="animate__heartBeat">
                        <n-button   text style="font-size: 40px" @click="handleBtn">
                            <n-icon v-if="videoUrls[0].isLike" color="rgb(254, 44, 85)">
                                <Heart  />
                            </n-icon>
                            <n-icon v-else>
                                <Heart  />
                            </n-icon>
                        </n-button>
                        <p v-if="videoUrls[0].isLike">123</p>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped lang="scss">


.video-container{

    .el-button{
        z-index: 9999;
    }
    .dplayer{
        height: calc(100vh - 60px);
    }
    .video-info-box{
        display: block;
        height: auto;
        pointer-events: none;
        position: absolute;
        z-index: 2;
        width: calc(100vw - 200px);
        bottom: 48px;
        height: 100px;
        padding: 0 20px;
        .title, .time, .content{
            color: white;
        }

        .header{
            display: flex;
            .time{
                margin: 25px 0 0 15px;
                font-size: small;
            }
            .title{
                font-size: larger;
                font-weight: bold;
            }
        }

        .content{
            text-align: left;
        }
        

    }

    .video-interaction-box{
        display: block;
        height: auto;
        position: absolute;
        z-index: 999;
        width: 100px;
        bottom: 150px;
        right: 0;
        height: 400px;
        padding: 0 20px;
        .like{
            
            p{
                display: block;
                margin: 0;
            }
        }
    }
}

</style>