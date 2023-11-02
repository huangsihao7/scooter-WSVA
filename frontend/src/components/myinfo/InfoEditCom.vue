<!--
 * @Author: Xu Ning
 * @Date: 2023-10-29 17:10:06
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-02 19:06:19
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\components\myinfo\InfoEditCom.vue
-->
<template>
  <NModal
    v-model:show="visibel"
    width="30vw"
    preset="dialog"
    title="编辑资料"
    :show-icon="false"
    positive-text="确认"
    negative-text="取消"
    :closable="false"
    @positive-click="submitCallback"
    @negative-click="cancelCallback"
  >
    <NForm
      ref="formRef"
      class="form-area"
      :model="formValue"
      label-placement="left"
      label-width="auto"
    >
      <NFormItem label="头像">
        <NUpload
          :action="uploadApi"
          :default-file-list="[
            {
              id: '1',
              name: 'avatar',
              status: 'finished',
              url: formValue.avatar,
            },
          ]"
          list-type="image-card"
          :max="1"
          @before-upload="beforeAvatarUpload"
          @finish="handleAvatarFinish"
        >
          点击上传
        </NUpload>
      </NFormItem>
      <NFormItem label="背景">
        <NUpload
          :action="uploadApi"
          :default-file-list="[
            {
              id: '1',
              name: 'background',
              status: 'finished',
              url: formValue.background,
            },
          ]"
          list-type="image-card"
          :max="1"
          @before-upload="beforeBgUpload"
          @finish="handleBgFinish"
        >
          点击上传
        </NUpload>
      </NFormItem>
      <NFormItem label="昵称">
        <NInput v-model:value="formValue.name" placeholder="请输入姓名" />
      </NFormItem>

      <NFormItem label="性别">
        <NSelect
          v-model:value="formValue.gender"
          placeholder="请输入性别"
          :options="genderOptions"
        />
      </NFormItem>
      <NFormItem label="简介">
        <NInput v-model:value="formValue.signature" placeholder="请输入简介" />
      </NFormItem>
    </NForm>
  </NModal>
</template>

<script lang="ts" setup>
import { computed, onMounted, ref } from "vue";
import {
  FormInst,
  NForm,
  NFormItem,
  NInput,
  NModal,
  NSelect,
  NUpload,
  UploadFileInfo,
  useMessage,
} from "naive-ui";
import { UserType } from "@/apis/interface";
import { userStore } from "@/stores/user";
import { baseURL } from "@/axios";
import { updateUserInfo } from "@/apis/user";

const uploadApi = baseURL + "/user/uploadImg";
const message = useMessage();

interface propsType {
  isVisible: boolean;
  userInfo: UserType;
}

const props = defineProps<propsType>();
const emit = defineEmits(["visible-update"]);
const formRef = ref<FormInst | null>(null);
const visibel = computed(() => props.isVisible);
const formValue = ref({
  name: "",
  gender: 0,
  signature: "",
  background: "",
  avatar: "",
});
const genderOptions = [
  {
    label: "男",
    value: 1,
  },
  {
    label: "女",
    value: 0,
  },
];

onMounted(() => {
  formValue.value.name = userStore().name;
  formValue.value.gender = userStore().gender;
  formValue.value.signature = userStore().signature;
  formValue.value.background = userStore().background_image;
  formValue.value.avatar = userStore().avatar;
});

// 上传限制
const beforeAvatarUpload = (data: {
  file: UploadFileInfo;
  fileList: UploadFileInfo[];
}) => {
  if (data.file.file?.type !== "image/jpeg") {
    message.error("Avatar picture must be JPG format!");
    return false;
  } else if (data.file.file.size / 1024 / 1024 > 2) {
    message.error("Avatar picture size can not exceed 2MB!");
    return false;
  }
};

// 上传限制
const beforeBgUpload = (data: {
  file: UploadFileInfo;
  fileList: UploadFileInfo[];
}) => {
  if (data.file.file?.type !== "image/jpeg") {
    message.error("Avatar picture must be JPG format!");
    return false;
  } else if (data.file.file.size / 1024 / 1024 > 10) {
    message.error("Avatar picture size can not exceed 2MB!");
    return false;
  }
};

// 获取背景url
const handleBgFinish = ({
  file,
  event,
}: {
  file: UploadFileInfo;
  event?: ProgressEvent;
}) => {
  console.log(file);
  const response = (event?.target as XMLHttpRequest)?.response;
  let responseJson = JSON.parse(response);
  formValue.value.background = responseJson.url;
};

// 获取头像url
const handleAvatarFinish = ({
  file,
  event,
}: {
  file: UploadFileInfo;
  event?: ProgressEvent;
}) => {
  console.log(file);
  const response = (event?.target as XMLHttpRequest)?.response;
  let responseJson = JSON.parse(response);
  formValue.value.avatar = responseJson.url;
};

// 模态框取消回调
const cancelCallback = () => {
  emit("visible-update", false);
};

// 修改资料
const submitCallback = () => {
  let userInfo = formValue.value;
  updateUserInfo(
    userInfo.name,
    userInfo.gender,
    userInfo.avatar,
    userInfo.signature,
    userInfo.background,
  ).then((res: any) => {
    if (res.status_code == 200) {
      userStore().name = userInfo.name;
      userStore().gender = userInfo.gender;
      userStore().avatar = userInfo.avatar;
      userStore().signature = userInfo.signature;
      userStore().background_image = userInfo.background;
      message.success("修改成功");
      cancelCallback();
      window.location.reload();
    } else {
      message.error("修改失败");
    }
  });
};
</script>

<style scoped>
.form-area {
  margin-top: 30px;
}


</style>
