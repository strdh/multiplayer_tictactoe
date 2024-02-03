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

const progressBar = document.querySelector('.progress-bar');
let progress = 0;

const loadingInterval = setInterval(() => {
  progress += 1;
  progressBar.style.width = progress + '%';

  if (progress >= 100) {
    clearInterval(loadingInterval);
  }
}, 100);

let seconds = 10;
const timerDisplay = document.getElementById('timer');

const interval = setInterval(() => {
  seconds--;
  timerDisplay.textContent = seconds;

  if (seconds === 0) {
    clearInterval(interval);
    timerDisplay.textContent = "Time's up!"; // Optional message after countdown
  }
}, 1000); // Update every 1 second (1000 milliseconds)
