<!--
 * @Author: Xu Ning
 * @Date: 2023-10-27 22:00:03
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-27 22:11:16
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\components\Comment.vue
-->
<script lang="ts" setup>
import { defineComponent, ref } from 'vue'
import { CashOutline as CashIcon } from '@vicons/ionicons5'
import { NIcon, NButton, NSpace,NThing,NAvatar } from 'naive-ui'

interface propsType {
  comment: Object
}
const props = defineProps<propsType>()

const avatar = ref(true)
const header = ref(true)
const headerExtra = ref(true)
const action = ref(true)

</script>

<template>
    <n-thing content-indented>
      <template v-if="avatar" #avatar>
        <n-avatar>
          <n-icon>
            <cash-icon />
          </n-icon>
        </n-avatar>
      </template>
      <template v-if="header" #header>
        货币
      </template>
      <template v-if="headerExtra" #header-extra>
        <n-button circle size="small">
          <template #icon>
            <cash-icon />
          </template>
        </n-button>
      </template>
      货币是为了提高交易效率而用于交换的中介商品。货币有多种形式，如贝壳粮食等自然物、金属纸张等加工品、银行卡信用卡等磁条卡、移动支付加密货币等APP。
      <template v-if="action" #action>
        <n-space>
          <n-button size="small">
            <template #icon>
              <n-icon>
                <cash-icon />
              </n-icon>
            </template>
            1 块钱
          </n-button>
          <n-button size="small">
            <template #icon>
              <n-icon>
                <cash-icon />
              </n-icon>
            </template>
            10 块钱
          </n-button>
          <n-button size="small">
            <template #icon>
              <n-icon>
                <cash-icon />
              </n-icon>
            </template>
            100 块钱
          </n-button>
        </n-space>
      </template>
    </n-thing>
  </template>
  
