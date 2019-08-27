# ambari design

- [ ] ambari design pdf 다 읽기
  - dependency tracker 가 어느 정도까지의 dependency 를 책임져 줄까? 설치 말고 실행 순서도 가능할까?
- [ ] 이전 ambari server 코드에서 agent 들이 대량 죽는 원인을 알아낼 수 있을까?
  - 부팅과정을 빠르게 하는 방법?
  - ambari server 재부팅 과정 알아보기

https://cwiki.apache.org/confluence/display/AMBARI/Ambari
https://cwiki.apache.org/confluence/display/AMBARI/Ambari+Design


## goal

- ambari 를 이해하며 코드를 분석하는 법, 아키텍처를 바라보는 능력을 기르자.
  - e.g. ambari design pdf 에서는 `pluggable components` 라는 개념이 있다. 나는 pluggable 하게 구조를 만드는 것에 대한 개념이 없는데, 실제 어떻게 이를 구현해 냈는가를 찾아보자.





## Agent registration flow

https://www.slideshare.net/hortonworks/ambari-agentregistrationflow-17041261

1. ambari.repo 를 배포한다.
1. setupAgent.py 를 배포한다.
1. setupAgent.py 를 실행한다.
1. epel-release 를 설치한다. (yum)
1. yum 으로 ambari-agent package 를 설치한다.
1. /etc/ambari-agent/ambari-agent.ini 를 변경한다. hostname 을 ambari server hostname 으로 변경한다.
1. ambari-agent 를 실행한다.
1. registration 절차를 실행한다.

### agent registration 절차

(아래 사항을 보면 왜 8441, 8440 포트를 열어두었어야 하는지 이해 된다.)
- 8441: handshake port
- 8440: registration port

- [ ] certificate 을 주고받는 건 ssl 인가? 서명을 하는 과정 이런것들은 다 뭘까?
  - 8441 포트에서 주고받는 과정은 무엇일까?


- [ ] 코드로 따라가보기
  - 왜 이렇게 했을까?
  - cert 는 무엇일까?

아래 설명이 약간 반대되는 것 같다. 실제 코드에선 다음과 같은 주석이 있다.

```
# AmbariServer.java
// Agents download cert on on-way connector but always communicate on
// two-way connector for server-agent communication
ServerConnector agentOneWayConnector = createSelectChannelConnectorForAgent(serverForAgent, configs.getOneWayAuthPort(), false, agentAcceptors);
ServerConnector agentTwoWayConnector = createSelectChannelConnectorForAgent(serverForAgent, configs.getTwoWayAuthPort(), configs.isTwoWaySsl(), agentAcceptors);
```

https://www.slideshare.net/hortonworks/ambari-agentregistrationflow-17041261 에서 
8440 포트와 8441 포트의 역할이 반대로 기술된 것으로 보인다.

c3s 에서는 one way ssl 이라, 아래과정을 거치치않고 바로 8441 포트에서 stomp 연결을 맺는다.
/register 를 destination 으로 등록요청한다. (HeartbeatThread.py)


1. 8441 포트로 핸드쉐이크 실행
1. server -> agent 로 server certificate 를 보내준다.
1. server <- agent 로 agent certificate 요청
   - server 에 agent certificate 를 서명해달라고 요청한다.
1. server 에서 agent certificate 를 서명
   - server 가 agent certificate 를 비밀번호로 서명한다.
1. server -> agent 로 agent certificate 주고 연결을 끊는다.
1. server <- agent 로 8440 registration 포트로 연결
1. server <- agent 로 agent certificate 를 이용하여 2way auth 수행
1. agent 에서 Host Script 를 이용하여 FQDN 알아오기?
   - agent 는 FQDN 을 알아와야 한다. 만일 여러 개의 hostname 을 갖고 있을 때에는 host script 를 이용하여 registration 을 위해 쓸 hostname 을 정할 수 있도록 한다.
   - https://docs.hortonworks.com/HDPDocuments/HDP1/HDP-1.2.1/bk_using_Ambari_book/content/ambari-chap7a.html
   - `hostname_script` in `[agent]` section on /etc/ambari-agent/conf/ambari-agent.ini
1. server <- agent 로 FQDN 을 이용하여 호스트 등록
1. server 는 호스트 등록을 마친다.
   - server 는 ambari db 에 호스트를 추가하면서 등록을 마친다.
1. server <- agent 로 Heartbeating 시작

## stacks and services

- [ ] https://cwiki.apache.org/confluence/display/AMBARI/Stacks+and+Services 읽기

## automated kerberization

- [ ] https://cwiki.apache.org/confluence/display/AMBARI/Automated+Kerberizaton 읽기

