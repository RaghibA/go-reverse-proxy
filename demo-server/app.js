const express = require('express')

const app = express()
const port = 3012

app.get('/', (req, res) => {
  res.send({
    message: 'Hello!'
  })
})

app.listen(port, () => {
  console.log(`Demo server running on port ${port}`)
})