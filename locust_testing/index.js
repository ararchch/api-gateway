const express = require("express");
const app = express();
const PORT = 3000;

//Basic load testing request 
app.get("/hello", (req, res) => {
  res.send("Hello");
});

app.get("/slow", (req, res) => {
  setTimeout(() => {
    res.send("Slow");
  }, 2000);
});

//Authentication request 
//Simple authentication procedure where server sends client the token 
app.get("/login", (req, res) => {
  res.json({
    //secret token is sent back to the client for authentication
    token: "secret-token",
  });
});

//This func does the authentication and checks for correct token
const loginMiddleware = (req, res, next) => {
  const token = req.headers.authorization;
  if (token !== "secret-token") {
    res.status(403).send("Not logged in");
  } else {
    next();
  }
};

//To access /secret page, include headers; 'Authorization': 'secret-token' 
app.get("/secret", loginMiddleware, (req, res) => {
  res.send("Welcome to the secret page");
});

//Listening to port 3000
app.listen(PORT, () => {
  console.log("Server is up");
});
