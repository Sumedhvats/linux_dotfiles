#!/usr/bin/env bash
set -e

echo "🔧 Installing apt packages..."
sudo apt-get update
xargs sudo apt-get install -y < ./packages/apt-packages.txt

echo "📦 Installing flatpak apps..."
xargs -n1 flatpak install -y < ./packages/flatpak-list.txt

echo "🐚 Setting up Zsh & Oh My Zsh..."
cp ./configs/zsh/.zshrc ~/
cp ./configs/zsh/.p10k.zsh ~/
mkdir -p ~/.oh-my-zsh/custom
cp -r ./configs/zsh/custom/* ~/.oh-my-zsh/custom/ 2>/dev/null || true

echo "🖥️ Setting up Ghostty..."
mkdir -p ~/.config/ghostty
cp -r ./configs/ghostty/* ~/.config/ghostty/

echo "📝 Restoring VS Code..."
mkdir -p ~/.config/Code/User
cp -r ./configs/vscode/User/* ~/.config/Code/User/
xargs -n1 code --install-extension < ./configs/vscode/extensions.txt

echo "🔤 Installing fonts..."
mkdir -p ~/.local/share/fonts
cp -r ./fonts/* ~/.local/share/fonts/
fc-cache -fv

echo "🧩 Installing GNOME extensions..."
xargs -n1 gnome-extensions enable < ./extensions/gnome-extensions.txt
dconf load /org/gnome/shell/extensions/ < ./extensions/extensions-settings.conf

echo "✅ Setup complete!"
