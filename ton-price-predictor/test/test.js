const predictPrice = require('../src/predictor');
const assert = require('assert');

const testPrices = [5.0, 5.2, 5.1, 5.3, 5.4];
const result = predictPrice(testPrices);

assert(result.prediction !== null, 'Prediction should not be null');
assert(['up', 'down'].includes(result.direction), 'Direction should be up or down');
assert(result.confidence >= 0 && result.confidence <= 1, 'Confidence should be between 0 and 1');

console.log('Test passed: Prediction works correctly');
console.log(`Prediction: $${result.prediction} (${result.direction}, confidence: ${result.confidence})`);