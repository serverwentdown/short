package main

const AssetIndex = `
<!DOCTYPE html>
<link href="https://fonts.googleapis.com/css?family=Lato:400,700" rel="stylesheet">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>Shorten URL</title>
<style>
*, *:before, *:after {
	box-sizing: border-box;
}
html {
	font-size: 16px;
	line-height: 1.2;
	font-family: Lato;
}
body {
	margin: 0 auto;
	padding: 1rem 1rem;
	max-width: 640px;
	display: flex;
	flex-direction: column;
	align-items: stretch;
	justify-content: center;
}
input, textarea {
	font-size: inherit;
	font-family: inherit;
	line-height: inherit;
	-webkit-appearance: none;
	margin: 0;
	padding: 0.5rem;
	border: none;
	background: none;
	resize: none;
	display: inline-block;
}
input[type=submit], input[type=button] {
	font-weight: bold;
	background: #00cecf;
	color: #fff;
}
input[type=submit]:active, input[type=button]:active {
	box-shadow: -0.2rem 0.2rem 0.4rem #009e9f inset;
	background: #00bebf;
}
.wrapper {
	display: flex;
	margin: 1rem;
	border: 2px solid #00cecf;
	border-radius: 0.3rem;
	overflow: hidden;
}
.wrapper > :nth-child(1) {
	flex-grow: 1;
	width: 120px;
}
</style>
<form class="wrapper" action="/new" method="post">
	<input type="text" placeholder="https://" name="url">
	<input type="submit" value="Shorten">
</form>
<div class="wrapper">
	<input type="text" id="short" readonly>
	<input type="button" id="copy" value="Copy">
</div>
<script>
	let shortEle = document.querySelector("#short")
	let copyEle = document.querySelector("#copy")
	document.querySelector("form").addEventListener("submit", e => {
		e.preventDefault()
		copyEle.value = "Copy"
		fetch("/new", {
			method: e.target.method,
			body: new FormData(e.target),
		}).then(res => {
			return { text: res.text(), res: res }
		}).then(o => {
			if (!o.res.ok) throw Error(o.text)
			return o.text
		}).then(short => {
			console.log(short)
			document.querySelector("[name=url]").value = "";
			shortEle.value = short
		}).catch(err => {
			console.error(err)
			shortEle.value = err
		})
	})
	copyEle.addEventListener("click", e => {
		shortEle.select()
		let success = document.execCommand("copy")
		if (success) copyEle.value = "Copied!"
		setTimeout(() => copyEle.value = "Copy", 2000)
	})
</script>
`
