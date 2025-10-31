## **Project’s Audit Test Cases**

- **Test Case 1: Combined Operations (cap, low, up with or without count and punctuation)**

    **Input:**
    `it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.`
    **Output:**
    `It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.`

    * Covers:
        * `(low, N)` applied to multiple words
        * `(up)` capitalization at the end
        * `(cap)` rule applied to last word
        * `(cap, N)` rule applied to multiple words
        * Space correction around punctuation

- **Test Case 2: Hexadecimal and Binary Conversion**

    **Input:**
    `Simply add 42 (hex) and 10 (bin) and you will see the result is 68.`
    **Output:**
    `Simply add 66 and 2 and you will see the result is 68.`

    * Covers:
        * Hexadecimal to decimal conversion
        * Binary to decimal conversion

- **Test Case 3: Punctuation and Grouping**

    **Input:**
    `Punctuation tests are ... kinda boring ,what do you think ?`
    **Output:**
    `Punctuation tests are... kinda boring, what do you think?`

    * Covers:
        * Punctuation spacing correction (commas, exclamation marks)
        * Multiple punctuation handling

- **Test Case 4: Article Correction**

    **Input:**
    `There is no greater agony than bearing a untold story inside you.`
    **Output:**
    `There is no greater agony than bearing an untold story inside you.`

    * Covers:
        * Article correction (“a” → “an”)
        * Mix of grammar and formatting logic

## **Success Test Cases**

- **Test Case 1: Hexadecimal and Binary Conversion**

    **Input:**
    `In total, we found 1E (hex) boxes and 101 (bin) coins in the cave.`
    **Output:**
    `In total, we found 30 boxes and 5 coins in the cave.`

    * Covers:
        * Hexadecimal to decimal conversion
        * Binary to decimal conversion

- **Test Case 2: Case Conversion with Counts**

    **Input:**
    `He whispered: This is the final test (up, 3) and it will be great.`
    **Output:**
    `He whispered: This is THE FINAL TEST and it will be great.`

    * Covers:
        * `(up, N)` transformation applied to multiple words

- **Test Case 3: Punctuation and Apostrophes**

    **Input:**
    `I can’t believe it ,this is amazing !! She said ‘ incredible work ‘ indeed.`
    **Output:**
    `I can’t believe it, this is amazing!! She said ‘incredible work‘ indeed.`

    * Covers:
        * Punctuation spacing correction (commas, exclamation marks)
        * Multiple punctuation handling
        * Apostrophe placement around multiple words

- **Test Case 4: Article Correction and Capitalization**

    **Input:**
    `A honest person and a elephant entered the arena in athens (cap).`
    **Output:**
    `An honest person and an elephant entered the arena in Athens.`

    * Covers:
        * Article correction (`a` → `an`)
        * `(cap)` rule applied to last word
        * Mix of grammar and formatting logic

- **Test Case 5: Combined Operations (low, up, punctuation, and grouping)**

    **Input:**
    `The SECRET WILL (low, 2) be revealed soon … or will it ? Maybe never (up)!`
    **Output:**
    `The secret will be revealed soon… or will it? Maybe NEVER!`

    * Covers:
        * `(low, N)` applied to multiple words
        * Grouped punctuation `…`
        * Space correction around punctuation
        * `(up)` capitalization at the end

## **Comprehensive Test Paragraph (“Stress Test”)**

**Input:**
`A old story tells that 1A (hex) warriors fought for 111 (bin) days. They shouted this is sparta (up, 3) ,but no one retreated ! The leader said ‘ courage and honor ‘ before entering the gate … In ancient greece (cap) ,it was a sacred rule never to surrender. Truly ,a hero’s fate awaits those who fight WITH PASSION (low, 2) !`

**Expected Output:**
`An old story tells that 26 warriors fought for 7 days. They shouted THIS IS SPARTA, but no one retreated! The leader said ‘courage and honor‘ before entering the gate… In ancient Greece, it was a sacred rule never to surrender. Truly, a hero’s fate awaits those who fight with passion!`

**Covers:**
- `(hex)` and `(bin)` conversions
- `(low, N)` multi-word lowercase
- Punctuation spacing and grouping
- Apostrophe rules with multiple words
- `(cap)` capitalization of last word in phrase
- Article correction (`a` → `an`)
- `(up, N)` multi-word uppercase
- Complex mixed case with realistic sentence flow
