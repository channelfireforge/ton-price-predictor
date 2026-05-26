# TON Price Predictor

A minimal machine learning project to predict TON cryptocurrency prices using linear regression.

## Setup

```bash
pip install -r requirements.txt
```

## Usage

```bash
python -m src.predictor
```

## Test

```bash
pytest tests/
```

## How it works

1. Fetches historical TON/USD prices from CoinGecko API
2. Trains a linear regression model on sliding windows of prices
3. Predicts the next price based on the most recent window