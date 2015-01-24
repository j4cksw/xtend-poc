function onNameEntered(){

    var playerName = document.getElementById("player_name").value;

    if (playerName === ""){
        return;
    }

    disablePlayerRegistration();
    popWaitingMessage(playerName);
    //connect(playerName, popGameScreen);
    popGameScreen();
}

function disablePlayerRegistration(){
    $('#player_name').prop('disabled',true);
    $('#player_ok').prop('disabled',true);
}

function popWaitingMessage(playerName){
    $('#player_wait_message').text("Ok " + playerName + " Please wait...");
}

function popGameScreen(){
    initGame();
    removeRegistrationElements();
}

function removeRegistrationElements(){
    $('#player_registration').hide();
    $('#player_wait_message').hide();
}
