const { insert, getAll, getById, updateById, deleteById } = require('../utils/abstractCrud.js');

module.exports = class Train {
  id;
  model;
  manufacturer;

  constructor(row) {
    this.id = row.id;
    this.model = row.model;
    this.manufacturer = row.manufacturer;
  }

  static async insert(train) {
    const rows = await insert('trains', train);

    return new Train(rows);
  }

  static async getAll() {
    const rows = await getAll('trains');

    const trains = rows.map(row => new Train(row));
    return trains;
  }

  static async getById(id) {
    const row = await getById('trains', id);

    if(!row) return null;
    return new Train(row);
  }
  
  static async updateById(id, train) {
    const row = await updateById('trains', id, train);

    return new Train(row);
  }

  static async deleteById(id) {
    const row = await deleteById('trains', id);
    return new Train(row);
  }
};
