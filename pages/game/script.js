console.log("Username is", username)

const raiseButton = document.getElementById('button-raise');
const callButton = document.getElementById('button-call');
const foldButton = document.getElementById('button-fold');

callButton.addEventListener('click', function () {
    let sum = Bet - players.get(4).CurrentBet
    if (sum > players.get(4).Balance) {
        sum = players.get(4).Balance
    }
    if (curStep) {
        const data = `{
            "position" : ${playerPosition},
            "action" : "call",
            "sum" : ${sum}
        }`;
        SendMessage(data);
        curStep = false;
    }

})

foldButton.addEventListener('click', function () {
    if (curStep) {
        const data = `{
            "position" : ${playerPosition},
            "action" : "fold",
            "sum" : 0 
        }`;
        inGame = false;
        curStep = false;
        SendMessage(data);
    }

})

raiseButton.addEventListener('click', function () {
    let sum = 0
    do {
        sum = prompt("Type your sum");
    } while (sum > players.get(4).Balance)
    if (sum && curStep) {
        curStep = false
        const data = `{
        "position" : ${playerPosition},
        "action" : "raise",
        "sum" : ${sum}
    }`;
        SendMessage(data)
    }
})

function changeBank(sum) {
    bankSum+=sum
    changeElement("bank", bankSum + "$")
}

function changeCard(cardID, card) {
    document.getElementById(cardID).textContent=(card["Value"] + "\n" + card["Suit"]);
}

function changeElement(ID, text) {
    document.getElementById(ID).textContent=text;
}


function addPlayer(player) {
    const parentElement = document.getElementById('table');
    let gameStatus = 'Out of game'
    if (player["InGame"]){
        gameStatus = 'In game'
    }

    // Создаем новый элемент div
    const newDiv = document.createElement("div");
    const pos = getGamePosition(player["Position"]);
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
                <span id="player_${pos}_cur_bet">$${player["CurrentBet"]}</span>
            </div>
        </div>
        <div id="player${pos}inf" class="info-left">
            <div class="username">${player["username"]}</div>
            <div id="player_${pos}_sum" class="total-money">$${player["Balance"]}</div>
            <div id="player_${pos}_status" class="status">${gameStatus}</div>
        </div>
        
    `;
    } else {
        newDiv.innerHTML = `
        <div class="player-cards">
            <div class="current-bet">
                <img src="/pages/game/poker-chip.png" alt="Image" height="30px" width="30px">
                <div class="spacer-vertical"></div>
                <span id="player_${pos}_cur_bet">$${player["CurrentBet"]}</span>
            </div>
            <div class="spacer-vertical"></div>  
            <div class="player-card-right card1 card"></div>
            <div class="player-card-right card2 card"></div>
        </div>
        <div id="player${pos}inf" class="info-right">
            <div class="username">${player["username"]}</div>
            <div id="player_${pos}_sum" class="total-money">$${player["Balance"]}</div>
            <div id="player_${pos}_status" class="status">${gameStatus}</div>
        </div>
    `;
    }

    parentElement.appendChild(newDiv);
}
