/**
 * Return an array of size *returnSize.
 * Note: The returned array must be malloced, assume caller calls free().
 */
#include <stdlib.h>
#include <string.h>
const int TRUE = 1;
const int FALSE = 0;


int isDigit(char c) {
    if ('0' <= c && c <= '9') {
        return TRUE;
    }
    return FALSE;
}
// 0: letter-log
// 1: digit-log
int classifyLog(char* log) {
    int i = 0;
    while(log[i] != ' ') {
        i++;
    }
    if (isDigit(log[i + 1])) {
        return 1;
    } else {
        return 0;
    }
}

// return 1: a <= b
// return 0: a > b
int compareIdentifier(char* a, char* b) {
    int idx = 0;
    while(a[idx] != ' ' && b[idx] != ' ') {
        if (a[idx] < b[idx]) {
            return 1;
        }
        if (a[idx] > b[idx]) {
            return 0;
        }
        idx++;
    }
    if (a[idx] == ' ' && b[idx] != ' ') {
        return 1;
    }
    if (a[idx] == ' ') {
        return 1;
    } else {
        return 0;
    }
}



// return 1: a < b
// return 0: a == b
// return -1: a > b
int compareStr(char* a, char* b) {
    int aLen = strlen(a);
    int bLen = strlen(b);
    int smallest = aLen < bLen ? aLen : bLen;
    for(int i=0;i<smallest;i++) {
        if(a[i] < b[i]) {
            return 1;
        }
        if(a[i] > b[i]) {
            return -1;
        }
    }
    if(aLen == bLen) {
        return 0;
    }
    if(aLen < bLen) {
        return 1;
    } else {
        return -1;
    }
}

// return 1: a < b
// return -1: a > b
int compareLetterLog(char* a, char* b) {
    int nonIDIdxA = 0;
    int nonIDIdxB = 0;
    for(int i=0;i<strlen(a);i++) {
        if(a[i] == ' ') {
            nonIDIdxA = i + 1;
            break;
        }
    }
    for(int i=0;i<strlen(b);i++) {
        if(b[i] == ' ') {
            nonIDIdxB = i + 1;
            break;
        }
    }
    int val = compareStr(a + nonIDIdxA, b + nonIDIdxB);
    if(val == 1) {
        return 1;
    }
    if(val == -1) {
        return -1;
    }
    val = compareIdentifier(a, b);
    if (val == 1) {
        return 1;
    } else {
        return -1;
    }
}


void sortLetterLogs(char** logs, int len) {
    for(int i=0;i<len;i++) {
        for(int j=i+1;j<len;j++) {
            if (compareLetterLog(logs[i], logs[j]) != 1) {
                char* tmp = logs[i];
                logs[i] = logs[j];
                logs[j] = tmp;
            }
        }
    }
}

char** reorderLogFiles(char** logs, int logsSize, int* returnSize) {
    char* digitLogs[100];
    char* letterLogs[100];
    int digitLogsLen = 0;
    int letterLogsLen = 0;
    for(int i=0;i<logsSize;i++) {
        if (classifyLog(logs[i]) == 1) {
            digitLogs[digitLogsLen++] = logs[i];
        } else {
            letterLogs[letterLogsLen++] = logs[i];
        }
    }
    sortLetterLogs(letterLogs, letterLogsLen);
    char** ret = (char**)malloc(sizeof(char*) * logsSize);
    *returnSize = logsSize;
    int retIdx = 0;
    for(int i=0;i<letterLogsLen;i++) {
        ret[retIdx++] = letterLogs[i];
    }
    for(int i=0;i<digitLogsLen;i++) {
        ret[retIdx++] = digitLogs[i];
    }
    return ret;
}

