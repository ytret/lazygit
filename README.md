Fork of https://github.com/jesseduffield/lazygit

Changes:

1. Hide nested repos from the files panel.
    - Config: `git.excludeNestedRepos: true` (disabled by default).
    - Hides directories containing a directory/symlink named `.git/`, except for
      registered submodules and linked worktrees.
2. Customizable search-match colors.
    - Config `gui.theme.searchMatchFgColor` controls the match text color.
    - Config `gui.theme.searchMatchBgColor` controls the match background color.
    - Config `gui.theme.searchMatchSelectedBgColor` controls the selected-match
      background color.

Example `config.yml`:

```yaml
git:
  excludeNestedRepos: true

gui:
  theme:
    searchMatchFgColor: [black]
    searchMatchBgColor: [yellow]
    searchMatchSelectedBgColor: [cyan]
```
