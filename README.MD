# 🛠️ Sumedh's Dotfiles

This repository contains my personal Linux configuration files, package lists, and setup scripts.  
It allows me to **backup and restore** my entire development environment on any fresh Linux install with minimal effort.  

---

## 📂 What's Inside

- **configs/**
  - `zsh/` → `.zshrc`, `.p10k.zsh`, and Oh My Zsh plugins/themes
  - `ghostty/` → Ghostty terminal configuration
  - `vscode/` → VS Code settings, keybindings, and extensions list
- **packages/**
  - `apt-packages.txt` → List of apt-installed packages
  - `flatpak-list.txt` → List of Flatpak apps
  - `snap-list.txt` → List of Snap apps (if any)
- **fonts/** → Custom fonts
- **extensions/**
  - `gnome-extensions.txt` → Installed GNOME extensions
  - `extensions-settings.conf` → GNOME extension settings
- **backup.sh** → Collects all current configs into this repo
- **install.sh** → Installs packages, restores configs, fonts, and GNOME extensions

---

## 🚀 Usage

### 1. Backup your current setup
```bash
cd ~/dotfiles
./backup.sh
git add .
git commit -m "update configs"
git push
```

### 2. Restore on a fresh install

```bash
# Install Git
sudo apt install git -y

# Clone this repo
git clone https://github.com/sumedhvats/dotfiles.git ~/dotfiles
cd ~/dotfiles

# Run setup
./install.sh
```

## 🐚 Zsh + Oh My Zsh
This repo includes my `.zshrc` and Powerlevel10k theme. For plugins and themes, you can either:
* Store them in this repo (**heavier**), or
* Auto-clone them during setup (**lighter, recommended**):

```bash
git clone https://github.com/marlonrichert/zsh-autocomplete.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autocomplete
git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions
git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting
git clone https://github.com/romkatv/powerlevel10k.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/themes/powerlevel10k
```

## 🧩 GNOME Extensions
Backup & restore extensions:

```bash
# Backup
gnome-extensions list > ./extensions/gnome-extensions.txt
dconf dump /org/gnome/shell/extensions/ > ./extensions/extensions-settings.conf

# Restore
xargs -n1 gnome-extensions enable < ./extensions/gnome-extensions.txt
dconf load /org/gnome/shell/extensions/ < ./extensions/extensions-settings.conf
```

## 📦 Package Reinstall
* **Apt packages**

```bash
sudo apt-get update
xargs sudo apt-get install -y < ./packages/apt-packages.txt
```

* **Flatpak apps**

```bash
xargs -n1 flatpak install -y < ./packages/flatpak-list.txt
```

* **Snap apps**

```bash
snap install <app>
```

## ✅ Workflow Summary
* Run `backup.sh` regularly to keep configs updated.
* Commit & push to GitHub.
* On a new machine, just clone this repo and run `install.sh`.

---

✦ Made with ❤️ by Sumedh Vats
