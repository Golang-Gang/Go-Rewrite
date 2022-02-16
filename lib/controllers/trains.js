const { Router } = require('express');
const Train = require('../models/Train.js');
const { validateTrain } = require('../utils/validation.js');

module.exports = Router()
  .post('/trains', async (req, res, next) => {
    try {
      const train = req.body;
      validateTrain(train);
      const result = await Train.insert(train);
      res.json(result);
    } catch(error) {
      next(error);
    }
  })
  .get('/trains', async (req, res, next) => {
    try {
      const trains = await Train.getAll();
      res.json(trains);
    } catch(error) {
      next(error);
    }
  })
  .get('/trains/:id', async (req, res, next) => {
    try {
      const { id } = req.params;
      const train = await Train.getById(id);
      res.json(train);
    } catch(error) {
      next(error);
    }
  })
  .put('/trains/:id', async (req, res, next) => {
    try {
      const { id } = req.params;
      const train = req.body;
      validateTrain(train);
      const existingTrain = await Train.getById(id);

      if(!existingTrain) {
        const err = new Error(`train with id: ${id} not found`);
        err.status = 404;
        throw err;
      }

      const newTrain = await Train.updateById(id, train);
      res.json(newTrain);
    } catch(error) {
      next(error);
    }
  })
  .delete('/trains/:id', async (req, res, next) => {
    try {
      const { id } = req.params;
      const existingTrain = await Train.getById(id);

      if(!existingTrain) {
        const err = new Error(`train with id: ${id} not found`);
        err.status = 404;
        throw err;
      }
      
      const deletedTrain = await Train.deleteById(id);
      res.json(deletedTrain);
    } catch(error) {
      next(error);
    }
  });
