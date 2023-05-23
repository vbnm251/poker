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
    console.log(`Message is ${event.data}`);

    if (message["event"] === "new_player") {
        if (message["player"]["username"] === username) {
            playerPosition = message["player"]["Position"]
            console.log("Self player position is", playerPosition);
        }
        if (!(message["player"]["Position"] === playerPosition) ){
            addPlayer(message["player"]);
        }
    }

    else if (message["event"] === "preflop") {
        changeCard("table1", message["cards"][0]);
        changeCard("table2", message["cards"][1]);
        changeCard("table3", message["cards"][2]);
    }

    else if (message["event"] === "distribution") {
        changeCard("player_card1", message["cards"][0]);
        changeCard("player_card2", message["cards"][1]);
    }

    else if (message["event"] === "gamePlayers") {
        for (const player of message["players"]) {
            if (player) {
                players.push(player)
                if (player["Role"] === "small_blind") {
                    const pos = getGamePosition(player["Position"])
                    addGameTurnIndicator(pos)
                }
            }
        }
    }

    else if (message["action"] === "call") {
        removeGameTurnIndicator(getGamePosition(message["position"]))
        bankSum+=message["sum"]
        changeElement("bank", bankSum + "$")
        if (message["next"] !== -1) {
            addGameTurnIndicator(getGamePosition(message["next"]))
        }
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

function SendMessage(data) {
    socket.send(data)
    console.log("data has been sent")
}
