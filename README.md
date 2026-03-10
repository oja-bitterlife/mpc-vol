# mpc-vol

A CUI (Character User Interface) tool to manipulate [mpc](https://musicpd.org/clients/mpc/) volume in the terminal.

## Requirements

- Python 3.6+
- [mpc](https://musicpd.org/clients/mpc/) (Music Player Client)
- A running [MPD](https://www.musicpd.org/) (Music Player Daemon) instance

## Installation

```bash
# Clone the repository
git clone https://github.com/oja-bitterlife/mpc-vol.git
cd mpc-vol

# Make the script executable and install it
chmod +x mpc-vol
sudo cp mpc-vol /usr/local/bin/mpc-vol
```

## Usage

```bash
mpc-vol
```

Launch the interactive volume control UI. The current volume is displayed as a
bar and percentage. Use the keyboard controls below to adjust the volume.

### Key Bindings

| Key              | Action                  |
|------------------|-------------------------|
| `↑` / `+` / `k` | Increase volume by 1%  |
| `↓` / `-` / `j` | Decrease volume by 1%  |
| `Page Up`        | Increase volume by 10%  |
| `Page Down`      | Decrease volume by 10%  |
| `q` / `Esc`      | Quit                    |

## Screenshot

```
                          mpc-vol

          Volume:  75%
          [█████████████████████████████████████░░░░░░░░░░░░░]

                    ↑ / +    Increase volume (+1)
                    ↓ / -    Decrease volume (-1)
                    PgUp     Increase volume (+10)
                    PgDn     Decrease volume (-10)
                    q        Quit
```
