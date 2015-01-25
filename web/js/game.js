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
        game.physics.startSystem(Phaser.Physics.ARCADE);

        gameData.players.forEach(function(baseInfo){
            console.log("building base");
            minions[baseInfo.name] = game.add.group();
            minions[baseInfo.name].enableBody = true;
            minions[baseInfo.name].physicsBodyType = Phaser.Physics.ARCADE;

            drawBase(baseInfo.color, baseInfo.x, baseInfo.y, 100, baseInfo.name);
        });

        game.input.onTap.add(function(pointer){

            setRallyPoint(pointer.x, pointer.y);

            minions[playerName].forEach(function(minion){
                moveMinion(minion, pointer);
            }, this);

        });
    }

    function update(){
        // if (game.input.)
        // {
        //     console.log(game.input.activePointer);
        // }
    }
}
