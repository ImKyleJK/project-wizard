name | description | languages
--- | --- | ---
Express | Simple Express.JS Server (No Rendering) | javascript | ImKyleJK

# index.js
```
const express = require('express')
const app = express()
const port = 3000

app.get('/', (req, res) => {
  return res.send('Hello World!');
})

app.post('/', (req, res) => {
  return res.json({"data":"success"});
})

app.listen((port || 80), () => {
  console.log(`Server running on the following port: ${port}`)
})
```

command name | command | params
--- | --- | ---
Install All Required Packaged | npm install |
Start Express | node index.js |
