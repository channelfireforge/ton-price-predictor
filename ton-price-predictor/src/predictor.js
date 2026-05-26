function predictPrice(prices) {
  if (prices.length < 3) {
    return { prediction: null, confidence: 0 };
  }

  const recent = prices.slice(-3);
  const avg = recent.reduce((a, b) => a + b, 0) / recent.length;
  const change = (recent[2] - recent[0]) / recent[0] * 100;

  let confidence = Math.min(Math.abs(change) / 5, 1);
  let direction = change >= 0 ? 'up' : 'down';
  let predictedPrice = avg * (1 + change / 100);

  return {
    prediction: predictedPrice.toFixed(2),
    direction,
    confidence: confidence.toFixed(2)
  };
}

module.exports = predictPrice;