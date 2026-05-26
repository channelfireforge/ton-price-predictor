"""Fetch TON price data from public API."""
import requests
import time
from typing import List, Dict

COINGECKO_URL = "https://api.coingecko.com/api/v3/coins/the-open-network/market_chart"

def fetch_historical_prices(days: int = 30) -> List[float]:
    """
    Fetch daily TON prices from CoinGecko.
    Returns list of closing prices in USD.
    """
    params = {
        "vs_currency": "usd",
        "days": days,
        "interval": "daily"
    }
    try:
        response = requests.get(COINGECKO_URL, params=params, timeout=10)
        response.raise_for_status()
        data = response.json()
        prices = [item[1] for item in data["prices"]]
        return prices
    except requests.RequestException as e:
        print(f"API error: {e}")
        return []

if __name__ == "__main__":
    prices = fetch_historical_prices(7)
    print(f"Fetched {len(prices)} prices: {prices[:5]}...")