const txtCode = document.getElementById("txt_code");
const joinRoom = document.getElementById("join_room");
const iclose = document.getElementById("iclose");

joinRoom.addEventListener('click', function() {
  const code = txtCode.value;
  if (code &&code != "") {
    const url = 'http://localhost:5000/ttt/room/join';
    const data = {
      code: code
    };

    fetch(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    })
    .then(response => {
      console.log('Response Status Code:', response.status);
      if (response.status === 200) {
        txtUsername.value = "";
      }
      return response.json();
    })
    .then(data => {
      console.log(data);
    })
    .catch((error) => {
      console.log('Error: ', error);
    });
  } else {
    alert("Code cannot be null");
  }
});

iclose.addEventListener('click', function() {
  window.location.href = "home.html";
});