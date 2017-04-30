## apache httpd 실습

 cf) mac os 환경
 - configure, make and make install?
   - http://www.codecoffee.com/tipsforlinux/articles/27.html
 - 현재 문서는 2.4 버전 기준으로 작성
   - http://httpd.apache.org/docs/2.4/en/install.html
 - requirement
   - APR, APR-Util
   - PCRE(Perl-Compatible Regular Expressions Library)
   - gcc
   - Accurate time keeping: ntpdate, xntpd(Network Time Protocol, NTP)
   - Perl 5[Optional]

### apr, apr-util 설치

 - 다운로드
   - http://apr.apache.org/
 - 설치
   - 압축풀기
     - tar xvf apr-x.x.x.tar.gz
     - tar xzvf apr-util-x.x.x.tar.gz
   - apr-x.x.x 디렉터리에서:
     - ./configure
     - make
     - sudo make install
   - apr-util-x.x.x 디렉터리에서:
     - ./configure --with-apr=/usr/local/apr
     - make
     - sudo make install

### pcre 설치

 - 다운로드
   - https://ftp.pcre.org/pub/pcre/
 - 설치
   - 압축풀기
     - tar xvf pcre2-x.x.tar.gz
   - pcre2-x.x 디렉터리에서:
     - ./configure
     - make
     - sudo make install

### apache httpd 설치

  - 다운로드
    - http://httpd.apache.org/download.cgi
  - 설치
    - 압축풀기
      - tar xvf httpd-x.x.x.tar.gz
    - httpd-x.x.x 디렉터리에서:
      - ./configure
      - make
      - sudo make install
    - 기본값으로 /usr/local/apache2 에 설치됨.
    - /usr/local/apache2/conf/httpd.conf 파일을 수정하여 서버 설정 변경 가능
      - [주소 및 포트 설정](http://httpd.apache.org/docs/2.4/en/bind.html)
      - http://httpd.apache.org/docs/2.4/en/configuring.html
    - 실행 테스트
      - sudo apachectl -k start
        - 브라우저에서 localhost 를 입력하여 `It works!` 문구 확인
      - sudo apachectl -k stop


