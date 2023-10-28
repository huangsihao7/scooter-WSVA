import { service } from '@/axios'

export function getUserInfo(uid: number) {
  return service({
    url: '/user/userinfo',
    method: 'post',
    data: { uid }
  })
}

