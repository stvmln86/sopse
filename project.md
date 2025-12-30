# Project Notes 

## 2025-12-30
- [x] Complete SQLite rewrite.
- [ ] Add all handlers (create user, set pair, get user/pair, delete user/pair).
- [ ] Add background routine to clean out expired Pairs and rate limits.
  - [x] Add a `-taskWait` flag for background task sleep durations.
  - [ ] Delete pairs using SQLite `where` query, not Go code.
