## **Task 6 – Punctuation Spacing & Grouping Rules**

### **Goal / Functionality**

- Ensure punctuation `. , ! ? : ;` is attached to the previous word and followed by a single space (except end-of-line), and handle grouped punctuation like `...` (ellipsis) preserving grouping as `…` or `...` behavior specified.

### **TDD Steps**

1. **Write failing tests:** From Golden Test: `Punctuation tests are ... kinda boring ,what do you think ?` → `Punctuation tests are... kinda boring, what do you think?`.

2. **Implement minimal code:** Normalize spaces around punctuation during token rejoin stage: ensure no space before punctuation token and ensure single space after unless punctuation ends the line or followed by another punctuation that groups.

3. **Refactor:** Add special-case grouping for `...` sequences to remain concatenated and treated as a single punctuation token.

### **Validation**

- Unit tests for comma spacing, exclamation grouping `!!`, ellipsis `...`, and sequences like `!!?`.

- Confirm rejoined output matches Golden Test expectations.

### **Learning Summary**

- Learn deterministic whitespace normalization, rules about punctuation adjacency, and writing reassembly logic for tokens.

### **Why TDD**

- Whitespace and punctuation problems are brittle; tests capture formatting expectations precisely.

### **Alternatives**

- Use a two-pass approach: (1) normalize punctuation tokens, (2) collapse spaces — clearer and more maintainable.

### **AI assistance**

- AI can generate many punctuation edge-case tests automatically.

### **Quick Learning Note**

- When reassembling tokens, treat `space` tokens as conditional: include them only if previous token wasn't punctuation, etc.

### **Reference**

- https://pkg.go.dev/strings#Builder

### **Challenge**

- How will you treat `She said ‘ incredible work ‘ indeed.` so quotes touch inner words but quotes themselves are spaced correctly outside?