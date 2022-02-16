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

  it('can get all cats', async () => {
    const { body } = await request(app).get('/cats');

    expect(body).toEqual(seedData);
  });

  it('can get a cat', async () => {
    const { body } = await request(app).get('/cats/1');

    expect(body).toEqual(seedData[0]);
  });

  it('can post a cat', async () => {
    const data = {
      name: 'albert',
      weight: 1.5
    };

    const { body } = await request(app).post('/cats')
      .send(data)
      .set('Content-Type', 'application/json')
      .set('Accept', 'application/json');

    data.id = '4';
    expect(body).toEqual(data);
  });

  it('can put a cat', async () => {
    const data = {
      name: 'albert',
      weight: 1.7
    };

    const { body } = await request(app).put('/cats/1')
      .send(data)
      .set('Content-Type', 'application/json')
      .set('Accept', 'application/json');

    data.id = '1';
    expect(body).toEqual(data);

  });

  it('can delete a cat', async () => {
    await request(app).delete('/cats/1');

    const { body } = await request(app).get('/cats');

    // eslint-disable-next-line no-unused-vars
    const [deleted, ...expected] = [...seedData];
    expect(body).toEqual(expected);
  });
});

const seedData = [
  {
    id: '1',
    name: 'kevin',
    weight: 1.2
  },
  {
    id: '2',
    name: 'chungus',
    weight: 42
  },
  {
    id: '3',
    name: 'pico',
    weight: 0.001
  }
];
