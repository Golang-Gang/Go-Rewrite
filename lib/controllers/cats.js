const { Router } = require('express');
const Cat = require('../models/Cat.js');
const { validateCat } = require('../utils/validation.js');

module.exports = Router()
  .post('/cats', async (req, res, next) => {
    try {
      const cat = req.body;
      validateCat(cat);
      const result = await Cat.insert(cat);
      res.json(result);
    } catch(error) {
      next(error);
    }
  })
  .get('/cats', async (req, res, next) => {
    try {
      const cats = await Cat.getAll();
      res.json(cats);
    } catch(error) {
      next(error);
    }
  })
  .get('/cats/:id', async (req, res, next) => {
    try {
      const { id } = req.params;
      const cat = await Cat.getById(id);
      res.json(cat);
    } catch(error) {
      next(error);
    }
  })
  .put('/cats/:id', async (req, res, next) => {
    try {
      const { id } = req.params;
      const cat = req.body;
      validateCat(cat);
      const existingCat = await Cat.getById(id);

      if(!existingCat) {
        const err = new Error(`cat with id: ${id} not found`);
        err.status = 404;
        throw err;
      }

      const newCat = await Cat.updateById(id, cat);
      res.json(newCat);
    } catch(error) {
      next(error);
    }
  })
  .delete('/cats/:id', async (req, res, next) => {
    try {
      const { id } = req.params;
      const existingCat = await Cat.getById(id);

      if(!existingCat) {
        const err = new Error(`cat with id: ${id} not found`);
        err.status = 404;
        throw err;
      }
      
      const deletedCat = await Cat.deleteById(id);
      res.json(deletedCat);
    } catch(error) {
      next(error);
    }
  });
