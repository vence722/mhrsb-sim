const items = require('./data.json');
const fs = require('fs');

const equipTypes = ['head', 'torso', 'arms', 'waist', 'legs']

const equipments = [];

for (const item of items) {
    const equipment = {
        name: item[0],
        type: equipTypes[item[3]],
        slots: item[10],
        skills: item[11],
        defense: item[12],
        resistance: item[13]
    }
    equipments.push(equipment);
}

fs.writeFileSync('./equipments.json', JSON.stringify(equipments, null, 4));