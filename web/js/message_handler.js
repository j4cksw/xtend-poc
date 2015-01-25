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
        // setTimeout(function(){
        //     eventData.data.players.forEach(function(baseInfo){
        //         console.log("building base");
        //         drawBase(baseInfo.color, baseInfo.x, baseInfo.y, 100, baseInfo.name);
        //     });
        // }, 500);
    }
    if( eventData.action === "render_minion" ){
        renderMinion(eventData.data)
    }
}

function removeRegistrationElements(){
    $('#player_registration').hide();
    $('#player_wait_message').hide();
}
