# Project Notes 

## 2025-12-30
- [x] Complete SQLite rewrite.
- [ ] Add check for maximum pairs per user.
- [ ] Add remaining handlers (get pair, delete user/pair).
- [ ] Add `neat.Name` to force pair names to lowercase alphanumeric with dashes.
- [ ] Add background routine to clean out expired Pairs and rate limits.
  - [x] Add a `-taskWait` flag for background task sleep durations.
  - [ ] Delete pairs using SQLite `where` query, not Go code.
