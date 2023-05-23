console.log("Username is", username)

let playerPosition = -1;
let bankSum = 0

const raiseButton = document.getElementById('button-raise');
const callButton = document.getElementById('button-call');
const foldButton = document.getElementById('button-fold');

function changeCard(cardID, card) {
    document.getElementById(cardID).textContent=(card["Value"] + "\n" + card["Suit"]);
}

function changeElement(ID, text) {
    document.getElementById(ID).textContent=text;
}

callButton.addEventListener('click', function () {
    const data = `{
        "position" : ${playerPosition},
        "action" : "call",
        "sum" : 200 
    }`
    SendMessage(data)
})

function addGameTurnIndicator(pos) {
    let cont = `player${pos}inf`;
    if (pos === 4) {
        cont = `player4`;
    }

    const container = document.getElementById(cont);

    // Создаем элемент красного кружка
    const gameTurnIndicator = document.createElement("div");
    gameTurnIndicator.classList.add("game-turn");

    // Добавляем красный кружок к элементу с определенным ID
    container.appendChild(gameTurnIndicator);
}

function removeGameTurnIndicator(pos) {
    let cont = `player${pos}inf`;
    if (pos === 4) {
        cont = `player4`;
    }
    const container = document.getElementById(cont);

    // Удаляем красный кружок из элемента с определенным ID
    const gameTurnIndicator = container.querySelector(".game-turn");
    if (gameTurnIndicator) {
        container.removeChild(gameTurnIndicator);
    }
}

function addPlayer(player) {
    const parentElement = document.getElementById('table');

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
                <span>$${player["CurrentBet"]}</span>
            </div>
        </div>
        <div id="player${pos}inf" class="info-left">
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
        <div id="player${pos}inf" class="info-right">
            <div class="username">${player["username"]}</div>
            <div class="total-money">$${player["Balance"]}</div>
        </div>
    `;
    }

    parentElement.appendChild(newDiv);
}

function getGamePosition(pos)  {
    return ( 4 + pos - playerPosition) % 7
}