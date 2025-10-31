## **Task 9 – Operation Ordering & Conflict Resolution**

### **Goal / Functionality**

- Define and implement rule order and conflict resolution (e.g., when punctuation and `(cap, N)` interact, or when numeric conversion and case transform target the same token).

### **TDD Steps**

1. **Write failing tests:** Combined operations from Golden Test Case 1 and the Comprehensive Stress Test that mix multiple command types and punctuation ordering.

2. **Implement minimal code:** Decide pipeline stages and implement them in order (tokenize → handle quotes → apply numeric conversions → apply multi-word case transforms → punctuation normalization → article correction → final assembly).

3. **Refactor:** Make pipeline configurable and document the chosen stage order with rationale from Analysis Document (pipeline preferred).

### **Validation**

- Run combined Golden Test cases covering interactions (e.g., `cap` near punctuation and `(low, N)` interactions).

### **Learning Summary**

- Learn how stage ordering changes outputs; architecture design trade-offs and building a deterministic pipeline.

### **Why TDD**

- Testing combined flows catches emergent bugs created by interacting transformations.

### **Alternatives**

- Event-driven FSM to handle dependent transforms; prefer pipeline for throughput and clarity (as Analysis Document argues).

### **AI assistance**

- AI can propose stage orders and generate integration tests exercising pairwise interactions.

### **Quick Learning Note**

- Keep transformations idempotent where possible — re-running a stage shouldn't change already-correct text.

### **Reference**

- https://pkg.go.dev/context (for cancellations/timeouts if pipeline becomes concurrent)

### **Challenge**

- Which should run first: numeric conversions or case transformations? Why?