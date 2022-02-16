const { insert, getAll, getById, updateById, deleteById } = require('../utils/abstractCrud.js');

module.exports = class Plane {
  id;
  model;
  cost;

  constructor(row) {
    this.id = row.id;
    this.model = row.model;
    this.cost = row.cost;
  }

  static async insert(plane) {
    const rows = await insert('planes', plane);

    return new Plane(rows);
  }

  static async getAll() {
    const rows = await getAll('planes');

    const planes = rows.map(row => new Plane(row));
    return planes;
  }

  static async getById(id) {
    const row = await getById('planes', id);

    if(!row) return null;
    return new Plane(row);
  }
  
  static async updateById(id, plane) {
    const row = await updateById('planes', id, plane);

    return new Plane(row);
  }

  static async deleteById(id) {
    const row = await deleteById('planes', id);
    return new Plane(row);
  }
};
