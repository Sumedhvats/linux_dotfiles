#!/usr/bin/env bash
set -e

echo "📦 Backing up packages..."
mkdir -p ./packages
dpkg --get-selections > ./packages/apt-packages.txt
flatpak list --app --columns=application > ./packages/flatpak-list.txt
snap list > ./packages/snap-list.txt || true   # skip if no snap

echo "🐚 Backing up Zsh + Oh My Zsh..."
mkdir -p ./configs/zsh
cp ~/.zshrc ./configs/zsh/ 2>/dev/null || true
cp ~/.p10k.zsh ./configs/zsh/ 2>/dev/null || true
cp -r ~/.oh-my-zsh/custom ./configs/zsh/ 2>/dev/null || true

echo "🖥️ Backing up Ghostty..."
mkdir -p ./configs/ghostty
cp -r ~/.config/ghostty/* ./configs/ghostty/ 2>/dev/null || true

echo "📝 Backing up VS Code..."
mkdir -p ./configs/vscode/User
cp -r ~/.config/Code/User/* ./configs/vscode/User/ 2>/dev/null || true
code --list-extensions > ./configs/vscode/extensions.txt || true

echo "🔤 Backing up Fonts..."
mkdir -p ./fonts
cp -r ~/.local/share/fonts/* ./fonts/ 2>/dev/null || true

echo "🧩 Backing up GNOME Extensions..."
mkdir -p ./extensions
gnome-extensions list > ./extensions/gnome-extensions.txt || true
dconf dump /org/gnome/shell/extensions/ > ./extensions/extensions-settings.conf || true

echo "✅ Backup complete! Don’t forget to git add/commit/push"
