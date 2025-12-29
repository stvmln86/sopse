# Project Notes 

## 2025-12-30
- [x] Rename `mockUser1` to just `mockUser`.
- [ ] Add capacity hints to all `make()` calls.
- [ ] Remove head from database paths, e.g.: `uuid` and `uuid.name`.
  - [ ] Remove `bolt.Join` and `bolt.Split`.
  - [ ] Change `Pair.Name` to just use `strings.Cut`. 
- [ ] Change `bolt.List` to return a list of maps all at once.
  - [ ] Change `User.ListPairs` to use `pair.New` to make new Pairs from maps.
- [ ] Add background routine to clean out expired Pairs and rate limits.
  - [ ] Add a `-taskWait` flag for background task sleep durations.
- [ ] Add remaining handlers (get pair, delete pair, delete user).
