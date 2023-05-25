console.log("Username is", username)

const raiseButton = document.getElementById('button-raise');
const callButton = document.getElementById('button-call');
const foldButton = document.getElementById('button-fold');

callButton.addEventListener('click', function () {
    if (curStep) {
        const data = `{
            "position" : ${playerPosition},
            "action" : "call",
            "sum" : 200 
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
    let sum = prompt("Type your sum");
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


