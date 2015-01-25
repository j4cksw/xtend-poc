var game;
var gameInfo;
var bases;
var minions = {};

function initGame(gameData){
    game = new Phaser.Game(800, 600, Phaser.CANVAS, 'phaser-example', {
        create: create,
        update: update
    });

    function create(){
        gameData.players.forEach(function(baseInfo){
            console.log("building base");
            minions[baseInfo.name] = game.add.group();
            drawBase(baseInfo.color, baseInfo.x, baseInfo.y, 100, baseInfo.name);
        });
    }

    function update(){
        if (game.input.activePointer.isDown)
        {
            console.log(game.input.activePointer);
        }
    }
}
