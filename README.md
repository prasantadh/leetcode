# LEETCODE

A simple approach to solve [leetcode](leetcode.com) problems on my local neovim instance.
Primarily to get realtime LSP support from my editor ([neovim](neovim.io) with [lazyvim](lazyvim.org)).

## Requisite

I use [bear](https://github.com/rizsotto/Bear) to generate `compile_commands.json` for better LSP support.

## Installation

```bash
git clone git@github.com:prasantadh/leetcode.git
```

## Usage

```bash
# to copy template.cpp into src/242-valid-anagram.cpp
./practi.sh setup 242 valid anagram

# to run 242-valid-anagram.cpp
./practi.sh run 242
```

## Notes

- There are tags on `tags.h` that can be used to categorize problems
- Need ideas on what problem to solve next? Consult `next.md`
- Realized something important? Put it on `notes.md`

## PAQ (Potentially Asked Questions)

**Q:** Aren't there already plugins for this?

**A:** Yes, there are. [leetcode.nvim](https://github.com/kawre/leetcode.nvim) seems popular.
I like browsing leetcode. Just don't like having to run code
to get my compilation error. Or the janky vim key bindings.
I have a working setup for coding and would just like to use
that for practice.

**Q:** This isn't allowed during interview. Isn't this bad for interview prep?

**A:** Yes, it probably is. I am not concerned about interview scenarios.
I just want to practice some interesting leetcode problems during my free time.

**Q:** I see git history here has python and go and rust and what not. Why C++ now?

**A:** Yeah, I like all of them. Python is super quick for prototyping.
Love using it for CTFs. Even though it taught me a bunch of bad OOP practices, I think.
And when an algorithmic solution goes wrong, it makes me sometimes wonder if I got the
algorithm wrong or the language feature wrong.
Go makes concurency so easy. I have heard Go makes deploying microservices super easy
with its static binaries (I haven't tried this but will get to it at some point).
Didn't really enjoy its type system though.
Rust is amazing and I am trying to learn it more. Yet, making a simple linked list
is difficult due to the self-referential nature of the data structure.

In the past, I have tried to use leetcode to learn a language. I think this was misguided.
Need to learn a language? Read some books, watch some videos, build something.
Leetcode is best at teaching algorithms. I like C++ for learning algorithms.
Low level yet with way more than enough language features.
