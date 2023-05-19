const urlParams = new URLSearchParams(window.location.search);
const host = window.location.host
const gameID = urlParams.get('id');
const username = sessionStorage.getItem('username')
console.log("Username is", username)

const socket = new WebSocket(`ws://${host}/api/ws?id=${gameID}`);
let playerPosition = 0;

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
            for (const player of data["players"]) {
                if (player && !(player["username"] === username)) {
                    addPlayer(player)
                }

            }
        })
        .catch(error => {
            console.error(error);
        });
});

// обработчик события получения сообщения от сервера
socket.addEventListener('message', (event) => {
    const message = JSON.parse(event.data);
    console.log(`Message is ${event.data}`)

    if (message["player"]["username"] === username) {
        playerPosition = message["player"]["Position"]
    }

    if (message["event"] === "new_player" && !(message["player"]["Position"] === playerPosition) ){
        addPlayer(message["player"])
        console.log("HERE")
    }
});

// обработчик события закрытия соединения
socket.addEventListener('close', (event) => {
    console.log('WebSocket connection closed');
});

// обработчик события ошибки соединения
socket.addEventListener('error', (event) => {
    console.log('WebSocket connection error');
});

function addPlayer(player) {
    const parentElement = document.getElementById('table');

    // Создаем новый элемент div
    const newDiv = document.createElement("div");
    const pos = ( 4 + player["Position"] - playerPosition) % 7
    newDiv.id = `player${pos}`;
    if (pos <= 3 && pos !== 0) {
        newDiv.innerHTML = `
        <div class="player-cards">
            <div class="player-card-left card1 card"></div>
            <div class="player-card-left card2 card"></div>
            <div class="spacer-vertical"></div>
            <div class="current-bet">
                <img src="/pages/game/poker-chip.png" alt="Image" height=30px width=30px>
                <div class="spacer-vertical"></div>
                <span>$${player["CurrentBet"]}</span>
            </div>
        </div>
        <div class="info-left">
            <div class="username">${player["username"]}</div>
            <div class="total-money">$${player["Balance"]}</div>
        </div>
        
    `;
    } else {
    newDiv.innerHTML = `
        <div class="player-cards">
            <div class="current-bet">
                <img src="/pages/game/poker-chip.png" alt="Image" height="30px" width="30px">
                <div class="spacer-vertical"></div>
                <span>$${player["CurrentBet"]}</span>
            </div>
            <div class="spacer-vertical"></div>  
            <div class="player-card-right card1 card"></div>
            <div class="player-card-right card2 card"></div>
        </div>
        <div class="info-right">
            <div class="username">${player["username"]}</div>
            <div class="total-money">$${player["Balance"]}</div>
        </div>
    `;
    }

    parentElement.appendChild(newDiv);
}