import { defineStore } from "pinia";
export const liveStore = defineStore({
  id: "live",
  state: () => ({
    live_play_url: "",
    live_cover_url: "",
    live_id: -1,
  }),
  persist: true,
});
