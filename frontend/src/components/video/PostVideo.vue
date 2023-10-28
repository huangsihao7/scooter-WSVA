<!--
 * @Author: Xu Ning
 * @Date: 2023-10-26 15:26:18
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-28 12:57:53
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\components\video\PostVideo.vue
-->
<script setup lang="ts">
import { reactive, onMounted, ref } from 'vue'
import { UploadFilled } from '@element-plus/icons-vue'
import { baseURl } from '@/axios';
import { ElMessage } from 'element-plus';

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
    selectValue:'',
    url:''
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

//上传前回调
 const beforeUploadVideo = (file:any) => {
    var fileSize = file.size / 1024 / 1024 < 100;   //控制大小  修改100的值即可
    if (
    [
        "video/mp4",
        "video/ogg",
        "video/flv",
        "video/avi",
        "video/wmv",
        "video/rmvb",
        "video/mov",
    ].indexOf(file.type) == -1     //控制格式
    ) {
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
    // isShowUploadVideo = false;
}

//进度条
const uploadVideoProcess = (event:any, file:any, fileList:any) => {    //注意在data中添加对应的变量名
    console.log(event, file, fileList)
}
//上传成功回调
const handleVideoSuccess = (res:any, file:any) => {
 
      console.log(res, file);
      //后台上传数据
      if (res.success == true) {  
        videoForm.url = res.data.url;    //上传成功后端返回视频地址 回显
      } else {
        ElMessage({
            message: res.data.msg,
            type: 'error'
        })
      }
    }

// 点击投稿的回调函数
const handlePostVideo = () => {
    console.log(videoForm)
    emit('visible-update', false)
}

const handleCancelVideo = () =>{
    emit('visible-update', false)
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
        <el-form label-position="right" :model="videoForm" label-width="100px">
            <el-form-item label="视频标题">
                <el-input clearable v-model="videoForm.title" />
            </el-form-item>
            <el-form-item label="视频上传">
                <el-upload
                    class="upload-demo"
                    drag
                    :action = postBaseURl
                    :on-progress="uploadVideoProcess"
                    :before-upload="beforeUploadVideo"
                    :on-success="handleVideoSuccess"
                    multiple
                >
                    <el-icon class="el-icon--upload"><upload-filled /></el-icon>
                    <div class="el-upload__text">
                    Drop file here or <em>click to upload</em>
                    </div>
                    <template #tip>
                        <div class="el-upload__tip">
                            视频大小在100MB以内
                        </div>
                    </template>
                </el-upload>
            </el-form-item>
            <el-form-item label="分类选择">
                <el-select v-model="videoForm.selectValue" class="m-2" placeholder="选择">
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
.el-input{
    width: 230px;
}

.el-form{
    margin-left: calc((30vw - 430px)/2);
}

</style>