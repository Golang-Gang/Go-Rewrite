const { Router } = require('express');
const Plane = require('../models/Plane.js');
const { validatePlane } = require('../utils/validation.js');

module.exports = Router()
  .post('/planes', async (req, res, next) => {
    try {
      const plane = req.body;
      validatePlane(plane);
      const result = await Plane.insert(plane);
      res.json(result);
    } catch(error) {
      next(error);
    }
  })
  .get('/planes', async (req, res, next) => {
    try {
      const planes = await Plane.getAll();
      res.json(planes);
    } catch(error) {
      next(error);
    }
  })
  .get('/planes/:id', async (req, res, next) => {
    try {
      const { id } = req.params;
      const plane = await Plane.getById(id);
      res.json(plane);
    } catch(error) {
      next(error);
    }
  })
  .put('/planes/:id', async (req, res, next) => {
    try {
      const { id } = req.params;
      const plane = req.body;
      validatePlane(plane);
      const existingPlane = await Plane.getById(id);

      if(!existingPlane) {
        const err = new Error(`plane with id: ${id} not found`);
        err.status = 404;
        throw err;
      }

      const newPlane = await Plane.updateById(id, plane);
      res.json(newPlane);
    } catch(error) {
      next(error);
    }
  })
  .delete('/planes/:id', async (req, res, next) => {
    try {
      const { id } = req.params;
      const existingPlane = await Plane.getById(id);

      if(!existingPlane) {
        const err = new Error(`plane with id: ${id} not found`);
        err.status = 404;
        throw err;
      }
      
      const deletedPlane = await Plane.deleteById(id);
      res.json(deletedPlane);
    } catch(error) {
      next(error);
    }
  });
