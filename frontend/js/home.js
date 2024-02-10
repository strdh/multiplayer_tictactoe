const playWithRandom = document.getElementById("play_wrandom");
const playWithFriend = document.getElementById("play_wfriend");

playWithRandom.addEventListener('click', function() {
  const isUsernameExist = localStorage.getItem("tttusername");
  if (isUsernameExist) {
    window.location.href = "waitroom.html";
  } else {
    window.location.href = "username.html";
  }
});

playWithFriend.addEventListener('click', function() {
  const isUsernameExist = true;
  if (isUsernameExist) {
    window.location.href = "lobby.html";
  } else {
    window.location.href = "username.html";
  }
});