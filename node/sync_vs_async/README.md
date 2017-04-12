# sync vs. async
 - 참고: https://nodejs.org/en/docs/guides/blocking-vs-non-blocking/

## blocking
 해당 operation 이 끝날 때까지 기다려야 한다.
 - Native module
 - Node.js std. lib. 중, Sync 로 끝나는 method 들

```js
const fs = require('fs');
const data = fs.readFileSync('file.md');
console.log(data);
console.log('read finished');
```

## non-blocking
 - Sync 로 끝나지 않는 Node.js std. lib. 

```js
const fs = require('fs');
fs.readFile('file.md', (err, data) => {
        if (err) throw err;
        console.log(data);
});
console.log('read not finished');
```

## Blocking 과 Non-blocking 을 섞어 쓰지말 것.

```js
const fs = require('fs');
fs.readFile('/file.md', (err, data) => {
  if (err) throw err;
  console.log(data);
});
fs.unlinkSync('/file.md');
```

위 코드 처럼 쓰지말 것, 아래와 같이 수정.

```js
const fs = require('fs');
fs.readFile('/file.md', (err, data) => {
  if (err) throw err;
  console.log(data);
  fs.unlink('/file.md', (err) => {
    if (err) throw err;
  });
});
```
