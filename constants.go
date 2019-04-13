package main

var html string = `
<!DOCTYPE html>
    <head>
        <meta charset="utf-8">
        <title></title>
        <meta name="viewport" content="width=device-width, initial-scale=1">
		<link rel="stylesheet" href="">
		<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
		<script>
			$(document).ready(function(){
				$("#pitchbutton").click(function(){
					$.get("startPitch", function(data, status){
						//nothing to be done here
					});
				});
				$("#qnabutton").click(function(){
					$.get("startQna", function(data, status){
						//nothing to be done here
					});
				});
				$("#nextpitchbutton").click(function(){
					$.get("nextPitch", function(data, status){
						//nothing to be done here
					});
				});
				$("#nextphasebutton").click(function(){
					$.get("nextPhase", function(data, status){
						//nothing to be done here
					});
				});
			});
		</script>
		<style>
			body {
				margin: 1em;
				background: #333;
				color: #ccc;
				font: normal 3vw/1 "Montserrat", sans-serif;
				text-align: center;
				text-transform: uppercase;
			  }
			  
			  h1 {
				font-size: 5vw;
			  }
			  
			  .container {
				width: 100%;
				padding: 10px 0px;
				box-sizing: border-box;
			  }
			  .container .inner {
				width: 100%;
				position: relative;
				min-height: 150px;
				display: inline-block;
			  }

			  .btn {
				border: none;
				text-decoration: none;
				padding: 15px 25px;
				color: #ffffff;
				border-radius: 50px;
				display: inline-block;
				position: relative;
			  }
			  .btn.btn-1 {
				background-color: #e67e22;
				box-shadow: 0 8px #bf6516;
			  }
			  .btn.btn-1:active {
				box-shadow: 0 0px #bf6516;
				top: 8px;
			  }
			  .btn.btn-1:focus {
				outline: 0;
			  }
			  .btn.btn-2 {
				background-color: #3498db;
				box-shadow: 0 8px #217dbb;
			  }
			  .btn.btn-2:active {
				box-shadow: 0 0px #217dbb;
				top: 8px;
			  }
			  .btn.btn-2:focus {
				outline: 0;
			  }
		  #wrapper {
			width: 100%;
			margin: 0 auto;
			overflow: hidden;
			}
			#first, #second {
				width: 50%;
				float:left;
				box-sizing: border-box;
				-moz-box-sizing: border-box;
				-webkit-box-sizing: border-box;
				padding: 0px 20px 0px 0px;
			}
		  </style>
    </head>
	<body>
		<h1>Dragon's Den Control Panel</h1>
		<div id="wrapper" class="container">
			<div id="first"><button class="btn btn-1" id="pitchbutton" type="button">P</button></div>
			<div id="second"><button class="btn btn-1" id="qnabutton" type="button">Q</button></div>
		</div>
		<br />
		<div id="wrapper" class="container">
			<div id="first"><button class="btn btn-2" id="nextpitchbutton" type="button">Name</button></div>
			<div id="second"><button class="btn btn-2" id="nextphasebutton" type="button">Phase</button></div>
		</div>
    </body>
</html>
`
