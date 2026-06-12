#!/bin/bash
set -e

# Auto-detect the repo directory (where this script lives)
BACKUP_DIR="$(cd "$(dirname "$0")" && pwd)"

echo "📦 Backing up apt packages (manually installed)..."
mkdir -p "$BACKUP_DIR/packages"
apt-mark showmanual | sort > "$BACKUP_DIR/packages/apt-packages.txt"

echo "📦 Backing up Flatpak apps..."
flatpak list --app --columns=application > "$BACKUP_DIR/packages/flatpak-list.txt" 2>/dev/null || true

echo "🐚 Backing up Zsh + Oh My Zsh..."
mkdir -p "$BACKUP_DIR/configs/zsh"
cp ~/.zshrc "$BACKUP_DIR/configs/zsh/.zshrc"
cp ~/.p10k.zsh "$BACKUP_DIR/configs/zsh/.p10k.zsh" 2>/dev/null || true
if [ -d ~/.oh-my-zsh/custom ]; then
  mkdir -p "$BACKUP_DIR/configs/zsh/custom"
  # Copy custom plugins and themes, but skip git repos inside them
  rsync -a --exclude='.git' ~/.oh-my-zsh/custom/ "$BACKUP_DIR/configs/zsh/custom/"
fi

echo "🖥️ Backing up Ghostty..."
mkdir -p "$BACKUP_DIR/configs/ghostty"
cp -r ~/.config/ghostty/* "$BACKUP_DIR/configs/ghostty/" 2>/dev/null || true

echo "📝 Backing up VS Code..."
mkdir -p "$BACKUP_DIR/configs/vscode/User"
cp ~/.config/Code/User/settings.json "$BACKUP_DIR/configs/vscode/User/" 2>/dev/null || true
cp ~/.config/Code/User/keybindings.json "$BACKUP_DIR/configs/vscode/User/" 2>/dev/null || true
code --list-extensions > "$BACKUP_DIR/configs/vscode/extensions.txt" 2>/dev/null || true

echo "🔤 Backing up Fonts..."
mkdir -p "$BACKUP_DIR/fonts"
cp -r ~/.local/share/fonts/* "$BACKUP_DIR/fonts/" 2>/dev/null || true

echo "🧩 Backing up GNOME Extensions..."
mkdir -p "$BACKUP_DIR/extensions"
gnome-extensions list > "$BACKUP_DIR/extensions/gnome-extensions.txt" 2>/dev/null || true
dconf dump /org/gnome/shell/extensions/ > "$BACKUP_DIR/extensions/extensions-settings.conf" 2>/dev/null || true

echo "⚙️ Backing up GNOME settings (including shortcuts)..."
mkdir -p "$BACKUP_DIR/gnome"
dconf dump /org/gnome/ > "$BACKUP_DIR/gnome/gnome-settings.dconf" 2>/dev/null || true

echo "✅ Backup complete! Stored in $BACKUP_DIR"
