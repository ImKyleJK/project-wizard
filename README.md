# 🧙 project-wizard
> Easily setup a new project with templates, select from hundreds of template projects from frameworks to tech stacks all from the CLI.

## ⚡️ Quick start
First, [download](https://go.dev/dl/) and install Go. Version 1.19 or higher is required.

> Installation is done by using the [`go install`](https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies) command and rename installed binary in `$GOPATH/bin:`

```bash
go install github.com/NotReeceHarris/project-wizard/cmd/pwizard@latest
```

Also, macOS and GNU/Linux users available way to install via Homebrew:
```bash
# Tap a new formula:
brew tap NotReeceHarris/project-wizard

# Installation:
brew install NotReeceHarris/project-wizard/pwizard
```
Let's create a new project via interactive console UI (or CUI for short) in current folder:

```bash
pwizard create
```

That's all you need to know to create a project! 🎉

## ⚙️ Commands & Options
### `create`

CLI command for create a new project with the interactive console UI.

```
cgapp create [OPTION]
```

Option |	Description | Type | Default | Required?
--- | --- | --- | --- | ---
-t | 	Enables to define custom backend and frontend templates. | 	bool |	false |	No

- 📖 Docs: https://github.com/NotReeceHarris/project-wizard/wiki/Command-create
