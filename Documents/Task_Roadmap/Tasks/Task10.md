## **Task 10 – Pipeline Implementation & Concurrency (Processing Stages)**

### **Goal / Functionality**

- Implement the pipeline architecture: separated stages, clear interfaces between stages, and optional concurrency for throughput (worker pools or goroutines with channels).

### **TDD Steps**

1. **Write failing tests:** Integration tests that exercise stage boundaries (e.g., ensure token produced by Stage A is correctly transformed by Stage B).

2. **Implement minimal code:** Implement sequential pipeline first (functions invoked in order). Then add an optional concurrent pipeline where safe (stateless stages).

3. **Refactor:** Add instrumentation (timings), and unit tests ensuring concurrency doesn't change results.

### **Validation**

- End-to-end tests from Golden Test Set must pass when pipeline runs sequentially and when run concurrent-safe mode.

- Benchmark tests to detect regressions.

### **Learning Summary**

- Learn channels, goroutines, race conditions, and `go test -race` for detecting concurrency issues.

### **Why TDD**

- Concurrency bugs are subtle; tests + race detector catch regressions early.

### **Alternatives**

- Keep pipeline sequential for simplicity; add concurrency only for heavy inputs.

### **AI assistance**

- AI can suggest concurrency-safe refactors and generate benchmarks.

### **Quick Learning Note**

- Use `go test -bench .` and `go test -race` to check performance and races.

### **Reference**

- https://golang.org/doc/effective_go#concurrency

### **Challenge**

- Which pipeline stages are safe to parallelize without shared-state locking?