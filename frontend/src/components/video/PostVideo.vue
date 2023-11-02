<!--
 * @Author: Xu Ning
 * @Date: 2023-10-26 15:26:18
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-02 15:54:52
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\components\video\PostVideo.vue
-->
<script setup lang="ts">
import { reactive, onMounted, ref, computed } from "vue";
import {
  NUpload,
  NUploadDragger,
  NIcon,
  NText,
  NModal,
  NForm,
  NFormItemRow,
  NInput,
  NButton,
  NSelect,
} from "naive-ui";
import { baseURL } from "@/axios";
import type { UploadFileInfo } from "naive-ui";
import { CloudUpload, DiceOutline, Dice } from "@vicons/ionicons5";
import { userStore } from "@/stores/user";
import { postVideo } from "@/apis/video";
import { useMessage } from "naive-ui";
const message = useMessage();
interface propsType {
  videoFormVisible: boolean;
}
interface ClassifyList {
  label: string;
  value: number;
}

const props = defineProps<propsType>();
const emit = defineEmits(["visible-update"]);

// 表单是否可见
const formVisible = computed(() => props.videoFormVisible);
// 用户token
const token = computed(() => userStore().token);
// 表单数据
const videoForm = reactive({
  title: "",
  category: 1,
  url: "",
  coverUrl: "",
});
// 上传绑定
const fileUploadRef = ref();
const titleIptRef = ref();

// 分类下拉框数据
const classifyList = ref<Array<ClassifyList>>([
  {
    value: 1,
    label: "娱乐",
  },
  {
    value: 2,
    label: "体育",
  },
  {
    value: 3,
    label: "美食",
  },
  {
    value: 4,
    label: "二次元",
  },
  {
    value: 5,
    label: "知识",
  },
]);

// 上传video的URL
const postBaseURL = baseURL + "/user/upload";

onMounted(() => {
  getClassifyList();
});

// 获取分类
const getClassifyList = () => {};

// 点击投稿的回调函数
const handlePostVideo = () => {
  if (
    videoForm.url != "" &&
    videoForm.coverUrl != "" &&
    videoForm.title != "" &&
    videoForm.category
  ) {
    postVideo(
      videoForm.url,
      videoForm.coverUrl,
      videoForm.title,
      videoForm.category,
    ).then((res: any) => {
      fileUploadRef.value.clear();
      titleIptRef.value.clear();
      message.success(res.status_message);
    });
    emit("visible-update", false);
  } else if (videoForm.url != "" || videoForm.coverUrl != "") {
    message.error("请等待视频上传完成后重试");
  } else {
    message.error("请完整填写视频信息");
  }
};

const handleCancelVideo = () => {
  emit("visible-update", false);
};

// 上传前的回调，判断视频大小等
const beforeUpload = (data: {
  file: UploadFileInfo;
  fileList: UploadFileInfo[];
}) => {
  let fileSize;
  if (data.file.file) {
    fileSize = data.file.file.size / 1024 / 1024 < 100; //控制大小  修改100的值即可
    if (
      [
        "video/mp4",
        "video/ogg",
        "video/flv",
        "video/avi",
        "video/wmv",
        "video/rmvb",
        "video/mov",
      ].indexOf(data.file.file.type) == -1
    ) {
      message.error("请上传正确的视频格式");
      return false;
    }
    if (!fileSize) {
      message.error("视频大小不能超过100MB");
      return false;
    }
  }
  return true;
};

// 上传成功回调, 把url和coverurl添加
const handleFinish = ({
  file,
  event,
}: {
  file: UploadFileInfo;
  event?: ProgressEvent;
}) => {
  if (event && event.target) {
    let res: any = event.target;
    if (res.response) {
      let resJson = JSON.parse(res.response);
      let url = resJson.url;
      let coverUrl = resJson.coverUrl;
      videoForm.url = url;
      videoForm.coverUrl = coverUrl;
    }
  }
  return file;
};

function uploadHeader() {
  return {
    Authorization: token.value,
  };
}

const selectVisible = ref<boolean>(false);
</script>

<template>
  <NModal
    v-model:show="formVisible"
    class="custom-card"
    preset="card"
    title="投稿视频"
    :closable="false"
    size="huge"
    :bordered="false"
    width="30%"
  >
    <NForm :model="videoForm" label-placement="left">
      <NFormItemRow label="描述">
        <NInput
          v-model:value="videoForm.title"
          placeholder="请输入"
          autocomplete="off"
          clearable
          type="textarea"
        />
      </NFormItemRow>
      <NFormItemRow label="上传">
        <NUpload
          ref="fileUploadRef"
          multiple
          directory-dnd
          :action="postBaseURL"
          :headers="uploadHeader"
          @before-upload="beforeUpload"
          @finish="handleFinish"
        >
          <NUploadDragger>
            <div style="margin-bottom: 12px">
              <NIcon size="48" :depth="3">
                <CloudUpload />
              </NIcon>
            </div>
            <NText style="font-size: 16px">
              点击或者拖动文件到该区域来上传
            </NText>
          </NUploadDragger>
        </NUpload>
      </NFormItemRow>
      <NFormItemRow label="分类">
        <NSelect
          v-model:show="selectVisible"
          v-model:value="videoForm.category"
          placeholder="分类"
          :options="classifyList"
        >
          <template #arrow>
            <Transition name="slide-left">
              <Dice v-if="selectVisible" />
              <DiceOutline v-else />
            </Transition>
          </template>
        </NSelect>
      </NFormItemRow>
    </NForm>
    <div class="footer">
      <NButton
        class="post-btn"
        type="primary"
        block
        secondary
        strong
        @click="handlePostVideo"
      >
        发布
      </NButton>
      <NButton
        class="cancle-btn"
        block
        secondary
        strong
        @click="handleCancelVideo"
      >
        取消
      </NButton>
    </div>
  </NModal>
</template>

<style scoped>
.el-select {
  width: 100%;
}

.footer {
  display: flex;
  width: 100%;
  .cancle-btn {
    width: 45%;
  }

  .post-btn {
    width: 45%;
    margin-right: 10%;
  }
}
</style>
