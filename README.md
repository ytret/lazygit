Fork of https://github.com/jesseduffield/lazygit

Changes:

- **Hide nested repositories from the files panel.** Set
  `git.excludeNestedRepos: true` to hide directories containing a `.git`
  entry, including symlinks. Registered submodules and linked worktrees stay
  visible. Disabled by default.
- **Customize search-match colors.** Configure `gui.theme.searchMatchFgColor`,
  `searchMatchBgColor`, and `searchMatchSelectedBgColor` to control the text
  color, match background, and selected-match background. Defaults are black,
  yellow, and cyan, respectively.

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
