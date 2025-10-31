## **Task 11 – Comprehensive Golden Tests Integration**

### **Goal / Functionality**

- Implement full acceptance tests using all cases from Golden Test Set and ensure they pass as a suite.

### **TDD Steps**

1. **Write failing tests:** Convert all Golden Test inputs → expected outputs into integration tests (`testdata` or table-driven).

2. **Implement minimal code:** Fix gaps surfaced by failing tests iteratively.

3. **Refactor:** Clean codebase, add modular packages (e.g., `tokenizer`, `transform`, `pipeline`, `util`) and ensure tests still pass.

### **Validation**

- All Golden Test cases (Success & Audit/Stress) must pass.

- Add CI gating so PRs must pass the Golden Test suite.

### **Learning Summary**

- Practice combining unit tests into a reliable integration suite and maintaining test fixtures.

### **Why TDD**

- Acceptance tests define “done”; they ensure features meet requirements from the Analysis and Golden Test documents.

### **Alternatives**

- Use BDD frameworks (e.g., Godog) for acceptance tests, but table-driven native tests are sufficient here.

### **AI assistance**

- AI can translate the Word doc test cases into machine-readable test fixtures automatically.

### **Quick Learning Note**

- Place large input/output pairs in `testdata/` and load with `os.ReadFile` in tests.

### **Reference**

- https://pkg.go.dev/os#ReadFile

### **Challenge**

- Where should edge-case tests (invalid hex, unmatched quotes) live: unit tests or integration tests — why?