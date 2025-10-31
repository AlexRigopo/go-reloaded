## **Task 7 – Apostrophe / Quote Grouping (`‘ ... ‘`) Rules**

### **Goal / Functionality**

- Ensure opening and closing single-quotes `‘`/`’` (or ASCII `'`) wrap the enclosed words with no internal leading/trailing space and no spaces between the quote and its adjacent word(s).

### **TDD Steps**

1. **Write failing tests:** From Golden Test: `She said ‘ incredible work ‘ indeed.` → `She said ‘incredible work‘ indeed.` and multi-word cases like `‘ Teamwork makes the dream work ‘`.

2. **Implement minimal code:** In tokenization or a dedicated pass, detect quote-open and quote-close and trim internal leading/trailing spaces inside them; reattach quotes to first/last enclosed tokens.

3. **Refactor:** Make logic robust to nested or unmatched quotes (log or ignore).

### **Validation**

- Tests for single and multi-word quoted fragments, apostrophes in contractions (`can’t`) should be preserved.

### **Learning Summary**

- Learn contextual token operations and bracket-matching algorithms (stack-based).

### **Why TDD**

- Ensures quote handling doesn't break contractions or mis-handle nested cases.

### **Alternatives**

- Normalize all typographic quotes to a consistent character set first, or handle ASCII and Unicode quotes separately.

### **AI assistance**

- AI can produce tests for contractions and nested quoting scenarios.

### **Quick Learning Note**

- Use a stack to match quote tokens and operate on the span between matched indices.

### **Reference**

- https://pkg.go.dev/container/list (idea for stack) or simply use slice as stack.

### **Challenge**

- How would you differentiate an opening quote from a closing quote in ambiguous contexts?