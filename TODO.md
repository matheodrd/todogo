# TODO

- [ ] Support stateful execution for the `delete` command
- [ ] Support fuzzy finding using `fzf` for the `select` command
    - `todogo select aa338563-5903-4f63-b989-12a93f9d0716` -> select todo with given ID
    - `todogo select` -> open `fzf` and let user select a todo
- [ ] Consider creating a distinct package for the `update` command since it has subcommands
    - [Video that explains how to do this](https://www.youtube.com/watch?v=SSRIn5DAmyw&t=479s)
    - [Cobra user guide's section on how to organize subcommands](https://github.com/spf13/cobra/blob/4cafa37bc4bb85633b4245aa118280fe5a9edcd5/site/content/user_guide.md#organizing-subcommands)
- [ ] Consider using `XDG_STATE_HOME` (`~/.local/state`) instead of `XDG_CACHE_HOME` for caching state variables
    - The standard is pretty new ([2021 spec](https://specifications.freedesktop.org/basedir-spec/latest/))
    - There isn't a `os.UserStateDir` function yet (mentioned in [this issue](https://github.com/golang/go/issues/68988))
    - I think the [spec](https://specifications.freedesktop.org/basedir-spec/latest/) is a bit ambiguous and there are discussions on its meaning ([one of them](https://forum.atuin.sh/t/xdg-state-home-for-the-location-of-history-data/67/9))
    - Seems to be mainly used for logs and history
