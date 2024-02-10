const uclose = document.getElementById("uclose");
const txtUsername = document.getElementById("txt_username");
const submitUsername = document.getElementById("submit_username");

uclose.addEventListener('click', function() {
  window.location.href = "home.html";
});

submitUsername.addEventListener('click', function() {
  const username = txtUsername.value;
  if (username && username != "") {
    console.log(username);
    // make request to the backend
  } else {
    alert("username cannot be null");
  }
});

