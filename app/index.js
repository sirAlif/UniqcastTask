var natsHandler = require("./natsHandler")
const express = require("express")
var app = express()

app.get("/", function (request, response) {
  response.send("Hello World!")
})

const server = app.listen(7009, () => {
  console.log("Started application on port 7009");
  natsHandler.provideNatsConnection();
});