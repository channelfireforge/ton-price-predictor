"""Tests for TON price predictor."""
import sys
import os
sys.path.insert(0, os.path.abspath(os.path.join(os.path.dirname(__file__), '..')))

from src.model import TONPriceModel
from src.data_fetcher import fetch_historical_prices

def test_model_train_and_predict():
    """Test model with synthetic data."""
    prices = [1.0, 1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7]
    model = TONPriceModel(window=3)
    model.train(prices)
    pred = model.predict_next([1.5, 1.6, 1.7])
    assert 1.7 < pred < 2.0, f"Expected ~1.8, got {pred}"

def test_model_short_data():
    """Test that model raises error with insufficient data."""
    model = TONPriceModel(window=5)
    try:
        model.train([1.0, 1.1, 1.2])
        assert False, "Should have raised ValueError"
    except ValueError:
        pass

def test_fetch_api():
    """Test API fetch returns data (may be skipped if offline)."""
    prices = fetch_historical_prices(days=1)
    if prices:
        assert len(prices) > 0
        assert all(isinstance(p, (int, float)) for p in prices)
    else:
        print("API fetch returned empty (offline?)")