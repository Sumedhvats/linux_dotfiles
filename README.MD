# 🛠️ Sumedh's Linux Dotfiles

Personal Linux configuration files, package lists, and setup scripts.  
**Backup and restore** your entire dev environment on any fresh Ubuntu install with minimal effort.

> Tested on **Ubuntu 26.04 LTS** (Resolute Raccoon)

---

## 📂 What's Inside

```
linux_dotfiles/
├── configs/
│   ├── zsh/              # .zshrc, .p10k.zsh, custom plugins/themes
│   ├── ghostty/          # Ghostty terminal config
│   └── vscode/           # VS Code settings, keybindings, extensions list
├── packages/
│   ├── apt-packages.txt  # Manually installed apt packages
│   └── flatpak-list.txt  # Flatpak apps
├── fonts/                # FiraCode Nerd Font, JetBrains Mono, MesloLGS
├── extensions/
│   ├── gnome-extensions.txt       # Installed GNOME extensions
│   └── extensions-settings.conf   # GNOME extension settings (dconf)
├── gnome/
│   └── gnome-settings.dconf       # Full GNOME settings dump (shortcuts, etc.)
├── backup.sh             # Saves current configs into this repo
└── install.sh            # Restores everything on a fresh install
```

---

## 🚀 Usage

### 1. Backup your current setup

```bash
cd ~/linux_dotfiles
./backup.sh
git add .
git commit -m "update configs"
git push
```

### 2. Restore on a fresh Ubuntu install

```bash
# Install Git & curl
sudo apt install git curl -y

# Clone this repo
git clone https://github.com/Sumedhvats/linux_dotfiles.git ~/linux_dotfiles
cd ~/linux_dotfiles

# Run setup
chmod +x install.sh
./install.sh
```

> Log out and back in (or reboot) after running `install.sh` for Zsh and GNOME changes to take effect.

---

## 🔧 What `install.sh` Does

1. **Installs apt packages** from `packages/apt-packages.txt`
2. **Sets up Flatpak** with Flathub remote and installs saved apps
3. **Installs Zsh + Oh My Zsh** with plugins and Powerlevel10k theme
4. **Installs Ghostty** terminal (uses PPA on Ubuntu < 26.04, official repos on 26.04+)
5. **Restores VS Code** settings and extensions
6. **Installs fonts** (FiraCode Nerd Font, JetBrains Mono)
7. **Enables GNOME extensions** and restores extension + desktop settings
8. **Sets Zsh** as default shell

---

## 🐚 Zsh + Oh My Zsh

`install.sh` automatically installs Oh My Zsh and clones these plugins/themes:

| Plugin / Theme | Source |
|---|---|
| zsh-autosuggestions | [zsh-users/zsh-autosuggestions](https://github.com/zsh-users/zsh-autosuggestions) |
| zsh-autocomplete | [marlonrichert/zsh-autocomplete](https://github.com/marlonrichert/zsh-autocomplete) |
| zsh-syntax-highlighting | [zsh-users/zsh-syntax-highlighting](https://github.com/zsh-users/zsh-syntax-highlighting) |
| Powerlevel10k | [romkatv/powerlevel10k](https://github.com/romkatv/powerlevel10k) |

To clone them manually instead:

```bash
git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions
git clone https://github.com/marlonrichert/zsh-autocomplete.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autocomplete
git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting
git clone --depth=1 https://github.com/romkatv/powerlevel10k.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/themes/powerlevel10k
```

---

## 🧩 GNOME Extensions

Backup & restore manually:

```bash
# Backup
gnome-extensions list > ./extensions/gnome-extensions.txt
dconf dump /org/gnome/shell/extensions/ > ./extensions/extensions-settings.conf

# Restore
xargs -n1 gnome-extensions enable < ./extensions/gnome-extensions.txt
dconf load /org/gnome/shell/extensions/ < ./extensions/extensions-settings.conf
```

> **Note:** Extensions must already be installed on the system for `gnome-extensions enable` to work. Install them from [extensions.gnome.org](https://extensions.gnome.org/) or via Extension Manager first.

---

## 📦 Package Management

### Apt packages

`backup.sh` uses `apt-mark showmanual` to save only packages **you** explicitly installed (not auto-installed dependencies). This keeps the list clean and portable across Ubuntu versions.

```bash
# Backup
apt-mark showmanual | sort > ./packages/apt-packages.txt

# Restore
xargs sudo apt install -y < ./packages/apt-packages.txt
```

### Flatpak apps

```bash
# Backup
flatpak list --app --columns=application > ./packages/flatpak-list.txt

# Restore (Flathub remote must be added first)
flatpak remote-add --if-not-exists flathub https://dl.flathub.org/repo/flathub.flatpakrepo
while IFS= read -r app; do flatpak install -y flathub "$app"; done < ./packages/flatpak-list.txt
```

---

## ⚙️ GNOME Desktop Settings

Full GNOME settings (keyboard shortcuts, appearance, etc.) are backed up via `dconf`:

```bash
# Backup
dconf dump /org/gnome/ > ./gnome/gnome-settings.dconf

# Restore
dconf load /org/gnome/ < ./gnome/gnome-settings.dconf
```

---

## ✅ Workflow Summary

1. Run `./backup.sh` regularly to keep configs updated
2. Commit & push to GitHub
3. On a new machine, clone this repo and run `./install.sh`
4. Log out and back in — you're done ✨

---

✦ Made with ❤️ by Sumedh Vats
