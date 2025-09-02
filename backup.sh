#!/bin/bash
set -e

BACKUP_DIR="$HOME/dotfiles"

echo "📦 Backing up apt packages..."
dpkg --get-selections > "$BACKUP_DIR/packages.list"

echo "🐚 Backing up Zsh + Oh My Zsh..."
cp -r ~/.zshrc "$BACKUP_DIR/"
cp -r ~/.oh-my-zsh "$BACKUP_DIR/oh-my-zsh"

echo "🖥️ Backing up Ghostty..."
mkdir -p "$BACKUP_DIR/ghostty"
cp -r ~/.config/ghostty/* "$BACKUP_DIR/ghostty/"

echo "📝 Backing up VS Code..."
mkdir -p "$BACKUP_DIR/vscode"
cp -r ~/.config/Code/User/* "$BACKUP_DIR/vscode/"

echo "🔤 Backing up Fonts..."
mkdir -p "$BACKUP_DIR/fonts"
cp -r ~/.local/share/fonts/* "$BACKUP_DIR/fonts/"

echo "🧩 Backing up GNOME Extensions..."
mkdir -p "$BACKUP_DIR/gnome/extensions"
cp -r ~/.local/share/gnome-shell/extensions/* "$BACKUP_DIR/gnome/extensions/"

echo "=== Backing up GNOME settings (including shortcuts) ==="
mkdir -p "$BACKUP_DIR/gnome"
dconf dump /org/gnome/ > "$BACKUP_DIR/gnome/settings.dconf"

echo "✅ Backup complete! Stored in $BACKUP_DIR"
