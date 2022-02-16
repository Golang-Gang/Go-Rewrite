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

  it('can get all trains', async () => {
    const { body } = await request(app).get('/trains');

    expect(body).toEqual(seedData);
  });

  it('can get a train', async () => {
    const { body } = await request(app).get('/trains/1');

    expect(body).toEqual(seedData[0]);
  });

  it('can post a train', async () => {
    const data = {
      model: 'albert',
      manufacturer: 'better trains co'
    };

    const { body } = await request(app).post('/trains')
      .send(data)
      .set('Content-Type', 'application/json')
      .set('Accept', 'application/json');

    data.id = '4';
    expect(body).toEqual(data);
  });

  it('can put a train', async () => {
    const data = {
      model: 'albert',
      manufacturer: 'better trains co'
    };

    const { body } = await request(app).put('/trains/1')
      .send(data)
      .set('Content-Type', 'application/json')
      .set('Accept', 'application/json');

    data.id = '1';
    expect(body).toEqual(data);

  });

  it('can delete a train', async () => {
    await request(app).delete('/trains/1');

    const { body } = await request(app).get('/trains');

    // eslint-disable-next-line no-unused-vars
    const [deleted, ...expected] = [...seedData];
    expect(body).toEqual(expected);
  });
});

const seedData = [
  {
    id: '1',
    model: 'train1',
    manufacturer: 'train co'
  },
  {
    id: '2',
    model: 'train2',
    manufacturer: 'train co'
  },
  {
    id: '3',
    model: 'train3',
    manufacturer: 'train co'
  }
];
