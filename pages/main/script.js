const randomGameBtn = document.getElementById('randomGameBtn');
const gameCodeBtn = document.getElementById('gameCodeBtn'); //TODO

const usernameInput = document.getElementById('username')

// Code for connecting to a random game
randomGameBtn.addEventListener('click', function() {
  const url = "http://localhost:8080/api/random"
  if (usernameInput.value === '') {
    alert('Пожалуйста, заполните поле "Username"!');
  } else {
    fetch(url, {
      headers: {
        'Content-Type': 'application/json'
      },
    }).then(response => response.json())
        .then(data => {
            sessionStorage.setItem('username', usernameInput.value);
            const gameID = data.id;
            window.location.href = `/connect?id=${gameID}`;
        })
        .catch(error => {
          console.error(error);
        });
  }

});