<!--
 * @Author: Xu Ning
 * @Date: 2023-10-26 15:26:18
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-28 21:29:53
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\components\video\PostVideo.vue
-->
<script setup lang="ts">
import { reactive, onMounted, ref, computed } from 'vue'
import { NUpload, NUploadDragger, NIcon } from 'naive-ui';
import { baseURl } from '@/axios';
import type { UploadFileInfo } from 'naive-ui'
import { ElMessage } from 'element-plus';
import { CloudUpload   } from '@vicons/ionicons5'
import { userStore } from '@/stores/user'
import { postVideo } from '@/apis/video'

interface propsType {
    VideoFormVisible: boolean
}
interface ClassifyList {
  label: string
  value: number
}

const props = defineProps<propsType>()
const token = computed(()=>userStore().token)
const emit = defineEmits(['visible-update'])
const videoForm = reactive({
    title: '',
    category:1,
    url:'',
    coverUrl:''
})

// 分类下拉框数据
const classifyList = ref<Array<ClassifyList>>([{
    value: 1,
    label: '娱乐',
  },
  {
    value: 2,
    label: '体育',
  },
  {
    value: 3,
    label: '美食',
  },
  {
    value: 4,
    label: '二次元',
  },
  {
    value: 5,
    label: '知识',
}])

// 上传video的URL
const postBaseURl = baseURl + '/user/upload'

onMounted(() => {
    getClassifyList()
})

// 获取分类
const getClassifyList = () =>{

}
// 点击投稿的回调函数
const handlePostVideo = () => {
    postVideo(videoForm.url, videoForm.coverUrl, videoForm.title, videoForm.category).then((res:any)=>{
        console.log(res)
    })
    emit('visible-update', false)
}

const handleCancelVideo = () =>{
    emit('visible-update', false)
}

const beforeUpload = (data: { file: UploadFileInfo, fileList: UploadFileInfo[] }) =>  {
    var fileSize
    if(data.file.file){
        fileSize = data.file.file.size / 1024 / 1024 < 100;   //控制大小  修改100的值即可
        if (["video/mp4",
                "video/ogg",
                "video/flv",
                "video/avi",
                "video/wmv",
                "video/rmvb",
                "video/mov"].indexOf(data.file.file.type) == -1) {
                ElMessage({
                    message: '请上传正确的视频格式',
                    type: 'error'
                })
                return false;
            }
            if (!fileSize) {
                ElMessage({
                    message: '视频大小不能超过100MB',
                    type: 'error'
                })
                return false;
            }
    }
    return true
}

// 上传成功回调, 把url和coverurl添加
const handleFinish = ({file,event}: {file: UploadFileInfo,event?: ProgressEvent}) => {
      console.log('111111',file, event)
      if(event && event.target){
        let res:any = event.target
        if(res.response){
            let resJson = JSON.parse(res.response)
            let url = resJson.url
            let coverUrl = resJson.coverUrl
            videoForm.url = url
            videoForm.coverUrl = coverUrl
        }
      }
      return file
}

function uploadHeader (){
    return {
        Authorization: token.value
    }
}

</script>


<template>
    <el-dialog 
        v-model="props.VideoFormVisible" 
        title="投稿视频" 
        :show-close=false 
        :close-on-click-modal=false
        width="30%"
    >
        <el-form label-position="right" :model="videoForm" label-width="50px">
            <el-form-item label="描述">
                <el-input
                    clearable 
                    v-model="videoForm.title"
                    autosize
                    type="textarea"
                    placeholder="请输入"
                />
            </el-form-item>
            <el-form-item label="上传">
                <n-upload
                    multiple
                    directory-dnd
                    :action= postBaseURl
                    :headers = "uploadHeader"
                    @before-upload="beforeUpload"
                    @finish="handleFinish"
                >
                    <n-upload-dragger>
                    <div style="margin-bottom: 12px">
                        <n-icon size="48" :depth="3">
                            <CloudUpload />
                        </n-icon>
                    </div>
                    <n-text style="font-size: 16px">
                        点击或者拖动文件到该区域来上传
                    </n-text>
                    </n-upload-dragger>
                </n-upload>
            </el-form-item>
            <el-form-item label="分类">
                <el-select v-model="videoForm.category" class="m-2" placeholder="选择">
                    <el-option
                        v-for="item in classifyList"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value"
                    />
                </el-select>
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="handleCancelVideo">取消</el-button>
                <el-button type="primary" @click="handlePostVideo">
                发布
                </el-button>
            </span>
        </template>
    </el-dialog>
      
</template>
    
    

<style scoped>
.el-select{
    width:100%;
}

</style>