const pool = require('../lib/utils/pool');
const setup = require('../data/setup');
const request = require('supertest');

const url = `http://localhost:${process.env.PORT}`;

describe('backend routes', () => {
  beforeEach(() => {
    return setup(pool);
  });

  afterAll(() => {
    pool.end();
  });

  it('can get all cats', async () => {
    const res = await request(url).get('/cats');
    const body = res.body;

    expect(body).toEqual(seedData);
  });

  it('can get a cat', async () => {
    const { body } = await request(url).get('/cats/1');

    expect(body).toEqual(seedData[0]);
  });

  it('can post a cat', async () => {
    const data = {
      name: 'albert',
      weight: 1.5
    };

    const { body } = await request(url).post('/cats')
      .send(data)
      .set('Content-Type', 'application/json')
      .set('Accept', 'application/json');

    data.id = 4;
    expect(body).toEqual(data);
  });

  it('can put a cat', async () => {
    const data = {
      name: 'albert',
      weight: 1.7
    };

    const { body } = await request(url).put('/cats/1')
      .send(data)
      .set('Content-Type', 'application/json')
      .set('Accept', 'application/json');

    data.id = 1;
    expect(body).toEqual(data);

  });

  it('can delete a cat', async () => {
    await request(url).delete('/cats/1');

    const { body } = await request(url).get('/cats');

    // eslint-disable-next-line no-unused-vars
    const [deleted, ...expected] = [...seedData];
    expect(body).toEqual(expected);
  });
});

const seedData = [
  {
    id: 1,
    name: 'kevin',
    weight: 1.2
  },
  {
    id: 2,
    name: 'chungus',
    weight: 42
  },
  {
    id: 3,
    name: 'pico',
    weight: 0.001
  }
];
