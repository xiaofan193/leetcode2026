// 1213 3个有序数组的交集
// intersection_of_three_sorted_arrays
//
// https://leetcode.cn/problems/intersection-of-three-sorted-arrays/description/
//
// 给定3个严格递增的排列的整数数组 arr1 arr2 arr3,返回一个仅由这3个数组中同时出现整数构成的有序数组
// 例如： arr1 = [1,2,3,4,5], arr2 = [1,2,5,7,9],arr3 = [1,3,4,5,8]
// 答案: [1,5]
function arraysIntersection(arr1: number[], arr2: number[], arr3: number[]): number[] {
    const res: number[] = [];

    // 三指针同向扫描：三个指针都指向各自数组当前最小的未比较元素
    let i = 0;
    let j = 0;
    let k = 0;

    while (i < arr1.length && j < arr2.length && k < arr3.length) {
        if (arr1[i] === arr2[j] && arr2[j] === arr3[k]) {
            // 数组严格递增，相等元素唯一，无需去重
            res.push(arr1[i]);
            i++;
            j++;
            k++;
        } else {
            // 把三者中最小的指针右移，向更大的值靠拢
            let minx = arr1[i];
            if (arr2[j] < minx) {
                minx = arr2[j];
            }
            if (arr3[k] < minx) {
                minx = arr3[k];
            }

            // 可能多个指针同时命中最小值，都要前进
            if (arr1[i] === minx) {
                i++;
            }

            if (arr2[j] === minx) {
                j++;
            }

            if (arr3[k] === minx) {
                k++;
            }
        }
    }

    return res;
}
