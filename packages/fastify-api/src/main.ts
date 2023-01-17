import fastify from 'fastify';

const app = fastify();

// eslint-disable-next-line @typescript-eslint/no-unused-vars
app.get('/', async (_req, _res) => {
  return { message: 'Hello API' };
});

const start = async () => {
  try {
    await app.listen({ port: 3000 });
  } catch (err) {
    // Errors are logged here
    process.exit(1);
  }
};

start();
