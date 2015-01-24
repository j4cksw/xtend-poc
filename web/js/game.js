var game;
var gameInfo;
var bases;

function initGame(gameInfo){
    //gameInfo = gameInfo;
    game = new Phaser.Game(800, 600, Phaser.CANVAS, 'phaser-example', {
        create: create
    });
    console.log(game.add);
    function create(){
        // var bases = [];
        // bases.push(drawBase(0xFFFF0B, 100, 300, 100));
        // bases.push(drawBase(0xBC1C22, 700, 300, 100));
        //
        // bases.forEach(function(base){
        //     base.timer.start();
        // });
        // ws.send(JSON.stringify({
        //     "action":"render_base",
        //     "data":{
        //         player:gameInfo.player,
        //     }
        // }));
    }

}
