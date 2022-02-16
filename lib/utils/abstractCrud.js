const pool = require('./pool.js');

async function insert(tableName, obj) {
  const columns = Object.keys(obj).join(', ');
  const vars = Object.values(obj).map((val, index) => `$${index + 1}`).join(', ');
  const vals = Object.values(obj);
  const { rows } = await pool.query(`
    INSERT INTO ${tableName} (${columns})
    VALUES (${vars})
    RETURNING *;
  `, vals);

  return rows[0];
}

async function getAll(tableName) {
  const { rows } = await pool.query(`
    SELECT * FROM ${tableName};
  `);

  return rows;
}

async function getById(tableName, id) {
  const { rows } = await pool.query(`
    SELECT * FROM ${tableName}
    WHERE id=$1;
  `, [id]);
  return rows[0];
}

async function updateById(tableName, id, obj) {
  const colVals = Object.entries(obj).map((entry, index) => `${entry[0]}=$${index + 2}`).join(', ');
  const vals = [id, ...Object.values(obj)];
  const { rows } = await pool.query(`
    UPDATE ${tableName}
    SET ${colVals}
    WHERE id=$1
    RETURNING *;
  `, vals);

  return rows[0];
}

async function deleteById(tableName, id) {
  const { rows } = await pool.query(`
    DELETE FROM ${tableName}
    WHERE id=$1
    RETURNING *;
  `, [id]);

  return rows[0];
}

module.exports = {
  insert,
  getAll,
  getById,
  updateById,
  deleteById
};
