const axios = require('axios');

async function fetchPrice() {
  try {
    const response = await axios.get('https://api.coingecko.com/api/v3/simple/price?ids=the-open-network&vs_currencies=usd');
    return response.data['the-open-network'].usd;
  } catch (error) {
    console.error('Error fetching TON price:', error.message);
    return null;
  }
}

module.exports = fetchPrice;