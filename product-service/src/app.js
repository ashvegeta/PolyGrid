const express = require('express');
const bodyParser = require('body-parser');
const routes = require('./routes/index');

const app = express();
const PORT = process.env.PORT || 3000;

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

app.use('/api/products', routes);

app.listen(PORT, () => {
    console.log(`Product service is running on port ${PORT}`);
});

module.exports = app;