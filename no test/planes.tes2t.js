const pool = require('../lib/utils/pool');
const setup = require('../data/setup');
const request = require('supertest');
const app = require('../lib/app');

describe('backend routes', () => {
  beforeEach(() => {
    return setup(pool);
  });

  afterAll(() => {
    pool.end();
  });

  it('can get all planes', async () => {
    const { body } = await request(app).get('/planes');

    expect(body).toEqual(seedData);
  });

  it('can get a plane', async () => {
    const { body } = await request(app).get('/planes/1');

    expect(body).toEqual(seedData[0]);
  });

  it('can post a plane', async () => {
    const data = {
      model: 'albert',
      cost: '$1.50'
    };

    const { body } = await request(app).post('/planes')
      .send(data)
      .set('Content-Type', 'application/json')
      .set('Accept', 'application/json');

    data.id = '4';
    expect(body).toEqual(data);
  });

  it('can put a plane', async () => {
    const data = {
      model: 'albert',
      cost: '$1.70'
    };

    const { body } = await request(app).put('/planes/1')
      .send(data)
      .set('Content-Type', 'application/json')
      .set('Accept', 'application/json');

    data.id = '1';
    expect(body).toEqual(data);

  });

  it('can delete a plane', async () => {
    await request(app).delete('/planes/1');

    const { body } = await request(app).get('/planes');

    // eslint-disable-next-line no-unused-vars
    const [deleted, ...expected] = [...seedData];
    expect(body).toEqual(expected);
  });
});

const seedData = [
  {
    id: '1',
    model: 'mustang',
    cost: '$3.50'
  },
  {
    id: '2',
    model: '777',
    cost: '$123,456.12'
  },
  {
    id: '3',
    model: '787',
    cost: '$7,654,321.01'
  }
];
