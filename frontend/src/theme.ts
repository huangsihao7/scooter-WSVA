/*
 * @Author: Xu Ning
 * @Date: 2023-10-31 19:56:09
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-06 14:19:32
 * @Description:
 * @FilePath: \scooter-WSVA\frontend\src\theme.ts
 */
import type { GlobalThemeOverrides } from "naive-ui";
const themeOverrides: GlobalThemeOverrides = {
  common: {
    primaryColor: "#409EFF",
    primaryColorHover: "#409EFF",
    primaryColorPressed: "#00A6E0",
    // primaryColorSuppl: '#4098FC'
  },
  Button: {
    // textColor: 'red',
  },
  Card: {
    borderRadius: "25px",
    color: "#fff9",
  },
};

export default themeOverrides;
