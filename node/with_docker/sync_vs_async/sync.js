const fs = require('fs');
const data = fs.readFileSync('file.md');
console.log(data);
console.log('read finished');
