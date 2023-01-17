import express from 'express';
const app = express();

app.get('/', (_req, res) => {
  res.send({ message: 'Express API GET' });
});

app.post('/', (_req, res) => {
  res.send({ message: 'Express API POST' });
});

const server = app.listen(3000, () => {
  console.log('Express started');
});

// eslint-disable-next-line @typescript-eslint/no-unused-vars
app.post('/shutdown', (_req, res) => {
  res.send('ok');
  server.close(() => {
    console.log('Express shutting down');
    process.exit(0);
  });
});
