var game;

function initGame(){
    game = new Phaser.Game(800, 600, Phaser.CANVAS, 'phaser-example', {
        create: create
    });
}

function create(){

    //var base = new Phaser.Circle(400, 300)

    var base = game.add.graphics(0, 0);
    base.beginFill(0xFFFF0B, 1);
    base.drawCircle(200, 300, 100);

    var base_2 = game.add.graphics(0, 0);
    base_2.beginFill(0xBC1C22, 1);
    base_2.drawCircle(600, 300, 100);
    console.log(base);
    console.log(base_2);

    console.log(base === base_2);
}
