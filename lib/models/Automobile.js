const { insert, getAll, getById, updateById, deleteById } = require('../utils/abstractCrud.js');

module.exports = class Automobile {
  id;
  model;
  hp;

  constructor(row) {
    this.id = row.id;
    this.model = row.model;
    this.hp = row.hp;
  }

  static async insert(automobile) {
    const rows = await insert('automobiles', automobile);

    return new Automobile(rows);
  }

  static async getAll() {
    const rows = await getAll('automobiles');

    const automobiles = rows.map(row => new Automobile(row));
    return automobiles;
  }

  static async getById(id) {
    const row = await getById('automobiles', id);

    if(!row) return null;
    return new Automobile(row);
  }
  
  static async updateById(id, automobile) {
    const row = await updateById('automobiles', id, automobile);

    return new Automobile(row);
  }

  static async deleteById(id) {
    const row = await deleteById('automobiles', id);
    return new Automobile(row);
  }
};
