## **Task 3 – Single-Word Case Transformations (`(up)`, `(low)`, `(cap)`)**

### **Goal / Functionality**

- Implement single-word `(up)`, `(low)`, and `(cap)` operations: transform the preceding word accordingly.

### **TDD Steps**

1. **Write failing tests:** Use Golden Test examples such as `it (cap)` → `It` and `LAST (low)` → `last`.

2. **Implement minimal code:** On encountering `(up)` apply strings.ToUpper on the previous word, `(low)` → strings.ToLower, `(cap)` → strings.Title-style capitalizing first character.

3. **Refactor:** Move transformation logic into separate functions `applyUp`, `applyLow`, `applyCap`.

### **Validation**

- Unit tests for ASCII words, mixed-case words, and punctuation-adjacent words.

- Edge case: empty previous token should be a no-op (and produce a controlled error or log).

### **Learning Summary**

- Understand `strings` package (`ToUpper`, `ToLower`), rune handling (UTF-8), and safe manipulation when previous token might be punctuation.

### **Why TDD**

- Isolates transformation logic and ensures correct behavior before adding multi-word counts.

### **Alternatives**

- `unicode` package to more exactly control capitalization of runes for Unicode-aware casing.

### **AI assistance**

- AI can suggest Unicode-safe capitalization helpers and create property-based tests (e.g., inverse properties).

### **Quick Learning Note**

- Use `[]rune` when altering first character to avoid breaking multi-byte Unicode characters.

### **Reference**

- https://pkg.go.dev/strings

### **Challenge**

- How do you capitalize the first letter of “greece” while preserving the rest exactly as-is?