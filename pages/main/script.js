const randomGameBtn = document.getElementById('randomGameBtn');
const gameCodeBtn = document.getElementById('gameCodeBtn');

randomGameBtn.addEventListener('click', function() {
  // Code for connecting to a random game
  const data = { username: "example" };
  const url = "http://localhost:8080/random"

  fetch(url, {
    method: 'POST',
    headers: {
      'Connection' : 'upgrade',
      'Upgrade': 'websocket',
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(data)
  }).then(response => {
    // Обработка ответа
  }).catch(error => {
    // Обработка ошибок
  });

});