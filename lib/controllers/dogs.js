const { Router } = require('express');
const Dog = require('../models/Dog.js');
const { validateDog } = require('../utils/validation.js');

module.exports = Router()
  .post('/', async (req, res, next) => {
    try {
      const dog = req.body;
      validateDog(dog);
      const result = await Dog.insert(dog);
      res.json(result);
    } catch(error) {
      next(error);
    }
  })
  .get('/', async (req, res, next) => {
    try {
      const dogs = await Dog.getAll();
      res.json(dogs);
    } catch(error) {
      next(error);
    }
  })
  .get('/:id', async (req, res, next) => {
    try {
      const { id } = req.params;
      const dog = await Dog.getById(id);
      res.json(dog);
    } catch(error) {
      next(error);
    }
  })
  .put('/:id', async (req, res, next) => {
    try {
      const { id } = req.params;
      const dog = req.body;
      validateDog(dog);
      const existingDog = await Dog.getById(id);

      if(!existingDog) {
        const err = new Error(`dog with id: ${id} not found`);
        err.status = 404;
        throw err;
      }

      const newDog = await Dog.updateById(id, dog);
      res.json(newDog);
    } catch(error) {
      next(error);
    }
  })
  .delete('/:id', async (req, res, next) => {
    try {
      const { id } = req.params;
      const existingDog = await Dog.getById(id);

      if(!existingDog) {
        const err = new Error(`dog with id: ${id} not found`);
        err.status = 404;
        throw err;
      }
      
      const deletedDog = await Dog.deleteById(id);
      res.json(deletedDog);
    } catch(error) {
      next(error);
    }
  });
