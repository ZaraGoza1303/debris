package types

type Config struct {
	Scanner Scanner `toml:"scanner"`
}

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
