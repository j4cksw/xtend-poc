var game;
var gameInfo;

function initGame(gameInfo){
    gameInfo = gameInfo;
    game = new Phaser.Game(800, 600, Phaser.CANVAS, 'phaser-example', {
        create: create
    });
}

function create(){
    var bases = [];
    bases.push(drawBase(0xFFFF0B, 200, 300, 100));
    bases.push(drawBase(0xBC1C22, 600, 300, 100));
}

function drawBase(color, x, y, size){
    var base = game.add.graphics(0, 0);
    base.beginFill(color, 1);
    base.drawCircle(x, y, size);
    return base;
}
