const { Router } = require('express');
const Automobile = require('../models/Automobile.js');
const { validateAutomobile } = require('../utils/validation.js');

module.exports = Router()
  .post('/automobiles', async (req, res, next) => {
    try {
      const automobile = req.body;
      validateAutomobile(automobile);
      const result = await Automobile.insert(automobile);
      res.json(result);
    } catch(error) {
      next(error);
    }
  })
  .get('/automobiles', async (req, res, next) => {
    try {
      const automobiles = await Automobile.getAll();
      res.json(automobiles);
    } catch(error) {
      next(error);
    }
  })
  .get('/automobiles/:id', async (req, res, next) => {
    try {
      const { id } = req.params;
      const automobile = await Automobile.getById(id);
      res.json(automobile);
    } catch(error) {
      next(error);
    }
  })
  .put('/automobiles/:id', async (req, res, next) => {
    try {
      const { id } = req.params;
      const automobile = req.body;
      validateAutomobile(automobile);
      const existingAutomobile = await Automobile.getById(id);

      if(!existingAutomobile) {
        const err = new Error(`automobile with id: ${id} not found`);
        err.status = 404;
        throw err;
      }

      const newAutomobile = await Automobile.updateById(id, automobile);
      res.json(newAutomobile);
    } catch(error) {
      next(error);
    }
  })
  .delete('/automobiles/:id', async (req, res, next) => {
    try {
      const { id } = req.params;
      const existingAutomobile = await Automobile.getById(id);

      if(!existingAutomobile) {
        const err = new Error(`automobile with id: ${id} not found`);
        err.status = 404;
        throw err;
      }
      
      const deletedAutomobile = await Automobile.deleteById(id);
      res.json(deletedAutomobile);
    } catch(error) {
      next(error);
    }
  });
