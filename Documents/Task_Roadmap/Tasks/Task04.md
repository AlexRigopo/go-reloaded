## **Task 4 – Multi-Word Case Transformations (`(up, N)`, `(low, N)`, `(cap, N)`)**

### **Goal / Functionality**

- Extend case transforms to apply to N previous words (e.g., `(cap, 2)` capitalizes the two words before the tag).

### **TDD Steps**

1. **Write failing tests:** From Golden Test: `jason statham (cap, 2)` → `Jason Statham`.

2. **Implement minimal code:** Parse numeric argument inside tag; apply transformation across previous N word tokens (skip punctuation/space tokens).

3. **Refactor:** Ensure boundary handling when fewer than N words available; create helper selecting previous N word tokens.

### **Validation**

- Tests with mixed punctuation, line breaks, and insufficient prior words.

- Test that `(up, 3)` turns three words uppercase in He whispered: `This is the final test (up, 3)` producing `THIS IS THE FINAL TEST`.

### **Learning Summary**

- Learn command parsing, iterating over token streams, and robustly handling token types.

### **Why TDD**

- Guarantees correctness of selecting prior word tokens vs naive “previous N tokens” which would fail around punctuation.

### **Alternatives**

- Implement look-back using an index stack vs scanning backward each time — stack is O(1) per token.

### **AI assistance**

- AI can auto-generate tests for permutations of punctuation within the window.

### **Quick Learning Note**

- Use integer parsing via `strconv.Atoi` to parse counts in `(cap, 6)`.

### **Reference**

- https://pkg.go.dev/strconv

### **Challenge**

- How should `(low, 2)` behave when the next token is a closing quote or punctuation—do you count punctuation as a word?