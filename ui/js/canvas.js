const canvas = document.getElementById("drawing-board");

function draw(x, y) {
  if (canvas.getContext) {
    const ctx = canvas.getContext("2d");

    ctx.fillStyle = "#f00";
    ctx.fillRect(x, y, 10, 10);
  } else {
    console.error("Your browser doesn't support the canvas element");
  }
}

function checkCoordinates(e) {
  const x = e.offsetX;
  const y = e.offsetY;

  console.log(`X: ${x}, Y: ${y}`);

  draw(x, y);
}

function resizeCanvas() {
  console.log("inner height", window.innerHeight);
  console.log("inner width", window.innerWidth);

  canvas.height = window.innerHeight * 0.9;
  canvas.width = window.innerWidth * 0.7;
}

resizeCanvas();

window.addEventListener("resize", resizeCanvas);
canvas.addEventListener("mousedown", checkCoordinates);
