<!--
 * @Author: Xu Ning
 * @Date: 2023-10-29 17:10:06
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-29 17:20:18
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\components\myinfo\InfoEditCom.vue
-->
<template>
    <!-- Form -->
    <el-dialog v-model="props.isVisible" title="Shipping address">
      <el-form :model="form">
        <el-form-item label="Promotion name" :label-width="formLabelWidth">
          <el-input v-model="form.name" autocomplete="off" />
        </el-form-item>
        <el-form-item label="Zones" :label-width="formLabelWidth">
          <el-select v-model="form.region" placeholder="Please select a zone">
            <el-option label="Zone No.1" value="shanghai" />
            <el-option label="Zone No.2" value="beijing" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="emit('visible-update',false)">Cancel</el-button>
          <el-button type="primary" @click="updateUserInfo">
            Confirm
          </el-button>
        </span>
      </template>
    </el-dialog>
  </template>
  
  <script lang="ts" setup>
  import { reactive, ref } from 'vue'
  
  interface propsType{
    isVisible : boolean
  }
  
  const props = defineProps<propsType>()
  const emit = defineEmits(['visible-update'])
  const formLabelWidth = '140px'
  
  const form = reactive({
    name: '',
    region: '',
    date1: '',
    date2: '',
    delivery: false,
    type: [],
    resource: '',
    desc: '',
  })

  const updateUserInfo = () =>{
    emit('visible-update',false)
  }
  
  </script>
  <style scoped>
  .el-button--text {
    margin-right: 15px;
  }
  .el-select {
    width: 300px;
  }
  .el-input {
    width: 300px;
  }
  .dialog-footer button:first-child {
    margin-right: 10px;
  }
  </style>
  