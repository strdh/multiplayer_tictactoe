var boxes = document.getElementsByClassName("cell");

for (var i = 0; i < boxes.length; i++) {
  let current = boxes[i];
  current.addEventListener('click', function () {
    let content = this.textContent;
    if (content === "" || content === "O") {
      this.textContent = "X";
    } else {
      this.textContent = "O";
    }
  });
}