function handleWebSocketMessage(evt){
    console.log(evt);
    var eventData = JSON.parse(evt.data);
    if( eventData.action == "init" ){
    }
    if( eventData.action === "render_base" ){
        initGame({
            players:eventData.data.players
        });
        removeRegistrationElements();
    }
    if( eventData.action === "render_minion" ){
        renderMinion(eventData.data)
    }
}

function removeRegistrationElements(){
    $('#player_registration').hide();
    $('#player_wait_message').hide();
}
