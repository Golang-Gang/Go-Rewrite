const { insert, getAll, getById, updateById, deleteById } = require('../utils/abstractCrud.js');

module.exports = class Cat {
  id;
  name;
  weight;

  constructor(row) {
    this.id = row.id;
    this.name = row.name;
    this.weight = row.weight;
  }

  static async insert(cat) {
    const row = await insert('cats', cat);
    return new Cat(row);
  }

  static async getAll() {
    const rows = await getAll('cats');
    return rows.map(row => new Cat(row));
  }

  static async getById(id) {
    const row = await getById('cats', id);
    if(!row) return null;
    return new Cat(row);
  }
  
  static async updateById(id, cat) {
    const row = await updateById('cats', id, cat);
    return new Cat(row);
  }

  static async deleteById(id) {
    const row = await deleteById('cats', id);
    return new Cat(row);
  }
};
