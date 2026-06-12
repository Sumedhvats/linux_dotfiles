#!/usr/bin/env bash
set -e

# Auto-detect the repo directory (where this script lives)
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"

# ──────────────────────────────────────────────
# 1. APT PACKAGES
# ──────────────────────────────────────────────
echo "🔧 Installing apt packages..."
sudo apt update
if [ -f "$SCRIPT_DIR/packages/apt-packages.txt" ]; then
  # Install packages, skip any that aren't available (e.g., from missing PPAs)
  xargs -a "$SCRIPT_DIR/packages/apt-packages.txt" sudo apt install -y 2>&1 | tail -5 || true
fi

# ──────────────────────────────────────────────
# 2. FLATPAK + FLATHUB
# ──────────────────────────────────────────────
echo "📦 Setting up Flatpak..."
sudo apt install -y flatpak 2>/dev/null || true
flatpak remote-add --if-not-exists flathub https://dl.flathub.org/repo/flathub.flatpakrepo 2>/dev/null || true

if [ -f "$SCRIPT_DIR/packages/flatpak-list.txt" ]; then
  echo "📦 Installing Flatpak apps..."
  while IFS= read -r app; do
    [ -z "$app" ] && continue
    flatpak install -y flathub "$app" 2>/dev/null || echo "  ⚠️  Skipped: $app"
  done < "$SCRIPT_DIR/packages/flatpak-list.txt"
fi

# ──────────────────────────────────────────────
# 3. ZSH + OH MY ZSH + PLUGINS
# ──────────────────────────────────────────────
echo "🐚 Setting up Zsh & Oh My Zsh..."

# Install Zsh if not present
if ! command -v zsh &>/dev/null; then
  sudo apt install -y zsh
fi

# Install Oh My Zsh if not present
if [ ! -d "$HOME/.oh-my-zsh" ]; then
  echo "  Installing Oh My Zsh..."
  sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" "" --unattended
fi

# Clone Zsh plugins & theme
ZSH_CUSTOM="${ZSH_CUSTOM:-$HOME/.oh-my-zsh/custom}"

if [ ! -d "$ZSH_CUSTOM/plugins/zsh-autosuggestions" ]; then
  git clone https://github.com/zsh-users/zsh-autosuggestions "$ZSH_CUSTOM/plugins/zsh-autosuggestions"
fi

if [ ! -d "$ZSH_CUSTOM/plugins/zsh-autocomplete" ]; then
  git clone https://github.com/marlonrichert/zsh-autocomplete.git "$ZSH_CUSTOM/plugins/zsh-autocomplete"
fi

if [ ! -d "$ZSH_CUSTOM/plugins/zsh-syntax-highlighting" ]; then
  git clone https://github.com/zsh-users/zsh-syntax-highlighting.git "$ZSH_CUSTOM/plugins/zsh-syntax-highlighting"
fi

if [ ! -d "$ZSH_CUSTOM/themes/powerlevel10k" ]; then
  git clone --depth=1 https://github.com/romkatv/powerlevel10k.git "$ZSH_CUSTOM/themes/powerlevel10k"
fi

# Copy Zsh configs
cp "$SCRIPT_DIR/configs/zsh/.zshrc" ~/
cp "$SCRIPT_DIR/configs/zsh/.p10k.zsh" ~/ 2>/dev/null || true
if [ -d "$SCRIPT_DIR/configs/zsh/custom" ]; then
  cp -r "$SCRIPT_DIR/configs/zsh/custom/"* "$ZSH_CUSTOM/" 2>/dev/null || true
fi

# Set Zsh as default shell
if [ "$SHELL" != "$(which zsh)" ]; then
  echo "  Setting Zsh as default shell..."
  chsh -s "$(which zsh)"
fi

# ──────────────────────────────────────────────
# 4. GHOSTTY TERMINAL
# ──────────────────────────────────────────────
echo "🖥️ Setting up Ghostty..."

# Ghostty is in official repos on Ubuntu 26.04+, needs PPA on older versions
if ! command -v ghostty &>/dev/null; then
  UBUNTU_VERSION=$(lsb_release -rs 2>/dev/null || echo "0")
  if dpkg --compare-versions "$UBUNTU_VERSION" "lt" "26.04" 2>/dev/null; then
    echo "  Adding Ghostty PPA (Ubuntu < 26.04)..."
    sudo add-apt-repository -y ppa:mkasberg/ghostty-ubuntu 2>/dev/null || true
    sudo apt update
  fi
  sudo apt install -y ghostty 2>/dev/null || echo "  ⚠️  Ghostty install skipped (may need manual setup)"
fi

mkdir -p ~/.config/ghostty
cp -r "$SCRIPT_DIR/configs/ghostty/"* ~/.config/ghostty/ 2>/dev/null || true

# ──────────────────────────────────────────────
# 5. VS CODE
# ──────────────────────────────────────────────
echo "📝 Restoring VS Code settings..."
mkdir -p ~/.config/Code/User
cp -r "$SCRIPT_DIR/configs/vscode/User/"* ~/.config/Code/User/ 2>/dev/null || true

if command -v code &>/dev/null && [ -f "$SCRIPT_DIR/configs/vscode/extensions.txt" ]; then
  echo "  Installing VS Code extensions..."
  xargs -n1 code --install-extension < "$SCRIPT_DIR/configs/vscode/extensions.txt" 2>/dev/null || true
fi

# ──────────────────────────────────────────────
# 6. FONTS
# ──────────────────────────────────────────────
echo "🔤 Installing fonts..."
mkdir -p ~/.local/share/fonts
cp -r "$SCRIPT_DIR/fonts/"* ~/.local/share/fonts/ 2>/dev/null || true
fc-cache -fv 2>/dev/null || true

# ──────────────────────────────────────────────
# 7. GNOME EXTENSIONS & SETTINGS
# ──────────────────────────────────────────────
echo "🧩 Restoring GNOME extensions..."
if [ -f "$SCRIPT_DIR/extensions/gnome-extensions.txt" ]; then
  xargs -n1 gnome-extensions enable < "$SCRIPT_DIR/extensions/gnome-extensions.txt" 2>/dev/null || true
fi

if [ -f "$SCRIPT_DIR/extensions/extensions-settings.conf" ]; then
  dconf load /org/gnome/shell/extensions/ < "$SCRIPT_DIR/extensions/extensions-settings.conf"
fi

echo "⚙️ Restoring GNOME settings (including shortcuts)..."
if [ -f "$SCRIPT_DIR/gnome/gnome-settings.dconf" ]; then
  dconf load /org/gnome/ < "$SCRIPT_DIR/gnome/gnome-settings.dconf"
fi

echo ""
echo "✅ Setup complete!"
echo "   ➜ Log out and back in (or reboot) for Zsh and GNOME changes to take effect."
