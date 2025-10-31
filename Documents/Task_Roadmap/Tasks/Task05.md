## **Task 5 – Numeric Conversions: `(hex)` and `(bin)`**

### **Goal / Functionality**

- Convert previous token from hex or binary string into decimal numerals.

### **TDD Steps**

1. **Write failing tests:** Examples: `1E (hex)` → `30`, `101 (bin)` → `5`, and Golden Test examples like `42 (hex)` and `10 (bin)` producing `66` and `2`.

2. **Implement minimal code:** Use `strconv.ParseInt` with base 16 for hex and base 2 for bin; replace token text with decimal string.

3. **Refactor:** Add validation and error handling (non-numeric input should leave original text or produce an error test).

### **Validation**

- Unit tests for uppercase/lowercase hex digits, `0x` prefixes (if you choose to support them), and invalid inputs.

- Integration tests from the Golden Test Set: combined paragraph conversions.

### **Learning Summary**

- How to parse numbers from strings, numeric bases, and error handling in Go.

### **Why TDD**

- Numeric parsing is fragile (bases, invalid digits); tests prevent accidental mis-parsing.

### **Alternatives**

- Use `fmt.Sscanf`, custom parsing, or big integers for very large inputs — prefer `strconv` for typical cases.

### **AI assistance**

- AI can produce fuzz tests for many numeric edge cases.

### **Quick Learning Note**

- `strconv.ParseInt(s, base, 64)` returns `(int64, error)`; handle errors explicitly.

### **Reference**

- https://pkg.go.dev/strconv#ParseInt

### **Challenge**

- What should the tool do if the hex token contains invalid characters (e.g., `1G (hex)`)?