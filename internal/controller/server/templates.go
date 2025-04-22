package server

import "html/template"

var tmpl = template.Must(template.New("tmpl").Parse(`
{{define "index"}}
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Home</title>
	<style>
		.button-form {
			display: inline-block;
			margin: 10px;
		}
		button {
			padding: 10px 20px;
			font-size: 16px;
			cursor: pointer;
		}
	</style>
</head>
<body>
	<h1>Welcome</h1>
	<div class="button-form">
		<form action="/registration" method="get">
			<button type="submit">Registration</button>
		</form>
	</div>
	<div class="button-form">
		<form action="/login" method="get">
			<button type="submit">Login</button>
		</form>
	</div>
</body>
</html>
{{end}}

{{define "registration"}}
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Registration</title>
	<style>
		form {
			max-width: 400px;
			margin: auto;
		}
		label {
			display: block;
			margin-top: 10px;
		}
		input {
			width: 100%;
			padding: 8px;
			margin-top: 5px;
		}
	</style>
</head>
<body>
	<h1>Registration</h1>
	<form method="post" action="/registration">
		<label for="username">Username:</label>
		<input type="text" id="username" name="username" required>
		
		<label for="password">Password:</label>
		<input type="password" id="password" name="password" required>
		
		<input type="submit" value="Register" style="margin-top:15px;">
	</form>
</body>
</html>
{{end}}


{{define "registrationSuccess"}}
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Registration Successful</title>
  <!-- after 3 seconds, go to /login -->
  <meta http-equiv="refresh" content="3;url=/login">
  <style>
    body { font-family: sans-serif; text-align: center; padding-top: 40px; }
    .message { font-size: 1.2em; }
    a { display: block; margin-top: 20px; }
  </style>
</head>
<body>
  <h1>Registration Complete</h1>
  <p class="message">User <strong>{{.Username}}</strong> registered successfully.</p>
  <p>If you’re not redirected automatically, <a href="/login">click here to log in</a>.</p>
</body>
</html>
{{end}}


{{define "login"}}
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Login</title>
	<style>
		form {
			max-width: 400px;
			margin: auto;
		}
		label {
			display: block;
			margin-top: 10px;
		}
		input {
			width: 100%;
			padding: 8px;
			margin-top: 5px;
		}
	</style>
</head>
<body>
	<h1>Login</h1>
	<form method="post" action="/login">
		<label for="username">Username:</label>
		<input type="text" id="username" name="username" required>
		
		<label for="password">Password:</label>
		<input type="password" id="password" name="password" required>
		
		<input type="submit" value="Login" style="margin-top:15px;">
	</form>
</body>
</html>
{{end}}

{{define "mainPage"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Main Page</title>
</head>
<body>
    <h1>Enter your wallet address</h1>
    <form method="POST" action="/main">
        <input type="text" name="walletAddress" required>
        <input type="submit" value="Submit">
    </form>
</body>
</html>
{{end}}

{{define "walletActions"}}
<!DOCTYPE html>
<html>
<head><title>Welcome</title></head>
<body>
    <h1>Your wallet address is {{.WAddress}}</h1>
	<hr>

    <h2>Balance</h2>
    {{range .WBalance}}
        <p>{{.}}</p>
    {{end}}
	<hr>

    <h2>Actions</h2>
	<form action="/transactions" method="GET">
        <button type="submit">Get Transactions History</button><br />
    </form>
    <form>
        <button type="button">Send Currency</button><br />
        <button type="button">Get Currency Transactions History</button><br />
        <button type="button">Delete Wallet</button><br />
    </form>
</body>
</html>
{{end}}

`))
