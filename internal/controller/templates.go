package controller

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
			<button type="submit">Sign Up</button>
		</form>
	</div>
	<div class="button-form">
		<form action="/signin" method="get">
			<button type="submit">Sign In</button>
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
	<title>Sign Up</title>
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
	<h1>Sign Up</h1>
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

{{define "signin"}}
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Sign In</title>
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
	<h1>Sign In</h1>
	<form method="post" action="/signin">
		<label for="username">Username:</label>
		<input type="text" id="username" name="username" required>
		
		<label for="password">Password:</label>
		<input type="password" id="password" name="password" required>
		
		<input type="submit" value="Login" style="margin-top:15px;">
	</form>
</body>
</html>
{{end}}
`))
