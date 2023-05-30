const socket = new WebSocket(`ws://${host}/api/ws?id=${gameID}`);

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
            for (const player of data["game"]) {

                if (player){
                    if (!(player["username"] === username)) {
                        addPlayer(player)
                    }
                }
            }
        })
        .catch(error => {
            console.error(error);
        });
});

socket.addEventListener('message', (event) => {
    const message = JSON.parse(event.data);
    console.log(message);

    if (message["event"] === "new_player") {
        if (message["player"]["username"] === username) {
            playerPosition = message["player"]["Position"]
            console.log("Self player position is", playerPosition);
        }
        if (!(message["player"]["Position"] === playerPosition) ){
            addPlayer(message["player"]);
        }
    }

    else if (message["event"] === "distribution") {
        gamePeriod = "preflop"

        playerRole = message["role"]

        if (playerRole === "small_blind") {
            const data = `{
            "position" : ${playerPosition},
            "action" : "raise",
            "sum" : ${smallBlindBet}
        }`;
            SendMessage(data);
        }

        clearAll()
        inGame = true
        changeCard("player_card1", message["cards"][0]);
        changeCard("player_card2", message["cards"][1]);
    }

    else if (message["event"] === "flop") {
        gamePeriod="flop"

        clearBets()

        changeCard("table1", message["cards"][0]);
        changeCard("table2", message["cards"][1]);
        changeCard("table3", message["cards"][2]);
    }

    else if (message["event"] === "turn") {
        gamePeriod="turn"

        clearBets()

        changeCard("table4", message["card"]);
    }

    else if (message["event"] === "river") {
        gamePeriod="river"

        clearBets()
        changeCard("table5", message["card"]);
    }


    else if (message["event"] === "gamePlayers") {
        for (const player of message["players"]) {
            //todo : handle game info
            if (player) {
                const pos = getGamePosition(player["Position"])
                changeElement(`player_${pos}_sum`, player.Balance)
                changeElement(`player_${pos}_cur_bet`, player.CurrentBet)
                players.set(pos, player)
                if (player["Role"] === "small_blind") {
                    // if (player["Position"] === playerPosition) {
                    //     curStep = true
                    // }
                    changeElement(`player_${pos}_status`, "Current")
                } else {
                    changeElement(`player_${pos}_status`, "In Game")
                }
            }
        }
    }

    else if (message["event"] === "step") {
        if (message["pos"] !== -1) {
            changeElement(`player_${getGamePosition(message["pos"])}_status`, "Current")
        }
        if (message["pos"] === playerPosition) {
            curStep = true
        }
    }

    else if (message["event"] === "winners") {
        for (const player of message["winners"]) {
            const pos = getGamePosition(player["Position"])
            changeElement(`player_${pos}_status`, "Winner")
            changeElement(`player_${pos}_sum`, players.get(pos).Balance+=message["sum"])
        }
    }

    //game events: fold, call, raise
    else  {
        bet(getGamePosition(message["position"]), message["sum"])
        if (message["next"] !== -1) {
            if (message["next"] === playerPosition) {
                if (gamePeriod === "preflop" && playerRole === "big_blind" && !betCompleted) {
                    const data = `{
                        "position" : ${playerPosition},
                        "action" : "raise",
                        "sum" : ${smallBlindBet * 2}
                    }`;
                    SendMessage(data);
                    betCompleted = true;
                } else {
                    curStep = true
                }
            }
            changeElement(`player_${getGamePosition(message["next"])}_status`, "Current")
        }
        if (message["action"] === "fold") {
            changeElement(`player_${getGamePosition(message["position"])}_status`, "Out of game")
            return
        } else if (message["action"] === "raise") {
            Bet = players.get(getGamePosition(message["position"])).CurrentBet
            updateCallSum(Bet - players.get(4).CurrentBet)
        }
        changeElement(`player_${getGamePosition(message["position"])}_status`, "In Game")
        changeBank(message["sum"])
    }

});

socket.addEventListener('close', (event) => {
    console.log('WebSocket connection closed');
});

socket.addEventListener('error', (event) => {
    console.log('WebSocket connection error');
});

function SendMessage(data) {
    socket.send(data)
}