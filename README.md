# Debris

Debris is a Linux cleanup tool written in Go. It scans applications installed through your distribution's package manager and other supported sources, uninstalls them through their native source, and helps you find and remove leftover files.

Residual files are never deleted automatically. Debris shows what it found and lets you choose which items to remove.

## Features

- Scan applications installed through your distribution's package manager and other supported sources.
- Uninstall each application through its native source.
- Find residual files in configurable paths, then choose what to remove.
- Interactive terminal UI for selecting applications and residual files.

## Requirements

- Linux. The distribution is detected from `/etc/os-release`.
- Go 1.27.2 or newer.
- `make` (optional, but used by the install commands below).

## Installation

Clone the repository and run `make`:

```bash
git clone https://github.com/ZaraGoza1303/debris.git
cd debris
make
```

`make` generates `~/.config/debris/config.toml`, builds the binary, and installs it to `~/.local/bin/debris`.

Make sure `~/.local/bin` is in your `PATH`:

```bash
export PATH="$HOME/.local/bin:$PATH"
```

To build without installing:

```bash
make generate
make build
./debris -h
```

Or build directly with Go:

```bash
cd src
go build -o ../debris .
```

There are currently no prebuilt release binaries. Build from source.

## Usage

```bash
debris list       # scan and list detected applications
debris uninstall  # uninstall an application interactively
debris -h         # show help
```

There is no separate `scan` command; `list` performs the scan. It prints every application Debris can detect, along with its source and version:

```text
NAME                 SOURCE    VERSION
firefox              apt       128.0
Calculator            flatpak   46.0
Obsidian              AppImage  1.6.7
```

`uninstall` opens an interactive flow:

1. Select an application.
2. Confirm the uninstall.
3. Debris runs the native uninstall command for that source.
4. Debris searches for residual files.
5. Select the residual files to remove.
6. Confirm the removal.

Running `debris` without arguments prints a short help hint.

## Configuration

Debris reads its configuration from:

```text
~/.config/debris/config.toml
```

The file is required. If it is missing, Debris exits with an error. `make generate` creates a default configuration.

Example configuration:

```toml
[scanner.appimage]
enabled = true
paths = [
    "~/Applications",
    "~/AppImages",
    "~/Downloads",
    "~/.local/bin",
    "/opt",
    "/usr/local/bin",
    "/usr/local/share"
]

[scanner.flatpak]
enabled = true

[lookup]
paths = [
    { path = "~/.config", category = "Configuration" },
    { path = "~/.local/share", category = "Application Data" },
    { path = "~/.cache", category = "Cache" },
    { path = "~/.local/state", category = "Application State" }
]

[lookup.flatpak]
paths = [
    { path = "~/.var/app", category = "Flatpak Data" }
]

[lookup.appimage]
paths = [
    { path = "~/.local/share/applications", category = "Desktop Integration" },
    { path = "~/.local/share/icons", category = "Icons" }
]
```

### Options

| Key | Type | Description |
| --- | --- | --- |
| `scanner.appimage.enabled` | boolean | Enable AppImage scanning. |
| `scanner.appimage.paths` | array of strings | Directories to search for AppImage files. |
| `scanner.flatpak.enabled` | boolean | Enable Flatpak scanning. |
| `lookup.paths` | array of tables | Global residual search roots, used for every source. |
| `lookup.paths[].path` | string | Directory to inspect for residual files. |
| `lookup.paths[].category` | string | Label shown in the residual selection UI. |
| `lookup.flatpak.paths` | array of tables | Additional search roots for Flatpak applications. |
| `lookup.appimage.paths` | array of tables | Additional search roots for AppImage applications. |

The same `path` and `category` fields apply to `lookup.flatpak.paths` and `lookup.appimage.paths`.

Notes:

- `~/` and `$HOME` are expanded in scanner and lookup paths.
- `make generate` writes absolute paths using `$HOME` at generation time.
- APT, Pacman, and DNF scanning is automatic based on the detected distribution. There is no enable/disable option for these sources.
- `lookup.flatpak` paths are used only when cleaning a Flatpak application.
- `lookup.appimage` paths are used only when cleaning an AppImage application.
- Both are added to `lookup.paths`, not used as replacements.

## Residual cleanup

After uninstalling an application, Debris searches the configured lookup paths for possible leftover files. Matching entries are shown for review, and you must select and confirm what to remove.

Residual detection is heuristic and based on names. It may miss leftovers or find unrelated files with similar names. Review the paths before confirming.

## Supported sources

| Source | Detected by | Distribution / package manager | Uninstall command | Status |
| --- | --- | --- | --- | --- |
| APT | `dpkg-query` | Ubuntu, Debian, Linux Mint, Pop!_OS | `sudo apt purge -y <package>` | Supported |
| Pacman | `pacman -Qe` | Arch Linux | `sudo pacman -Rns <package>` | Supported |
| DNF | `dnf repoquery --installed` | Fedora | `sudo dnf remove <package>` | Supported |
| Flatpak | `flatpak list --app` | Any distribution with Flatpak | `flatpak uninstall --delete-data <application-id>` | Supported |
| AppImage | Filesystem scan of configured paths | Any Linux distribution | `rm <file>` | Supported |

Other package managers and application sources are not supported yet.

## Limitations

- Automatic distribution detection uses the `ID` field from `/etc/os-release`. Distributions with a different ID, including some derivatives, may not get system package scanning.
- Lookup walks every configured path recursively, so large directories can take longer to scan.
- Residual matching is based on names, not package ownership. It can miss leftovers stored under a different name and can match unrelated files with the same name.
- AppImage uninstall removes only the AppImage file. Desktop entries and icons are removed only if they appear in the residual search results.
- There is no undo. Deleted files cannot be recovered by Debris.
- The application scan runs before the command is handled, so even `debris -h` waits for the scan to finish.

## Contributing

Contributions are welcome.

1. Fork the repository.
2. Create a branch for your change.
3. Run `gofmt` on the Go files.
4. Build with `make build` or `cd src && go build -o ../debris .`.
5. Open a pull request describing the change.

There is no automated test suite or CI yet, so please test your changes manually and include the distribution and source you tested with.

For bug reports, include your distribution, the application source, the command you ran, and the relevant output.

## Credits

Created by [ZaraGoza1303](https://github.com/ZaraGoza1303).

Built with:

- [charm.land/huh](https://github.com/charmbracelet/huh) for the interactive terminal UI.
- [pelletier/go-toml](https://github.com/pelletier/go-toml) for TOML configuration parsing.
- [probonopd/go-appimage](https://github.com/probonopd/go-appimage) for AppImage detection and metadata.