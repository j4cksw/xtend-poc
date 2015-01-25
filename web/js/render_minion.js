function renderMinion(info){
    console.log("Minion built");
    var g = game.add.graphics(info.minion.x, info.minion.y);
    g.beginFill(info.minion.color, 1);
    var minion = g.drawRect(0, 0, 8, 8);
    minions[info.minion.playerName].add(minion);

    game.physics.arcade.enable(minion);

    if( (playerName === info.minion.playerName) && isRallyPointSet() ){
        moveMinion(minion, {
            x: rallyPoint.x,
            y: rallyPoint.y
        });
    }
}
