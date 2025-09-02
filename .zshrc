# Enable Powerlevel10k instant prompt. Should stay close to the top of ~/.zshrc.
# Initialization code that may require console input (password prompts, [y/n]
# confirmations, etc.) must go above this block; everything else may go below.
if [[ -r "${XDG_CACHE_HOME:-$HOME/.cache}/p10k-instant-prompt-${(%):-%n}.zsh" ]]; then
  source "${XDG_CACHE_HOME:-$HOME/.cache}/p10k-instant-prompt-${(%):-%n}.zsh"
fi

# Set Oh My Zsh installation path
export ZSH="$HOME/.oh-my-zsh"
export PATH=$PATH:/usr/local/go/bin

# Load zsh-autocomplete early for circular autocomplete behavior
source ${ZSH_CUSTOM:-$HOME/.oh-my-zsh/custom}/plugins/zsh-autocomplete/zsh-autocomplete.plugin.zsh

# Set Powerlevel10k as the theme
ZSH_THEME="powerlevel10k/powerlevel10k"

# Load fzf if installed
[ -f ~/.fzf.zsh ] && source ~/.fzf.zsh

# Plugins to load
plugins=(
  git
  zsh-autosuggestions
  zsh-autocomplete
  zsh-syntax-highlighting
)

# Source Oh My Zsh to load plugins and theme
source $ZSH/oh-my-zsh.sh

# Source Powerlevel10k configuration if it exists
[[ ! -f ~/.p10k.zsh ]] || source ~/.p10k.zsh
export PATH="$HOME/.local/bin:$PATH"
export PATH="$HOME/go/bin:$PATH"

# --- Keybindings (only in interactive shells) ---
if [[ -o interactive ]]; then
  # Backspace vs Ctrl+Backspace handling
  bindkey '^?' backward-delete-char          # normal Backspace deletes char
  bindkey '^H' backward-kill-word            # Ctrl+Backspace (if it sends ^H)
  bindkey '^[[127;5u' backward-kill-word     # Ctrl+Backspace (if it sends this)
fi

alias e='exit'
alias cc='clear'
