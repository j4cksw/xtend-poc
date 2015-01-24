function handleWebSocketMessage(evt){
    console.log(evt);
    var eventData = JSON.parse(evt.data);
    if( eventData.action == "init" ){
        // initGame({
        //     player:playerName
        // });
        // removeRegistrationElements();
    }
    if( eventData.action === "render_base" ){
        initGame({
            player:playerName
        });
        removeRegistrationElements();
        setTimeout(function(){
            eventData.data.players.forEach(function(baseInfo){
                console.log("building base");
                drawBase(baseInfo.color, baseInfo.x, baseInfo.y, 100, baseInfo.name);
            });
        }, 500);
    }
    if( eventData.action === "render_minion" ){
        // eventData.data = {
        //     name: "player1",
        //     x: 10,
        //     y: 10,
        //     color: ""
        // }
        drawMinion(eventData.data)
    }
}

function removeRegistrationElements(){
    $('#player_registration').hide();
    $('#player_wait_message').hide();
}

function drawMinion(info){
    console.log("Minion built");
    var minion = game.add.graphics(0, 0);
    minion.beginFill(info.color, 1);
    minion.drawRect(info.x+(info.size/2), info.y, 8, 8);

}
