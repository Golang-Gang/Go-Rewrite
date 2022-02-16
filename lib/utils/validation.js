function isNullish(val) {
  return (typeof val === 'undefined' || val === null);
}

function areTypesMismatched(val1, val2) {
  return (typeof val1 !== typeof val2);
}

const dogStructure = {
  name: 'a',
  is_good_boy: true
};

const catStructure = {
  name: 'a',
  weight: 1.2
};

const planeStructure = {
  model: 'string',
  cost: '$123.21'
};

const trainStructure = {
  model: 'string',
  manufacturer: 'string'
};

const automobileStructure = {
  model: 'string',
  hp: 1.123
};

function validate(structure, instance) {
  Object.entries(structure).forEach(entry => {
    const [key, val] = entry;
    if(isNullish(instance[key]) || areTypesMismatched(val, instance[key]))
      throw new Error(`${key} must be specified and of type ${typeof val}`);
  });
}

function validateDog(dog) {
  validate(dogStructure, dog);
}

function validateCat(cat) {
  validate(catStructure, cat);
}

function validatePlane(plane) {
  validate(planeStructure, plane);
}

function validateTrain(train) {
  validate(trainStructure, train);
}

function validateAutomobile(automobile) {
  validate(automobileStructure, automobile);
}

module.exports = {
  validateDog,
  validateCat,
  validatePlane,
  validateTrain,
  validateAutomobile
};
