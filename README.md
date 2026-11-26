# Pascal

A terminal-first AI coding assistant. Personal fork of
[Crush](https://github.com/charmbracelet/crush) by Charm, used under its license
(see LICENSE.md).

## Build

Requires Go.

    git clone https://github.com/Ahmed9ali/Pascal.git
    cd Pascal
    go build -o pascal .

On Termux: `pkg install golang git ripgrep` first.

## Run

    export GEMINI_API_KEY=your_key   # or another provider's key
    ./pascal

Config lives in `~/.config/pascal/pascal.json`.
