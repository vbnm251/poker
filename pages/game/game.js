let bankSum = 0
let inGame = false
let curStep = false

function bet(pos, sum) {
    players.get(pos).CurrentBet=sum;
    players.get(pos).Balance-=sum;
    changeElement(`player_${pos}_sum`, players.get(pos)["Balance"])
    changeElement(`player_${pos}_cur_bet`, players.get(pos)["CurrentBet"])
}

function clearAll() {
    const card = {Value : "", Suit : ""}
    for (let i = 0; i < 5; i++) {
        changeCard(`table${i + 1}`, card);
    }
    changeCard("player_card1",card);
    changeCard("player_card2",card);
    changeBank(-bankSum);
}

function getGamePosition(pos)  {
    return ( 4 + pos - playerPosition) % 7
}

function updateCallSum(sum) {
    changeElement('button-call', 'Call' + sum + '$')
}