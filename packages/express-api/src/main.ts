import express from 'express';
const app = express();

app.get('/', (_req, res) => {
  res.send({ message: 'Hello API' });
});

app.listen(3000, () => {
  // Server is running
});
