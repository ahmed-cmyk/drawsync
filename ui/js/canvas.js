const canvas = document.getElementById("drawing-board");
let isDrawing = false;

canvas.addEventListener("mousedown", () => (isDrawing = true));
canvas.addEventListener("mouseup", () => (isDrawing = false));
canvas.addEventListener("mouseout", () => (isDrawing = false));
canvas.addEventListener("mousemove", draw);

function draw(e) {
  if (!isDrawing) return;

  const ctx = canvas.getContext("2d");

  ctx.lineWidth = 5;
  ctx.lineCap = "round";
  ctx.strokeStyle = "#f00";

  ctx.lineTo(e.offsetX, e.offsetY);
  ctx.stroke();
  ctx.beginPath();
  ctx.moveTo(e.offsetX, e.offsetY);
}

function resizeCanvas() {
  console.log("inner height", window.innerHeight);
  console.log("inner width", window.innerWidth);

  canvas.height = window.innerHeight * 0.9;
  canvas.width = window.innerWidth * 0.7;
}

resizeCanvas();

window.addEventListener("resize", resizeCanvas);
