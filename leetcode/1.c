/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
#include <stdlib.h>
int* twoSum(int* nums, int numsSize, int target) {
    int* ret = (int*)malloc(sizeof(int) * 2);
    for(int i=0;i<numsSize;i++) {
        int a = nums[i];
        for(int j=i+1;j<numsSize;j++) {
            int b = nums[j];
            if(a + b == target) {
                ret[0] = i;
                ret[1] = j;
                return ret;
            }
        }
    }
    return ret;
}
