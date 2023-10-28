<!--
 * @Author: Xu Ning
 * @Date: 2023-10-26 15:26:18
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-28 17:11:57
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\components\video\PostVideo.vue
-->
<script setup lang="ts">
import { reactive, onMounted, ref } from 'vue'
import { NUpload, NUploadDragger, NIcon } from 'naive-ui';
import { baseURl } from '@/axios';
import type { UploadFileInfo } from 'naive-ui'
import { ElMessage } from 'element-plus';
import { CloudUpload   } from '@vicons/ionicons5'

interface propsType {
    VideoFormVisible: boolean
}
interface ClassifyList {
  label: string
  value: string
}

const props = defineProps<propsType>()
const emit = defineEmits(['visible-update'])
const videoForm = reactive({
    title: '',
    category:'',
    url:'',
    coverUrl:''
})

// 分类下拉框数据
const classifyList = ref<Array<ClassifyList>>([{
    value: 'Option1',
    label: '娱乐',
  },
  {
    value: 'Option2',
    label: '体育',
  },
  {
    value: 'Option3',
    label: '美食',
  },
  {
    value: 'Option4',
    label: '二次元',
  },
  {
    value: 'Option5',
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

    emit('visible-update', false)
}

const handleCancelVideo = () =>{
    emit('visible-update', false)
}

const beforeUpload = (data: { file: UploadFileInfo, fileList: UploadFileInfo[] }) =>  {
    console.log(data.file,data.fileList)
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

//上传成功回调
const handleVideoSuccess = (data: { file: UploadFileInfo }) => {
 console.log(data.file);
 //后台上传数据
//  if (res.success == true) {  
//    videoForm.url = url;    //上传成功后端返回视频地址 回显
//  } else {
//    ElMessage({
//        message: res.data.msg,
//        type: 'error'
//    })
//  }
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
            <el-form-item label="标题">
                <el-input clearable v-model="videoForm.title" />
            </el-form-item>
            <el-form-item label="上传">
                <n-upload
                    multiple
                    directory-dnd
                    :action= postBaseURl
                    @before-upload="beforeUpload"
                    @on-finish = "handleVideoSuccess"
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