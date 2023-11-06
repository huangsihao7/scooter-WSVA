import { service } from "@/axios";

export function launchLive() {
  return service({
    url: "/live/start",
    method: "get",
  });
}

export function getLiveList() {
  return service({
    url: "/live/list",
    method: "get",
  });
}
