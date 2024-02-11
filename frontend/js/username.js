const uclose = document.getElementById("uclose");
const txtUsername = document.getElementById("txt_username");
const submitUsername = document.getElementById("submit_username");

uclose.addEventListener('click', function() {
  window.location.href = "home.html";
});

submitUsername.addEventListener('click', function() {
  const username = txtUsername.value;
  if (username && username != "") {
    const url = 'http://localhost:5000/ttt/username/create';
    const data = {
      username: username
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
    alert("username cannot be null");
  }
});

