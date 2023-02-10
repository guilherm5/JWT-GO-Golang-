# Autentication Golang + JWT 

<h1> Install and run Project </h1>
<p>git clone https://github.com/guilherm5/jwt-golang.git</p>
<p>go mod tidy</p>
<p>go run main.go</p>
<p>you can test this project by postman or by API Client you want. When running the project with go run main.go, routes will be opened to port 5000</p>

<h2>How to use</h2>

in your preferred Api client, access the routes, using the following paths:
<br> <br>

<p>method GET: localhost:5000/user - to view users</p>
<h3>Example post</h3>
<p>method POST: localhost:5000/user - to register users</p>
<img src="./img/post-user.png" alt="how to post user">

<h3>Example PUT</h3>
<p>method PUT: localhost:5000/user - to update user</p>
<img src="./img/put-user.png" alt="how to put user">

<h3>Example DELETE</h3>
<p>method DELETE: localhost:5000/user - to delete users</p>
<img src="./img/delete-user.png" alt="how to delete user">




<h2>Get JWT token</h2>

<p>first you must register a user (and keep your password, because after entering it, your password will be encrypted)

you must provide the registered email and password using the body/json

this will generate you a jwt token</p>

<br> <br>
<h3>example of how to get token</h3>
<p>method post: localhost:5000/postAuth</p>
<img src="./img/jwt-user.png" alt="how to jwt user">

<br> <br>

<h2>how to access route protected by Middleware</h2>

<p>method get: localhost:5000/getUser</p>
<img src="./img/get-user-middleware.png" alt="how to get user">
<h2> don't forget to update the credentials for your local bank, you can do this by changing the .env file</h2>


<h1>Technologies used:</h2>
<h4>Golang</h4>
<h4>Gin</h4>
<h4>Libpq</h4>
<h4>Postgresql</h4>
<h4>JWT</h4>
