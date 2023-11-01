<!--
 * @Author: Xu Ning
 * @Date: 2023-10-29 17:10:06
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-01 11:30:40
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
          action="https://www.mocky.io/v2/5e4bafc63100007100d8b70f"
          :default-file-list="[
            {
              id: '1',
              name: '头像',
              status: 'finished',
              url: formValue.avatar,
            },
          ]"
          list-type="image-card"
          :max="1"
          @before-upload="beforeAvatarUpload"
        >
          点击上传
        </NUpload>
      </NFormItem>
      <NFormItem label="昵称">
        <NInput v-model:value="formValue.name" placeholder="请输入姓名" />
      </NFormItem>
      <NFormItem label="手机号">
        <NInput v-model:value="formValue.phoneNum" placeholder="请输入手机号" />
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
  phoneNum: "",
  gender: 0,
  signature: "",
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
  formValue.value.phoneNum = userStore().phoneNum;
  formValue.value.gender = userStore().gender;
  formValue.value.signature = userStore().signature;
  formValue.value.avatar = userStore().avatar;
  console.log(formValue.value);
});

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

const cancelCallback = () => {
  emit("visible-update", false);
};

const submitCallback = () => {};
</script>

<style scoped>
.dialog-footer button:first-child {
  margin-right: 10px;
}

.form-area {
  margin-top: 30px;
}

.avatar-uploader .avatar {
  width: 178px;
  height: 178px;
  display: block;
}

.avatar-uploader .el-upload {
  border: 1px dashed var(--el-border-color);
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: var(--el-transition-duration-fast);
}

.avatar-uploader .el-upload:hover {
  border-color: var(--el-color-primary);
}

.el-icon.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 100px;
  height: 10vh;
  border: 2px dotted #b7daff;
  border-radius: 25px;
  text-align: center;
}
</style>
