// Create the canvas
var canvas = document.createElement("canvas");
var ctx = canvas.getContext("2d");
canvas.width = 800;
canvas.height = 600;
document.body.appendChild(canvas);


// Background image
var shipReady = false;
var shipImage = new Image();
shipImage.onload = function () {
	shipReady = true;
};
shipImage.src = "images/ship.png";

var ships = []

var ws = new WebSocket("ws://" + window.location.host + "/superstellar");

ws.onmessage = function(e) {
	ships = $.evalJSON(e.data).positions;
};

// Handle keyboard controls
var keysDown = {};

addEventListener("keydown", function (e) {
	var direction
	switch (e.keyCode) {
		case 38:
			direction = "up"
			break;
		case 40:
			direction = "down"
			break;
		case 37:
			direction = "left"
			break;
		case 39:
			direction = "right"
			break;
	}

	ws.send($.toJSON({client_id: 1, direction: direction}))
}, false);

// Draw everything
var render = function () {
	ctx.beginPath();
	ctx.rect(0, 0, 800, 600);
	ctx.fillStyle = "black";
	ctx.fill();

	if (shipReady) {
		for (i = 0; i < ships.length; i++) {
			ship = ships[i]

			ctx.translate(ship.x, ship.y);
			ctx.rotate(ship.angle);
			ctx.drawImage(shipImage, -30, -15);
			ctx.rotate(-ship.angle);
			ctx.translate(-ship.x, -ship.y);
		}
	}

	// Score
	ctx.fillStyle = "rgb(250, 250, 250)";
	ctx.font = "24px Helvetica";
	ctx.textAlign = "left";
	ctx.textBaseline = "top";
	ctx.fillText("Ships: " + ships.length, 680, 10);
};

// The main game loop
var main = function () {
	var now = Date.now();
	var delta = now - then;

	render();

	then = now;

	// Request to do this again ASAP
	requestAnimationFrame(main);
};

// Let's play this game!
var then = Date.now();
main();