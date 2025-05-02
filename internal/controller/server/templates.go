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
	<!-- GET Transactions History -->
    <form action="/get_transactions_history" method="GET">
		<input type="hidden" name="address" value="{{.WAddress}}">
        
		<button type="submit">Get Transactions History</button>
    </form>

	<!-- SEND Currency -->
    <form action="/send_currency" method="POST">
        <!-- Add input fields for amount & recipient as needed -->
		<input type="hidden" name="address" value="{{.WAddress}}">
        
		<button type="submit">Send Currency</button>

		<!-- choose exactly one of two currencies -->
		<label for="currency">Currency:</label>
		<select name="currency" id="currency">
			<option value="badgercoin">badgercoin</option>
			<option value="catfishcoin">catfishcoin</option>
		</select>

		<!-- amount to send -->
		<label for="amount">Amount:</label>
		<input
		  type="number"
		  id="amount"
		  name="amount"
		  placeholder="Enter amount"
		  step="any"
		  min="0"
		  required
		>

		<!-- recipient address -->
    	<label for="receiver">Recipient Address:</label>
    	<input
    	  type="text"
    	  id="receiver"
    	  name="receiver"
    	  placeholder="Enter recipient address"
    	  required
    	>

		<!-- mineNow -->
    	<label for="mineNow">mineNow:</label>
    	<input
    	  type="text"
    	  id="mineNow"
    	  name="mineNow"
    	  placeholder="mineNow"
    	  required
    	>
    </form>
</body>
</html>
{{end}}

{{define "getTransactionsHistory"}}
<!DOCTYPE html>
<html>
<head>
    <title>Transactions History</title>
</head>
<body>
    {{if .WCurrency}}
        <h1>Transactions history of {{.WCurrency}} for {{.WAddress}}</h1>
    {{else}}
        <h1>Transactions history for {{.WAddress}}</h1>
    {{end}}
    <hr>

    {{range .WHistory}}
        <p>{{.}}</p>
    {{else}}
        <p>No transactions found.</p>
    {{end}}
</body>
</html>
{{end}}

{{define "successSendCurrency"}}
<!DOCTYPE html>
<html>
<head>
    <title>Send Currency</title>
</head>
<body>
	<h1>{{.WResult}}</h1>
	<hr>

    <h2>{{.WSender}} sent {{.WAmount}} {{.WCurrency}} to {{.WReceiver}}</h2>
</body>
</html>
{{end}}

`))
