const express = require("express");
const app = express();
const port = 1234;

app.get("/", (req, res) => {
  setTimeout(() => res.send("Hello World!"), 5000);
});

app.listen(port, () => console.log(`Example app listening on port ${port}!`));
