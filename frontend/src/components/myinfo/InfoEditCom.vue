<!--
 * @Author: Xu Ning
 * @Date: 2023-10-29 17:10:06
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-29 20:14:00
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\components\myinfo\InfoEditCom.vue
-->
<template>
  <!-- Form -->
  <ElDialog
    v-model="visibel"
    width="30vw"
    title="修改资料"
    :before-close="closeFunc"
  >
    <ElForm :model="form">
      <ElFormItem label="头像" :label-width="formLabelWidth">
        <ElUpload
          class="avatar-uploader"
          action="#"
          :show-file-list="false"
          :before-upload="beforeAvatarUpload"
        >
          <img v-if="form.avatar" :src="userInfo.avatar" class="avatar" />
          <ElIcon v-else class="avatar-uploader-icon"><Plus /></ElIcon>
        </ElUpload>
      </ElFormItem>
      <ElFormItem label="昵称" :label-width="formLabelWidth">
        <ElInput v-model="form.name" />
      </ElFormItem>
      <ElFormItem label="手机号" :label-width="formLabelWidth">
        <ElInput v-model="form.phoneNum" />
      </ElFormItem>

      <ElFormItem label="性别" :label-width="formLabelWidth">
        <ElSelect v-model="form.gender" :placeholder="userInfo.gender">
          <ElOption label="女" value="0" />
          <ElOption label="男" value="1" />
        </ElSelect>
      </ElFormItem>

      <ElFormItem label="自我介绍" :label-width="formLabelWidth">
        <ElInput v-model="form.signature" />
      </ElFormItem>
    </ElForm>
    <template #footer>
      <span class="dialog-footer">
        <ElButton @click="emit('visible-update', false)">取消</ElButton>
        <ElButton type="primary" @click="updateUserInfo"> 修改 </ElButton>
      </span>
    </template>
  </ElDialog>
</template>

<script lang="ts" setup>
import { reactive, computed, onMounted } from "vue";
import { ElMessage } from "element-plus";
import { UploadProps } from "element-plus";
import { UserType } from "@/apis/interface";
import { userStore } from "@/stores/user";

interface propsType {
  isVisible: boolean;
  userInfo: UserType;
}

const props = defineProps<propsType>();
const emit = defineEmits(["visible-update"]);

const visibel = computed(() => props.isVisible);
const userInfo = computed(() => props.userInfo);
const formLabelWidth = "20%";

const form = reactive({
  name: "",
  phoneNum: "",
  gender: 0,
  signature: "",
  avatar: "",
});

onMounted(() => {
  form.name = userInfo.value.name;
  form.phoneNum = userStore().phoneNum;
  form.gender = userInfo.value.gender;
  form.signature = userInfo.value.signature;
  form.avatar = userInfo.value.avatar;
});

const beforeAvatarUpload: UploadProps["beforeUpload"] = (rawFile) => {
  if (rawFile.type !== "image/jpeg") {
    ElMessage.error("Avatar picture must be JPG format!");
    return false;
  } else if (rawFile.size / 1024 / 1024 > 2) {
    ElMessage.error("Avatar picture size can not exceed 2MB!");
    return false;
  }
  console.log("rawFile", rawFile);
};

const closeFunc = () => {
  emit("visible-update", false);
};

const updateUserInfo = () => {};
</script>

<style scoped>
.el-input {
  width: 90%;
}
.el-button--text {
  margin-right: 15px;
}
.el-select {
  width: 90%;
}
.dialog-footer button:first-child {
  margin-right: 10px;
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
