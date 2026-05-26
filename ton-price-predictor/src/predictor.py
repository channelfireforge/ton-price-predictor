"""Main predictor logic: fetch, train, predict."""
from src.data_fetcher import fetch_historical_prices
from src.model import TONPriceModel

def get_prediction(days_back: int = 30, window: int = 5) -> float:
    """
    Fetch TON prices, train model, return predicted next price.
    """
    prices = fetch_historical_prices(days_back)
    if not prices:
        raise RuntimeError("Failed to fetch price data")
    
    model = TONPriceModel(window=window)
    model.train(prices)
    
    recent = prices[-window:]
    predicted = model.predict_next(recent)
    return predicted

if __name__ == "__main__":
    try:
        pred = get_prediction()
        print(f"Predicted next TON price: ${pred:.4f}")
    except Exception as e:
        print(f"Prediction failed: {e}")