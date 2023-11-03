/*
 * @Author: Xu Ning
 * @Date: 2023-10-31 19:56:09
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-03 21:56:18
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
      colorTarget:'red'
  },
  Card:{
    borderRadius:'25px',
    color:'#eeeeee70'
  }
};

export default themeOverrides;
