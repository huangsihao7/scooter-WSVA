<!--
 * @Author: Xu Ning
 * @Date: 2023-10-26 18:39:00
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-28 16:49:33
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\components\video\VideoPlus.vue
-->

<script lang="ts" setup>
import dplayer from '@/components/video/Video.vue';
import Hls from 'hls.js';
import { ref, reactive, onMounted } from 'vue'
import  { NIcon, NButton } from 'naive-ui'
import { Heart, ArrowRedo, ChatbubbleEllipses, Star  } from '@vicons/ionicons5'
import { VideoType } from '@/apis/interface'

interface propsType {
  video: VideoType,
  key:number
}

const props = defineProps<propsType>()

const emit = defineEmits(['comment-visible-update'])
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
    }
])

const dplayerObj = reactive({
    autoplay:false,
    video: {
        url: props.video.playUrl, //视频地址
        type: 'mp4',
        customType: {
            customHls: function (video:any, player:any) {
                console.log(player)
                const hls = new Hls(); //实例化Hls  用于解析m3u8
                hls.loadSource(video.src);
                hls.attachMedia(video);
            }
        }
    },
    // danmaku: {
    //     id: '9E2E3368B56CDBB4',
    //     api: 'https://api.prprpr.me/dplayer/',
    //     token: 'tokendemo',
    //     maximum: 1000,
    //     addition: ['https://api.prprpr.me/dplayer/v3/bilibili?aid=4157142'],
    //     user: 'DIYgod',
    //     bottom: '15%',
    //     unlimited: true,
    //     speedRate: 0.5,
    // },
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

// 喜欢按钮的操作
const handleLikeBtn = () =>{
    if(!props.video.isFavorite){
        props.video.favoriteCount ++
    }else{
        props.video.favoriteCount --
    }
    if(!props.video.isFavorite){
        likeAnimateClass.value = 'animate__heartBeat'
    }
    else{
        likeAnimateClass.value = ''
    }
    props.video.isFavorite = !props.video.isFavorite
    
    // TODO: 发请求
}

// 收藏按钮的操作
const handleCollectBtn = () =>{
    if(!props.video.isCollect){
        props.video.collectCount ++
    }else{
        props.video.collectCount --
    }
    if(!videoUrls.value[0].isCollect){
        collectAnimateClass.value = 'animate__heartBeat'
    }
    else{
        collectAnimateClass.value = ''
    }
    videoUrls.value[0].isCollect = !videoUrls.value[0].isCollect
    // TODO: 发请求
}

// 评论按钮的操作
const handleCommentBtn = () =>{
    commentVisible.value = !commentVisible.value
    emit('comment-visible-update')
    
}

// 分享按钮的操作
const handleShareBtn = () =>{
    shareVisible.value = !shareVisible.value
}

const likeAnimateClass = ref<String>('')
const collectAnimateClass = ref<String>('')
const commentAnimateClass = ref<String>('')
const shareAnimateClass = ref<String>('')
const commentVisible = ref<boolean>(false)
const shareVisible = ref<boolean>(false)
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
                    <p class="title">@{{props.video.author.name}}</p>
                    <p class="time">{{ props.video.createTime }}</p>
                </div>
                <div class="content">{{ videoUrls[0].content }}</div>
            </div>
            <div class='video-interaction-box'>
                <div class="like">
                    <div :class=likeAnimateClass>
                        <n-button v-if="props.video.isFavorite" class="btn" text @click="handleLikeBtn" color="rgb(254, 44, 85)" >
                            <n-icon  >
                                <Heart  />
                            </n-icon>
                        </n-button>
                        <n-button v-else class="btn" text @click="handleLikeBtn" color="#ffffff" >
                            <n-icon  >
                                <Heart  />
                            </n-icon>
                        </n-button>
                    </div>
                    <p>{{ props.video.favoriteCount }}</p>
                </div>
                <div class="comment">
                    <div :class=commentAnimateClass>
                        <n-button class="btn" text @click="handleCommentBtn" color="#fff" >
                            <n-icon  >
                                <ChatbubbleEllipses  />
                            </n-icon>
                        </n-button>
                    </div>
                    <p>{{ props.video.commentCount }}</p>
                </div>
                <div class="collect">
                    <div :class=collectAnimateClass>
                        <n-button v-if="videoUrls[0].isCollect" class="btn" text @click="handleCollectBtn" color="#ffb802" >
                            <n-icon  >
                                <Star  />
                            </n-icon>
                        </n-button>
                        <n-button v-else class="btn" text @click="handleCollectBtn" color="#ffffff" >
                            <n-icon  >
                                <Star  />
                            </n-icon>
                        </n-button>
                    </div>
                    <p>{{ props.video.collectCount }}</p>
                </div>
                <div class="share">
                    <div :class=shareAnimateClass>
                        <n-button class="btn" text @click="handleShareBtn" color="#fff" >
                            <n-icon  >
                                <ArrowRedo  />
                            </n-icon>
                        </n-button>
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
        width: calc(100vw - 160px);
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
        .like, .collect, .comment, .share{
            margin-top: 10px;
            .btn{
                font-size: 35px
            }
            p{
                display: block;
                margin: 0;
            }
        }
    }
}
</style>