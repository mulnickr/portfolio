const express = require("express");
const http = require("http");
const path = require("path");

const port = process.env.PORT || 4200;
const host = process.env.HOST || "0.0.0.0";

const app = express();

app.use(express.static(path.join(__dirname, "dist/website/browser")));
app.get("/", function (req, res, next) {
  res.redirect("/");
});

app.listen(port, host, () =>
  console.log(`Running Portfolio on ${host}:${port}`)
);
