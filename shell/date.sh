#!/bin/bash
# 'date' 명령어 연습

echo "올해가 시작한 뒤로 지금까지 `date +%j` 일이 지났습니다."
# 날짜를 형식화 하려면 포매터 앞에 '+'를 써야 됩니다.
# %j 는 오늘이 연중 몇 번째 날인가를 알려줍니다.

echo "1970/01/01 이후로 지금까지 `date +%s` 초가 지났습니다."
# %s 는 "UNIX 에폭(epoch)" 이 시작한 뒤로 현재까지 몇 초가 지났는 지를 알려줍니다.
#+ 이걸 도대체 어디다 써 먹죠?

prefix=temp
suffix=`eval date +%s` # 'date'의 "+%s" 옵션은 GNU 전용 옵션입니다.
filename=$prefix.$suffix
echo $filename
# "유일한" 임시 파일 이름으로 $$ 를 쓰는  것 보다 더 훌륭합니다.

# 더 많은 형식화 옵션을 보려면 'date' 맨 페이지를 읽어 보세요.

exit 0
