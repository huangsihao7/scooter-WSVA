/*
 * @Author: Xu Ning
 * @Date: 2023-11-04 15:39:28
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-04 15:47:11
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\test\sum.test.ts
 */
/*
 * @Author: Xu Ning
 * @Date: 2023-11-04 15:39:28
 * @LastEditors: Xu Ning
 * @LastEditTime: 2023-11-04 15:40:34
 * @Description: 
 * @FilePath: \scooter-WSVA\frontend\src\test\sum.test.ts
 */
import {describe, expect, test} from '@jest/globals';
import { sumsa } from '@/test/sum';

describe('sum module', () => {
  test('adds 1 + 2 to equal 3', () => {
    expect(sumsa(1, 2)).toBe(3);
  });
});