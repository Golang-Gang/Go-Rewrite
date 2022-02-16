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

  it('can get all automobiles', async () => {
    const { body } = await request(app).get('/automobiles');

    expect(body).toEqual(seedData);
  });

  it('can get a automobile', async () => {
    const { body } = await request(app).get('/automobiles/1');

    expect(body).toEqual(seedData[0]);
  });

  it('can post a automobile', async () => {
    const data = {
      model: 'vroom vroom 2',
      hp: 1.597
    };

    const { body } = await request(app).post('/automobiles')
      .send(data)
      .set('Content-Type', 'application/json')
      .set('Accept', 'application/json');

    data.id = '4';
    expect(body).toEqual(data);
  });

  it('can put a automobile', async () => {
    const data = {
      model: 'vroom vroom 3',
      hp: 1.598
    };

    const { body } = await request(app).put('/automobiles/1')
      .send(data)
      .set('Content-Type', 'application/json')
      .set('Accept', 'application/json');

    data.id = '1';
    expect(body).toEqual(data);

  });

  it('can delete a automobile', async () => {
    await request(app).delete('/automobiles/1');

    const { body } = await request(app).get('/automobiles');

    // eslint-disable-next-line no-unused-vars
    const [deleted, ...expected] = [...seedData];
    expect(body).toEqual(expected);
  });
});

const seedData = [
  {
    id: '1',
    model: 'mustang',
    hp: 350.789654321
  },
  {
    id: '2',
    model: 'bug',
    hp: 63.2
  },
  {
    id: '3',
    model: 'transit',
    hp: 213.111
  }
];
