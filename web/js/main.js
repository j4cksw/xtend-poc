var playerName;

function onNameEntered(){

    playerName = document.getElementById("player_name").value;

    if (playerName === ""){
        return;
    }

    disablePlayerRegistration();
    popWaitingMessage(playerName);
    connect(playerName, handleWebSocketMessage);
}

function disablePlayerRegistration(){
    $('#player_name').prop('disabled',true);
    $('#player_ok').prop('disabled',true);
}

function popWaitingMessage(playerName){
    $('#player_wait_message').text("Ok " + playerName + " Please wait...");
}
