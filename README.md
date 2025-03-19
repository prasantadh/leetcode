# LEETCODE

A simple approach to solve leetcode problems on my local neovim instance.
Primarily to get realtime LSP support from my editor (neovim).

## Requisite

I use bear to generate `compile_commands.json` for better LSP support.

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
