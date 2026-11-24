package types

// Main
type Config struct {
	Scanner Scanner `toml:"scanner"`
	Lookup  Lookup  `toml:"lookup"`
}

// Lookup
type SearchRoot struct {
	Path     string `toml:"path"`
	Category string `toml:"category"`
}

type Lookup struct {
	Paths    []SearchRoot   `toml:"paths"`
	AppImage AppImageLookup `toml:"appimage"`
	Flatpak  FlatpakLookup  `toml:"flatpak"`
}

type AppImageLookup struct {
	Paths []SearchRoot `toml:"paths"`
}

type FlatpakLookup struct {
	Paths []SearchRoot `toml:"paths"`
}

// Scanner
type Scanner struct {
	AppImage AppImage `toml:"appimage"`
	Flatpak  Flatpak  `toml:"flatpak"`
}

type AppImage struct {
	Enabled bool     `toml:"enabled"`
	Paths   []string `toml:"paths"`
}

type Flatpak struct {
	Enabled bool `toml:"enabled"`
}
