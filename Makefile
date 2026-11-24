APP := debris

CONFIG_DIR  := $(HOME)/.config/$(APP)
CONFIG_FILE := $(CONFIG_DIR)/config.toml
BINARY_PATH := $(HOME)/.local/bin/$(APP)

define CONFIG_CONTENT
[scanner.appimage]
enabled = true
paths = [
  "$(HOME)/Applications",
  "$(HOME)/AppImages",
  "$(HOME)/Downloads",
  "$(HOME)/.local/bin",
  "/opt",
  "/usr/local/bin",
  "/usr/local/share"
]

[scanner.flatpak]
enabled = true

[lookup]
paths = [
  { path = "$(HOME)/.config", category = "Configuration" },
  { path = "$(HOME)/.local/share", category = "Application Data" },
  { path = "$(HOME)/.cache", category = "Cache" },
  { path = "$(HOME)/.local/state", category = "Application State" }
]

[lookup.flatpak]
paths = [
  { path = "$(HOME)/.var/app", category = "Flatpak Data" }
]

[lookup.appimage]
paths = [
  { path = "$(HOME)/.local/share/applications", category = "Desktop Integration" },
  { path = "$(HOME)/.local/share/icons", category = "Icons" }
]
endef
export CONFIG_CONTENT

.PHONY: all build generate install uninstall

all: install

build:
	cd src && go build -o ../$(APP) .

generate:
	@mkdir -p $(CONFIG_DIR)
	@printf '%s\n' "$$CONFIG_CONTENT" > $(CONFIG_FILE)
	@echo "Created $(APP) config file in $(CONFIG_FILE)"

install: generate build
	install -Dm755 $(APP) $(BINARY_PATH)
	@echo "Installed $(APP) in $(BINARY_PATH)"

uninstall:
	rm -f $(BINARY_PATH)