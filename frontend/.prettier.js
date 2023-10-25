/*
 * @Author: Xu Ning
 * @Date: 2023-10-22 19:46:36
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-10-22 19:46:41
 * @Description: 
 * @FilePath: \myMap\.prettier.js
 */
module.exports = {
    tabWidth: 2,
    jsxSingleQuote: true,
    jsxBracketSameLine: true,
    printWidth: 100,
    singleQuote: true,
    semi: false,
    overrides: [
      {
        files: '*.json',
        options: {
          printWidth: 200,
        },
      },
    ],
    arrowParens: 'always',
  }
  