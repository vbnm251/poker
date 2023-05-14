const urlParams = new URLSearchParams(window.location.search);
const gameID = urlParams.get('id');
const socket = new WebSocket(`ws://localhost:8080/ws?id=${gameID}`);
const username = sessionStorage.getItem('username')

// обработчик события открытия соединения
socket.addEventListener('open', (event) => {
    console.log('WebSocket connection established');
    socket.send(`{
    "action" : "new_player",
    "data" : {
        "username" : "${username}"
    }
}`)
});

// обработчик события получения сообщения от сервера
socket.addEventListener('message', (event) => {
    console.log(`Received message: ${event.data}`);
});

// обработчик события закрытия соединения
socket.addEventListener('close', (event) => {
    console.log('WebSocket connection closed');
});

// обработчик события ошибки соединения
socket.addEventListener('error', (event) => {
    console.log('WebSocket connection error');
});