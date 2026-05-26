"""Simple linear regression model for TON price prediction."""
import numpy as np
from sklearn.linear_model import LinearRegression
from typing import List, Tuple

class TONPriceModel:
    """Predict next TON price based on recent trend."""
    
    def __init__(self, window: int = 5):
        self.window = window
        self.model = LinearRegression()
        self.trained = False
    
    def _prepare_data(self, prices: List[float]) -> Tuple[np.ndarray, np.ndarray]:
        """Convert price list to features (windowed) and targets."""
        X, y = [], []
        for i in range(len(prices) - self.window):
            X.append(prices[i:i + self.window])
            y.append(prices[i + self.window])
        return np.array(X), np.array(y)
    
    def train(self, prices: List[float]) -> None:
        """Train model on historical prices."""
        if len(prices) < self.window + 1:
            raise ValueError(f"Need at least {self.window + 1} prices, got {len(prices)}")
        X, y = self._prepare_data(prices)
        self.model.fit(X, y)
        self.trained = True
    
    def predict_next(self, recent_prices: List[float]) -> float:
        """Predict next price given last `window` prices."""
        if not self.trained:
            raise RuntimeError("Model not trained yet")
        if len(recent_prices) != self.window:
            raise ValueError(f"Need exactly {self.window} prices")
        X = np.array([recent_prices])
        return float(self.model.predict(X)[0])