<!--
 * @Author: Xu Ning
 * @Date: 2023-10-22 19:33:20
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-02 19:07:17
 * @Description: 视频基础组件
 * @FilePath: \scooter-WSVA\frontend\src\components\video\VideoCom.vue
-->

<template>
  <div ref="videoRef"></div>
</template>

<script setup lang="ts">
import DPlayer from "dplayer";
import Hls from "hls.js";
import { onBeforeUnmount, onMounted, onUpdated, reactive, ref } from "vue";
import { videoStore } from "@/stores/video";

const videoRef = ref();
const state: any = reactive({
  instance: null,
});

const props = defineProps({
  // 是否是第一个视频
  // firstStart:{
  //   type: Boolean,
  //   default: false
  // },
  videoId: {
    type: Number,
    default: -1,
  },
  videoIndex: {
    type: Number,
    default: -1,
  },
  onPlayIndex: {
    type: Number,
    default: -1,
  },
  // 是否自动播放
  autoplay: {
    type: Boolean,
    default: false,
  },
  // 主题色
  theme: {
    type: String,
    default: "#0093ff",
  },
  // 视频是否循环播放
  loop: {
    type: Boolean,
    default: true,
  },
  // 语言(可选值: 'en', 'zh-cn', 'zh-tw')
  lang: {
    type: String,
    default: "zh-cn",
  },
  // 是否开启截图(如果开启，视频和视频封面需要允许跨域)
  screenshot: {
    type: Boolean,
    default: false,
  },
  // 是否开启热键
  hotkey: {
    type: Boolean,
    default: true,
  },
  // 视频是否预加载(可选值: 'none', 'metadata', 'auto')
  preload: {
    type: String,
    default: "auto",
  },
  // 默认音量
  volume: {
    type: Number,
    default: 0.7,
  },
  // 可选的播放速率，可以设置成自定义的数组
  playbackSpeed: {
    type: Array,
    default: function () {
      return [0.5, 0.75, 1, 1.25, 1.5, 2];
    },
  },
  // 在左上角展示一个 logo，你可以通过 CSS 调整它的大小和位置
  // logo: {
  //   type: String,
  //   default: "http://127.0.0.1:8080/OIP-C.jpg",
  // },
  // 视频信息
  // video: {
  //   type: Object,
  //   default: function () {
  //     return {
  //       url: "", //视频地址
  //       type: "mp4",
  //       customType: {
  //         customHls: function (video: any, player: any) {
  //           const hls = new Hls(); //实例化Hls  用于解析m3u8
  //           hls.loadSource(video.src);
  //           hls.attachMedia(video);
  //         },
  //       },
  //     };
  //   },
  // },
  video: {
    type: Object,
    // eslint-disable-next-line vue/require-valid-default-prop
    default: {
      url: "", //视频地址
      type: "mp4",
      customType: {
        customHls: function (video: any, _player: any) {
          const hls = new Hls(); //实例化Hls  用于解析m3u8
          hls.loadSource(video.src);
          hls.attachMedia(video);
        },
      },
    },
  },
  // 外挂字幕
  subtitle: {
    type: Object,
    // eslint-disable-next-line vue/require-valid-default-prop
    default: {},
  },
  // 显示弹幕
  danmaku: {
    type: Object,
    // eslint-disable-next-line vue/require-valid-default-prop
    default: {},
  },
  // 自定义右键菜单
  contextmenu: {
    type: Array,
    // eslint-disable-next-line vue/require-valid-default-prop
    default: [],
  },
  // 自定义进度条提示点
  highlight: {
    type: Array,
    // eslint-disable-next-line vue/require-valid-default-prop
    default: [],
  },
  // 阻止多个播放器同时播放，当前播放器播放时暂停其他播放器
  mutex: {
    type: Boolean,
    default: true,
  },
  //视频封面
  pic: {
    type: String,
    default: "http://127.0.0.1:8080/OIP-C.jpg",
  },
  preventClickToggle: {
    type: Boolean,
    default: false,
  },
});

onMounted(() => {
  let player: any = {
    container: videoRef.value,
    preventClickToggle: props.preventClickToggle,
    pic: props.pic,
    autoplay: props.autoplay,
    theme: props.theme,
    loop: props.loop,
    lang: props.lang,
    screenshot: props.screenshot,
    hotkey: props.hotkey,
    preload: props.preload,
    volume: props.volume,
    playbackSpeed: props.playbackSpeed,
    // logo: props.logo,
    video: props.video,
    contextmenu: props.contextmenu,
    highlight: props.highlight,
    mutex: props.mutex,
  };
  if (props.subtitle.url) {
    player.subtitle = props.subtitle;
  }
  if (props.danmaku) {
    player.danmaku = props.danmaku;
  }
  if (props.videoIndex == 0) {
    //自动播放开启
    player.autoplay = true;
    state.instance = new DPlayer(player);
    videoStore().video_id = props.videoId;
    // state.instance.video.play()
  } else {
    player.autoplay = false;
    state.instance = new DPlayer(player);
  }
});

onUpdated(() => {
  if (state.instance) {
    // 如果需要播放的不是当前video 则暂停
    if (props.onPlayIndex != props.videoIndex) {
      state.instance.video.pause();
    } else {
      state.instance.video.play();
      videoStore().video_id = props.videoId;
    }
  }
});

// 销毁
onBeforeUnmount(() => {
  if (state.instance != null) {
    state.instance.destroy();
  }
});
</script>

<style lang="scss" scoped></style>
