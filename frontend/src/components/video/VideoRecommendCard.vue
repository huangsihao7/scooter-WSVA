<script setup lang="ts">
import { NCard, NImage } from 'naive-ui';
import { NButton, NIcon} from "naive-ui";
import {
  Heart
} from "@vicons/ionicons5";
import { VideoType } from '@/apis/interface';
interface propsType{
    recommendlists: Array<VideoType>
}
const props = defineProps<propsType>()
</script>

<template>
    <NCard class="video-card" hoverable v-for="(video,index) in props.recommendlists" :key="index">
        <NImage  width="100" :src="video.cover_url"  preview-disabled/>
        <div class="video-info">
            <span class="title">{{ video.title }}</span>
            <div class="footer">
                <div class="like">
                    <NButton class="btn" text color="#3f3f3f" >
                        <NIcon>
                            <Heart />
                        </NIcon>
                    </NButton>
                    <p> {{ video.favorite_count }} </p>
                </div>
                <span> @ {{ video.author.name }}</span>
            </div>
        </div>
    </NCard>
</template>

<style lang="scss">
.video-card{
    height: 100%;
    .n-card__content{
        display: flex;
        flex-direction: row;
        justify-content: start;
    }
    
    .n-image{
        // float: left;
        width: 100px;
        height: 100px;
    }
    .video-info{
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        margin-left: 14px;
        .title{
            display: block;
            font-size: 1.2rem;
            font-weight: bold;
            text-align: left;
        }
        .footer{
            color: #333;
            display: flex;
            justify-content: space-between;
            .like{
                width: 20%;
                .btn, p{
                    display: inline;
                }
                p{
                    margin-left: 5px;
                }
            }
        }
    }
}
</style>
