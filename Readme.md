---

# Golang Price Calculator

A modular Go (Golang) CLI project that reads product prices, applies customizable tax rates, and outputs tax-included prices in JSON format.

### ✅ Features:

* Reads prices from a text file or CLI input (swappable IO Managers).
* Applies multiple tax rates (0%, 7%, 10%, 15%).
* Outputs structured JSON files for each tax rate.
* Clean, modular architecture using interfaces and package separation.

---

### 📂 Project Structure:

* `main.go` – Entry point, loops through defined tax rates.
* `prices/` – Core logic for tax calculation and job processing.
* `fileManager/` – File I/O operations (reading input, writing JSON output).
* `conversion/` – Utility to convert string slices to float64.
* `cmdmanager/` – (Optional) CLI-based input/output manager for interactive usage.
* `iomanager/` – IOManager interface for flexible input/output implementations.

---

### ⚙️ How It Works:

1. Prices are read from `prices.txt`.
2. For each tax rate, a `TaxInclPriceJob` is created.
3. It calculates the tax-included price for each input price.
4. Results are saved in JSON files like `result_0.json`, `result_7.json`, etc.

---

### 📄 Example Output (`result_7.json`)

```json
{
  "tax_rate": 0.07,
  "input_prices": [10.49, 12, 15.89, 9.99],
  "tax_included_prices": {
    "10.49": "11.22",
    "12.00": "12.84",
    "15.89": "17.00",
    "9.99": "10.69"
  }
}
```

---

### 🚀 How to Run:

```bash
go run main.go
```

Make sure `prices.txt` exists in the root directory with one price per line.

---
