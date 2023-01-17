import fastify from 'fastify';

const app = fastify();

// eslint-disable-next-line @typescript-eslint/no-unused-vars
app.get('/', async (_req, _res) => {
  return { message: 'Fastify API GET' };
});

// eslint-disable-next-line @typescript-eslint/no-unused-vars
app.post('/', async (_req, _res) => {
  return { message: 'Fastify API POST' };
});

// eslint-disable-next-line @typescript-eslint/no-unused-vars
app.post('/shutdown', (_req, res) => {
  res.send('ok');
  app.close(() => {
    console.log('Fastify shutting down');
    process.exit(0);
  });
});

const start = async () => {
  try {
    app.server.on('listening', () => {
      console.log('Fastify started');
    });
    await app.listen({ port: 3001 });
  } catch (err) {
    // Errors are logged here
    process.exit(1);
  }
};

start();
