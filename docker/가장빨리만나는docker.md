
- LXC(LinuX Container) -> libcontainer
 - exec driver:
   - libcontainer: native
   - LXC: lxc
- Control Groups(cgroups)
- Namespace isolation(namespaces)
- userland
- 레이어: 이미지를 통째로 생성하지 않고, 바뀐 부분만 생성한 뒤 부모 이미지를 계속 참조하는 방식
- Union mount, Union File System
 내용이 바뀌었을 때 이미지를 수정하지 않고, 쓰기 이미지를 생성한 뒤 내용을 기록

- 명령어
 - sudo docker attach [???]
 - ctl + p , ctl + q ==> 컨테이너를 정지하지 않고 빠져나오기
 - exec 명령으로 외부에서 컨테이너 실행
  - `sudo docker exec hello echo "Hello World"`
 
