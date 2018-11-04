#include <string.h>
#include <stdlib.h>
bool isPalindrome(char* s, int start, int end) {
    for(int i = start, j = end;i <= j ; i++, j--) {
        if (s[i] != s[j]) {
            return false;
        }
    }
    return true;
}

char* longestPalindrome(char* s) {
    int len = strlen(s);
    int maxLen = 1;
    int startIdx = 0;
    int endIdx = 0;
    for(int i=0;i < len - 1; i++) {
        for(int j=i+1; j < len; j++) {
            if (isPalindrome(s, i, j)) {
                int len = j - i + 1;
                if (maxLen < len) {
                    maxLen = len;
                    startIdx = i;
                    endIdx = j;
                }
            }
        }
    }
    char* ret = (char*)malloc(sizeof(char) * (endIdx - startIdx + 2));
    for(int i=0, j=startIdx; j<=endIdx;i++, j++) {
        ret[i] = s[j];
    }
    ret[endIdx - startIdx + 1] = 0;
    return ret;
}
