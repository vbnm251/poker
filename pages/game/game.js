let bankSum = 0
let inGame = false
let curStep = false

let Bet = 0

function bet(pos, sum) {
    let player = players.get(pos);
    player.CurrentBet += sum;
    player.Balance-=sum;
    players.set(pos, player)
    changeElement(`player_${pos}_sum`, players.get(pos)["Balance"]);
    changeElement(`player_${pos}_cur_bet`, players.get(pos)["CurrentBet"]);
}

function clearAll() {
    const card = {Value : "", Suit : ""}
    for (let i = 0; i < 5; i++) {
        changeCard(`table${i + 1}`, card);
    }
    changeCard("player_card1",card);
    changeCard("player_card2",card);
    changeBank(-bankSum);
    players.clear()
}

function getGamePosition(pos)  {
    return ( 4 + pos - playerPosition) % 7
}

function updateCallSum(sum) {
    let newTxt = 'Call ' + sum + '$'
    if (sum === 0) {
        newTxt = 'Check'
    }
    changeElement('button-call', newTxt)
}

function  clearBets() {
    console.log("CLeaning started")
    players.forEach(function (player, pos) {
        player.CurrentBet = 0
        changeElement(`player_${pos}_cur_bet`, 0);
    })
    updateCallSum(0)
    Bet = 0
}