const express = require('express');
const router = express.Router();

// Define routes for product-related requests
router.get('/', (req, res) => {
    res.send('List of products');
});

router.post('/', (req, res) => {
    res.send('Product created');
});

router.get('/:id', (req, res) => {
    res.send(`Product details for ID: ${req.params.id}`);
});

router.put('/:id', (req, res) => {
    res.send(`Product updated for ID: ${req.params.id}`);
});

router.delete('/:id', (req, res) => {
    res.send(`Product deleted for ID: ${req.params.id}`);
});

module.exports = router;