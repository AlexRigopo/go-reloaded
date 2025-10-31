## **Task 1 – Project & Test Harness Setup**

### **Goal / Functionality**

- Create the repository scaffold, module, and test harness so every subsequent task can run `go test` and produce deterministic results.

### **TDD Steps**

1. **Write failing test:** Add `transform_test.go` with a placeholder test `TestTransform_Noop` that asserts `Transform("") == ""`.

2. **Implement minimal code:** Create `transform.Transform(input string) string` returning `""`.

3. **Refactor:** Keep code minimal but add package docs and a simple README with `go test` instructions.

### **Validation**

- `go test` should pass after implementation; failing initially, then passing once the minimal function is added.

- CI (GitHub Actions) runs `go test ./....`

### **Learning Summary**

- Developer learns module initialization (`go mod init`), testing basics (`testing` pkg), simple package layout, and running tests locally and in CI.

### **Why TDD**

- Ensures a reproducible baseline and prevents scope creep; the first test documents the expected public API.

### **Alternatives**

- Use `gotests` generator vs hand-writing tests. Prefer hand-writing for clarity at project start.

### **AI assistance**

- AI can scaffold `go.mod`, CI YAML, and sample test templates automatically.

### **Quick Learning Note**

- `go test` discovers `*_test.go`; table-driven tests are handy for later.

### **Reference**

- https://pkg.go.dev/testing

### **Challenge**

- What's the `go test` exit code behavior when no tests are found vs when a test fails?