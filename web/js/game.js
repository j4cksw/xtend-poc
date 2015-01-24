var game;
var gameInfo;
var bases;

function initGame(gameInfo){
    gameInfo = gameInfo;
    game = new Phaser.Game(800, 600, Phaser.CANVAS, 'phaser-example', {
        create: create
    });
}

function create(){
    // var bases = [];
    // bases.push(drawBase(0xFFFF0B, 100, 300, 100));
    // bases.push(drawBase(0xBC1C22, 700, 300, 100));
    //
    // bases.forEach(function(base){
    //     base.timer.start();
    // });
    ws.send(JSON.stringify({
        "action":"render_base",
        "data":{
            player:gameInfo.player,
        }
    }));
}

function drawBase(color, x, y, size){
    var base = game.add.graphics(0, 0);
    base.beginFill(color, 1);
    base.drawCircle(x, y, size);

    base.timer = game.time.create(false);
    base.timer.loop(1000, drawMinion, this, {"x":x, "y":y, "color":color, "size":size});
    base.timer.start();

    return base;
}

function drawMinion(info){
    var minion = game.add.graphics(0, 0);
    minion.beginFill(info.color, 1);
    minion.drawRect(info.x+(info.size/2), info.y, 16, 16);
}
