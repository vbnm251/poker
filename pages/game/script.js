const urlParams = new URLSearchParams(window.location.search);
const host = window.location.host
const gameID = urlParams.get('id');
const username = sessionStorage.getItem('username')

const socket = new WebSocket(`ws://${host}/api/ws?id=${gameID}`);

// обработчик события открытия соединения
socket.addEventListener('open', (event) => {
    console.log('WebSocket connection established');
    socket.send(`{
            "username" : "${username}"
        }`
    )
    fetch(`http://${host}/api/gameInfo?id=${gameID}`, {
        headers: {
            'Content-Type': 'application/json'
        },
    }).then(response => response.json())
        .then(data => {
            console.log(data)
        })
        .catch(error => {
            console.error(error);
        });
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