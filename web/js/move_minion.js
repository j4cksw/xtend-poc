function moveMinion(minion, pointer){
    if(minion.tween && minion.tween.isRunning){
        minion.tween.stop();
        minion.tween = undefined;
    }
    minion.tween = game.add.tween(minion);
    var duration = (game.physics.arcade.distanceToPointer(minion, pointer) / MINION_SPEED) * 1000;
    minion.tween.to({
        x: pointer.x,
        y: pointer.y
    }, duration, Phaser.Easing.Linear.None, true);

    minion.tween.start();
}
