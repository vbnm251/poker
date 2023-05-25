let playerPosition = -1;
let bankSum = 0
let inGame = false
let curStep = false

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

function updateCallSum(sum) {
    changeElement('button-call', 'Call' + sum + '$')
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

function clearAll() {
    const card = {Value : "", Suit : ""}
    for (let i = 0; i < 5; i++) {
        changeCard(`table${i + 1}`, card)
    }
    changeCard("player_card1",card);
    changeCard("player_card2",card);
    changeBank(-bankSum);
}

function getGamePosition(pos)  {
    return ( 4 + pos - playerPosition) % 7
}