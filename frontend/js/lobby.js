const joinRoom = document.getElementById("join_room");
const makeRoom = document.getElementById("make_room");
const lclose = document.getElementById("lclose");

joinRoom.addEventListener('click', function() {
  window.location.href = "inputcode.html";
});

makeRoom.addEventListener('click', function() {
  window.location.href = "waitroom2.html";
});

lclose.addEventListener('click', function() {
  window.location.href = "home.html";
});