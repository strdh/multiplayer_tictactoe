const user1 = document.getElementById('user1');
const user2 = document.getElementById('user2');
const user1Status = document.getElementById('user1_status');
const user2Status = document.getElementById('user2_status');
const user1Score = document.getElementById('user1_score');
const user2Score = document.getElementById('user2_score');
const scoreRound = document.getElementById('score_round');
const pclose = document.getElementById('pclose');

pclose.addEventListener('click', function() {
  window.location.href = "home.html"
});

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
}, 1000);

let condition = [1, 2, 2, 0, 0, 2, 1, 0, 0];
const boxes = document.getElementsByClassName("cell");
for (let i = 0; i < boxes.length; i++) {
  let current = boxes[i];

  if (condition[i] == 1) {
    current.textContent = 'X';
  } else if (condition[i] == 2) {
    current.textContent = 'O';
  }
  // current.addEventListener('click', function () {
  //   let content = this.textContent;
  //   if (content === "" || content === "O") {
  //     this.textContent = "X";
  //   } else {
  //     this.textContent = "O";
  //   }
  // });
}