var rallyPoint = {};

function setRallyPoint(x, y){
    rallyPoint.x = x;
    rallyPoint.y = y;
}

function isRallyPointSet(){
    return (rallyPoint.x !== undefined) && (rallyPoint.y !== undefined);
}
