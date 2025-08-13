# Jujutsu (jj) - A Git-Compatible VCS with a Fresh Approach
Jujutsu (jj) is a modern, Git-compatible version control system that offers a simpler and more powerful workflow while keeping compatibility with existing Git repositories. It’s designed to make branching, rebasing, and history rewriting safer and easier.

## Key Features
- **Git-compatible** – You can use jj on top of your existing Git repository without breaking anything.
- **Simple branching model** – No detached HEAD confusion; jj always works with a single working copy and checkouts are fast.
- **Safe history rewriting** – Rewriting history (rebasing, editing commits) is straightforward and less error-prone.
- **Multiple working copies** – Work on different changes simultaneously without stashing.
- **Immutable commit IDs** – Commit IDs are based on content, making rebases and merges safer.
- **Undo for everything** – Almost all actions can be undone, so mistakes are less scary.

## Differences from Git
| Feature / Concept | Git                                    | Jujutsu (jj)                                    |
| ----------------- | -------------------------------------- | ----------------------------------------------- |
| Working directory | One at a time (switch with `checkout`) | Multiple working copies possible                |
| Branch model      | Named pointers (refs)                  | Changesets connected in a graph                 |
| Commit IDs        | SHA-1/SHA-256 hashes                   | Content-based immutable IDs                     |
| History rewriting | More manual and risky                  | Built-in safe rebase/edit commands              |
| Undo support      | Limited (`reflog`)                     | Built-in `jj undo` for almost all actions       |
| Learning curve    | Steeper, more plumbing commands        | Streamlined for common workflows                |
| Git repo usage    | Native                                 | Fully compatible (sync with `jj git push/pull`) |

## When to Use Jujutsu
- You like Git but want a **safer, simpler history editing experience**.
- You want **multiple working copies** without complex stash juggling.
- You want to **try a new VCS** while staying compatible with your Git repositories.

📚 **More info:** https://github.com/martinvonz/jj
