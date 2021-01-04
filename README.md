# Finstagram

#### A full stack minimal Instagram clone:
##### - Frontend / User-interface written in Vue.js
##### - Backend / Server written in Golang

The app allows registration, login as well as certain CRUD capabilities. 
Quasar framework is used for the frontend components. 
Fiber framework is used on the backend. 
The server connects to a MongoDB atlas cloud database.

## Install the dependencies
```bash
cd finstagram/
npm install
```
### Start the app in development mode
```bash
cd finstagram/
quasar dev
```
Update MongoDB connection details in server/handlers/database.go+70
### Start the Server
```bash
cd server/
go run server.go
```

![caption](https://media.giphy.com/media/8biE3Ur80n1piWVa0J/giphy.gif)
![caption](https://media.giphy.com/media/4Qx4f5Vh6BOPdzEpEQ/giphy.gif)

<br />

![caption](https://media.giphy.com/media/YhmXCpu8FfLm9GA7KO/giphy.gif)

<br />


![caption](https://media.giphy.com/media/fm4pRh4doeNbgBrfFQ/giphy.gif)

Todo -

- [ ] Encrypt password
