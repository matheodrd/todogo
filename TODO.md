# TODO

- [ ] Support stateful execution for the `delete` command
- [ ] Support fuzzy finding using `fzf` for the `select` command
    - `todogo select aa338563-5903-4f63-b989-12a93f9d0716` -> select todo with given ID
    - `todogo select` -> open `fzf` and let user select a todo
- [ ] Consider creating a distinct package for the `update` command since it has subcommands
    - [Video that explains how to do this](https://www.youtube.com/watch?v=SSRIn5DAmyw&t=479s)
    - [Cobra user guide's section on how to organize subcommands](https://github.com/spf13/cobra/blob/4cafa37bc4bb85633b4245aa118280fe5a9edcd5/site/content/user_guide.md#organizing-subcommands)
