const express = require("express");
const app = express();
const port = 1236;

function fibonacci(num) {
  if (num == 0) {
    return 0;
  }
  if (num == 1) {
    return 1;
  }
  return fibonacci(num - 1) + fibonacci(num - 2);
}

app.get("/", (req, res) => {
  var tmp = fibonacci(45);
  res.send(tmp.toString());
});

app.listen(port, () => console.log(`Example app listening on port ${port}!`));
