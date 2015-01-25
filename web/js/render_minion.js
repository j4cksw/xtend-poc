function renderMinion(info){
    console.log("Minion built");
    console.log(info);

    info.minions.forEach(function(minionInfo){
        var minion = game.add.graphics(0, 0);
        minion.beginFill(minionInfo.color, 1);
        minion.drawRect(minionInfo.x, minionInfo.y, 8, 8);
    });
    //minions[info.name].add(minion);
}
