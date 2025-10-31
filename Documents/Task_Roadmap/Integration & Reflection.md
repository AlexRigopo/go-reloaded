### **Final “Integration & Reflection” Summary**

- Pipeline architecture is used throughout: tokenization → quote handling → numeric transforms → case transforms → punctuation normalization → article correction → assembly. This aligns with the analysis recommendation to prefer Pipeline (scalable, stage-isolated) while keeping FSM-style checks inside stages when needed.

- Performance considerations:

    * Keep stages stateless where possible so concurrency (goroutines + channels) is safe.

    * Profile real inputs; avoid allocating large temporary buffers; use `strings.Builder` when reassembling strings.

    * If the article-correction DB grows, load it into an efficient in-memory Trie or hash map.

- Future improvements:

    * Pronunciation-based article correction using a phonetic/pronunciation dictionary.

    * Add an exceptions DB for `a/an` edge cases.

    * Support for additional commands and internationalization.