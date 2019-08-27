#include <stdlib.h>
double findMedianSortedArrays(int* nums1, int nums1Size, int* nums2, int nums2Size) {
    int idx1 = 0, idx2 = 0;
    int cur = 0;
    int* arr = (int*)malloc(sizeof(int) * (nums1Size + nums2Size));
    while(idx1 < nums1Size && idx2 < nums2Size) {
        if(nums1[idx1] < nums2[idx2]) {
            arr[cur] = nums1[idx1];
            cur++;
            idx1++;
        } else {
            arr[cur] = nums2[idx2];
            cur++;
            idx2++;
        }
    }
    while(idx1 < nums1Size) {
        arr[cur] = nums1[idx1];
        cur++;
        idx1++;
    }
    while(idx2 < nums2Size) {
        arr[cur] = nums2[idx2];
        cur++;
        idx2++;
    }
    int sum = nums1Size + nums2Size;
    if (sum % 2 == 0) {
        return (arr[sum / 2] + arr[sum / 2 - 1]) / 2.0;
    } else {
        return arr[sum / 2];
    }
}
