const { insert, getAll, getById, updateById, deleteById } = require('../utils/abstractCrud.js');

module.exports = class Dog {
  id;
  name;
  is_good_boy;

  constructor(row) {
    this.id = row.id;
    this.name = row.name;
    this.is_good_boy = row.is_good_boy;
  }

  static async insert(dog) {
    const rows = await insert('dogs', dog);

    return new Dog(rows);
  }

  static async getAll() {
    const rows = await getAll('dogs');

    const dogs = rows.map(row => new Dog(row));
    return dogs;
  }

  static async getById(id) {
    const row = await getById('dogs', id);

    if(!row) return null;
    return new Dog(row);
  }
  
  static async updateById(id, dog) {
    const row = await updateById('dogs', id, dog);

    return new Dog(row);
  }

  static async deleteById(id) {
    const row = await deleteById('dogs', id);
    return new Dog(row);
  }
};
