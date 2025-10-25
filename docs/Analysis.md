**Problem description**

The goal in go-reloaded is to create a tool with which the user can give a .txt file as input and the output that will be produced is the contents of the file in the desired form.


**Rules**
In order to achieve that outcome, the contents of the file will have to use certain code words immediately after every word the user desires to change.
The code words that can be used are:
(hex) : Replaces the word before with its decimal version. **NOTE:** In order for this to work the word before (hex) must be a hexadecimal number

(i.e: **Input:** In July of 164 (hex) BC Alexander the Great was born.

**Output:** In July of 356 BC Alexander the Great was born.)


(bin): Replace the word before with the decimal version of the word. **NOTE:** In order for this to work the word will always be a binary number

(i.e: **Input:** I’m 10100 (bin) years old.

**Output:** I’m 20 years old.) 


(up): Converts the word before with the Uppercase version of it.

(i.e: **Input:** I’m here at last (up).

**Output:** I’m here at LAST.) 


(low): Converts the word before with the Lowercase version of it.

(i.e: **Input:** It’s not of HIGH (low) risk.

**Output:** It’s not of high risk.) 


(cap): Converts the word before with the capitalized version of it.

(i.e: **Input:** I live in Athens, greece (cap).

**Output:** I live in Athens, Greece.) 


If (low), (up), (cap) are used with a number right after them, then the corresponding function would occur on the corresponding number of words.

(i.e: **Input:** I like the movies of jason statham (cap, 2).

**Output:** I like the movies of Jason Statham.) 


Every instance of the punctuations (.  ,  ! ? : ;) should be close to the previous word and with space apart from the next one.


(i.e: **Input:** Good morning ,it is nice to meet you all !


**Output:** Good morning, it is nice to meet you all!) 


If punctuation is used in a grouped format such as (…), then it should perform as a regular instance of punctuation, but it should also maintain the punctuation group as intended.

(i.e: **Input:** So … what do you think about this project ?

**Output:** So… what do you think about this project?) 


If there is an instance of the punctuation mark (‘), then it should be always found with another instance of it and they should enclose, with no spaces, the word they are being used for.

(i.e: **Input:** Totally ‘ meant this.

**Output:** Totally ‘meant’ this.) 

**NOTE:** In case there are more than one word between ‘  ‘ marks, then the program must place these marks right next to the corresponding words (first mark should be left of the first word and last mark right of the last word)

(i.e: **Input:** As the people say ‘ Teamwork makes the dream work ‘

**Output:** As the people say ‘Teamwork makes the dream work‘) 


In case there is an a used then the program must ensure that it changes to an if the next word begins with a vowel (a, e, i, o, u) or a h.

(i.e: **Input:** A apple a day keeps the doctor away.

**Output:** An apple a day keeps the doctor away.) 

**NOTE:** In the English grammar the usage of a and an it doesn’t actually refers to the first letter of the word that is after them, but to the sound they make, if it’s vowel sound or a consonant sound.

(i.e: I found an advertisement for a university not far from here.)

-advertisement sounds like ad-ver-tis-ment, and it needs an because it begins with a vowel sound.

-university sounds like you-ni-ver-si-ty, and it needs a because it begins with a consonant sound.

So this brings us to an impasse regarding this particular rule. One possible solution is to make a learning and growing database for the system that builds up a vocabulary in order to automate this particular use.


Pipeline vs Finite State Machine

Comparison

A pipeline focuses on speeding up computation by dividing the main task into stages and executing different subtasks. It’s like an assembly line for instructions.

(i.e: A CPU has different stages to process different data.)


An FSM focuses on governing control flow through predictable states and inputs. It’s like a traffic signal controller that reacts to inputs step-by-step.

(i.e: Traffic lights change their state mainly in a clock time transition, but there can be the factor of the pedestrian button being used.)

Benefits
Pipeline can increase instruction throughput and has efficient use of CPU components due to parallel operation. It’s also highly scalable, because you can add more stages to increase potential clock speed.

FSM has a simple and predictable behavior making it easier to model and verify and it can be ideal for real-time and event-driven systems.

Limitations

Pipeline control can become complex for handling hazards such as:

Structural Hazard where two stages need the same resource simultaneously.
Data Hazard where one instruction depends on the result of a previous one not yet completed.
Control Hazard which occurs due to branch or jump instructions (the next instruction’s address changes unexpectedly).

These hazards can be handled by forwarding (bypassing), stalling (inserting bubbles) or branch prediction.

Also, if there are uneven stage durations there can be bottlenecks and we can have diminishing returns beyond a certain number of stages.


FSM is not efficient for data-heavy or computationally complex tasks and the number of states can grow exponentially with input complexity which makes it harder to modify once implemented in hardware.

Personal choice

In my opinion Pipeline and FSM should be used together because a pipeline provides the performance backbone, while an FSM acts as its control mechanism. That means that the two work in harmony to create efficient, intelligent digital systems.

Since the question is about choosing one of them and why, I choose Pipeline because it forms the backbone of high-performance digital systems and can even contain FSMs as control mechanisms.

Why does Pipeline form the backbone of high-performance digital systems?
Because it has higher efficiency, it has highly scalable performance, it’s foundational in modern CPUs, FSM can still exist inside it and it excels in general utility in comparison to FSM. Pipelines aren’t limited to CPUs, they’re used in graphics processors, DSPs, neural accelerators, and network routers. Wherever throughput matters, Pipeline prevails over FSM.

Last but not least, Pipeline fits better in this project because text modification commands can be processed as sequential stages such as reading, parsing, transforming, and writing, which aligns naturally with the pipeline model.
