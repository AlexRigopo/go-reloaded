## **Task 8 – Article Correction (`a` → `an`) with Vowel & “h” Heuristic**

### **Goal / Functionality**

- Convert `a` to `an` when the next word begins with a vowel sound approximated by starting letter rules (vowels + `h` per spec) while leaving exceptions as future work (vowel sound exceptions like `university`).

### **TDD Steps**

1. **Write failing tests:** From Golden Test: `a untold story` → `an untold story`; `A honest person` → `An honest person`.

2. **Implement minimal code:** Token-scan pass: when token is `a` or `A` and next word begins with `[aeiouhAEIOUH]`, replace `a` with `an` or `A` with `An`.

3. **Refactor:** Abstract rule into `shouldUseAn(nextWord string) bool` to enable future DB/exception handling.

### **Validation**

- Tests for lowercase/uppercase, leading punctuation, and edge cases like `a` followed by quote.

- Golden Test examples with `A honest person and a elephant` producing `An honest person and an elephant`.

### **Learning Summary**

- Learn heuristic rules, building functions to be replaced by smarter modules (audio-based heuristics or dictionary lookup).

### **Why TDD**

- This grammar heuristic affects sentence-level correctness; tests guard behavior and future replacement of heuristic with a DB.

### **Alternatives**

- Use a pronunciation dictionary (CMU) data set (external) later vs simple heuristic now.

### **AI assistance**

- AI can generate a candidate list of exceptions and propose a small built-in exception dictionary.

### **Quick Learning Note**

- Keep heuristics encapsulated so they can be swapped for a dictionary-based approach later.

### **Reference**

- https://pkg.go.dev/strings#HasPrefix

### **Challenge**

- Why does `a university` remain `a university` even though `u` is a vowel letter?