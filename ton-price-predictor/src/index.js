const fetchPrice = require('./fetchPrice');
const predictPrice = require('./predictor');

const priceHistory = [];

async function main() {
  console.log('TON Price Predictor');
  console.log('Fetching current price...\n');

  const currentPrice = await fetchPrice();
  if (currentPrice) {
    priceHistory.push(currentPrice);
    console.log(`Current TON price: $${currentPrice}`);

    if (priceHistory.length >= 3) {
      const result = predictPrice(priceHistory);
      console.log(`Prediction: $${result.prediction} (${result.direction})`);
      console.log(`Confidence: ${result.confidence}`);
    } else {
      console.log('Need more data points for prediction.');
    }
  }
}

main();