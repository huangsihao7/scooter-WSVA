<!--
 * @Author: Xu Ning
 * @Date: 2023-10-22 19:33:20
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-26 18:34:47
 * @Description: 
 * @FilePath: \scooter-wsva\frontend\src\components\Video.vue
-->

<template>
  <div ref="videoRef"></div>
</template>
     
<script setup>
import DPlayer from 'dplayer'
import Hls from 'hls.js';
import { ref, reactive, onBeforeUnmount, onMounted } from 'vue'

const videoRef = ref()
const state = reactive({
  instance: null
})

const props = defineProps({
  // 是否自动播放
  autoplay: {
    type: Boolean,
    default: false
  },
  // 主题色
  theme: {
    type: String,
    default: '#0093ff'
  },
  // 视频是否循环播放
  loop: {
    type: Boolean,
    default: true
  },
  // 语言(可选值: 'en', 'zh-cn', 'zh-tw')
  lang: {
    type: String,
    default: 'zh-cn'
  },
  // 是否开启截图(如果开启，视频和视频封面需要允许跨域)
  screenshot: {
    type: Boolean,
    default: false
  },
  // 是否开启热键
  hotkey: {
    type: Boolean,
    default: true
  },
  // 视频是否预加载(可选值: 'none', 'metadata', 'auto')
  preload: {
    type: String,
    default: 'auto'
  },
  // 默认音量
  volume: {
    type: Number,
    default: 0.7
  },
  // 可选的播放速率，可以设置成自定义的数组
  playbackSpeed: {
    type: Array,
    default: [0.5, 0.75, 1, 1.25, 1.5, 2]
  },
  // 在左上角展示一个 logo，你可以通过 CSS 调整它的大小和位置
  logo: {
    type: String,
    default: 'http://jiuaibuni.top/wp-content/uploads/2020/12/logo.png'
  },
  // 视频信息
  video: {
    type: Object,
    default: {
      url: 'https://api.dogecloud.com/player/get.m3u8?vcode=5ac682e6f8231991&userId=17&ext=.m3u8', //视频地址
      type: 'customHls',
      customType: {
        customHls: function (video, player) {
          const hls = new Hls(); //实例化Hls  用于解析m3u8
          hls.loadSource(video.src);
          hls.attachMedia(video);
        }
      }
    },
  },
  // 外挂字幕
  subtitle: {
    type: Object,
    default: {}
  },
  // 显示弹幕
  danmaku: {
    type: Object,
    default: {}
  },
  // 自定义右键菜单
  contextmenu: {
    type: Array,
    default: []
  },
  // 自定义进度条提示点
  highlight: {
    type: Array,
    default: []
  },
  // 阻止多个播放器同时播放，当前播放器播放时暂停其他播放器
  mutex: {
    type: Boolean,
    default: true
  }
})
onMounted(() => {
  let player = {
    container: videoRef.value,
    autoplay: props.autoplay,
    theme: props.theme,
    loop: props.loop,
    lang: props.lang,
    screenshot: props.screenshot,
    hotkey: props.hotkey,
    preload: props.preload,
    volume: props.volume,
    playbackSpeed: props.playbackSpeed,
    logo: props.logo,
    video: props.video,
    contextmenu: props.contextmenu,
    highlight: props.highlight,
    mutex: props.mutex,
  }
  if (props.subtitle.url) {
    player.subtitle = props.subtitle
  }
  if (props.danmaku) {
    player.danmaku = props.danmaku
  }
  console.log(player);
  state.instance = new DPlayer(player)
})
// 销毁
onBeforeUnmount(() => {
  state.instance.destroy()
})


</script>
     
<style lang='scss' scoped>
</style>
