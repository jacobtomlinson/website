function initBinary() {
  canvas = document.getElementById("animation");
  if (!canvas) {
    return;
  }
  content = "";
  characters = "01";
  for (var i = 0; i < 12000; i++) {
    content += characters.charAt(Math.floor(Math.random() * characters.length));
    if (i % 250 == 0 && i > 0) {
      content += "<br />";
    }
  }
  canvas.innerHTML = content;

  setInterval(animateBinary, 100);
  window.addEventListener("mousemove", animateBinary);
}

function animateBinary() {
  canvas = document.getElementById("animation");
  content = canvas.innerHTML;
  idx = Math.floor(Math.random() * content.length);
  if (content.charAt(idx) === "0") {
    canvas.innerHTML =
      content.substring(0, idx) + "1" + content.substring(idx + 1);
  }
  if (content.charAt(idx) === "1") {
    canvas.innerHTML =
      content.substring(0, idx) + "0" + content.substring(idx + 1);
  }
}

window.addEventListener("load", initBinary);
