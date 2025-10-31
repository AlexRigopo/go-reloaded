## **Task 2 – Tokenization & Basic Word Model**

### **Goal / Functionality**

- Implement robust tokenization: preserve words, punctuation tokens, and inline command tokens like `(hex)`, `(up)`, `(cap, 2)`.

### **TDD Steps**

1. **Write failing tests:** Tests that input `"hello , world (up)"` produces token stream `["hello"," ",","," ","world"," ","(up)"]` (or equivalent structured tokens).

2. **Implement minimal code:** Add tokenizer that splits on spaces but extracts parenthesized tags as separate tokens and keeps punctuation separate.

3. **Refactor:** Introduce a Token struct with Text, Type (word/punct/command/space) for easier downstream processing.

### **Validation**

- Unit tests for typical inputs and edge cases (multiple spaces, leading/trailing punctuation).

- Verify tokens can be rejoined back to original text (idempotency test).

### **Learning Summary**

- Learn string handling in Go, regexp basics for token detection, and small data modeling (struct) for tokens.

### **Why TDD**

- Tokenization underpins all rules; tests prevent subtle parsing bugs from propagating.

### **Alternatives**

- Use bufio.Scanner with custom split function vs regex-based splitting — use regex for fine-grained token types.

### **AI assistance**

- AI can propose regex patterns and generate unit tests for tricky cases.

### **Quick Learning Note**

- Use regexp.MustCompile and FindAllStringIndex to locate commands/punctuation reliably.

### **Reference**

- https://pkg.go.dev/regexp

### **Challenge**

- How would you tokenize the input She said ‘ hello world ‘ (up) so the quotes and multi-word interior are preserved?