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

  it('can get all dogs', async () => {
    const { body } = await request(app).get('/dogs');

    expect(body).toEqual(seedData);
  });

  it('can get an dog', async () => {
    const { body } = await request(app).get('/dogs/1');

    expect(body).toEqual(seedData[0]);
  });

  it('can post an dog', async () => {
    const data = {
      name: 'albert',
      is_good_boy: true
    };

    const { body } = await request(app).post('/dogs')
      .send(data)
      .set('Content-Type', 'application/json')
      .set('Accept', 'application/json');

    data.id = '4';
    expect(body).toEqual(data);
  });

  it('can put an dog', async () => {
    const data = {
      name: 'albert',
      is_good_boy: true
    };

    const { body } = await request(app).put('/dogs/1')
      .send(data)
      .set('Content-Type', 'application/json')
      .set('Accept', 'application/json');

    data.id = '1';
    expect(body).toEqual(data);

  });

  it('can delete an dog', async () => {
    await request(app).delete('/dogs/1');

    const { body } = await request(app).get('/dogs');

    // eslint-disable-next-line no-unused-vars
    const [deleted, ...expected] = [...seedData];
    expect(body).toEqual(expected);
  });
});

const seedData = [
  {
    id: '1',
    name: 'spot',
    is_good_boy: true
  },
  {
    id: '2',
    name: 'jeep',
    is_good_boy: true
  },
  {
    id: '3',
    name: 'jeff',
    is_good_boy: true
  }
];
