/*
 * @Author: Xu Ning
 * @Date: 2023-11-01 23:26:31
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-01 23:31:07
 * @Description:
 * @FilePath: \scooter-WSVA\frontend\src\stores\historySearch.ts
 */
import { defineStore } from "pinia";
interface HistoryStoreState {
  historyData: searchHis[];
}
interface searchHis {
  label: string;
  key: string;
}
export const historyStore = defineStore({
  id: "historySearch",
  state: (): HistoryStoreState => ({
    historyData: [],
  }),
  persist: true,
});
