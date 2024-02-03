# Palworld-Appearance-Tool
Tool for appearance manipulation for Palworld written in Go.

## Setup
- Download and put [uesave](https://github.com/trumank/uesave-rs/releases/latest) binary in PAT's folder.

## Usage
Saves are stored here: `%LOCALAPPDATA%\Pal\Saved\SaveGames\<ID>\Players`

Export JSON appearance data:   
`pat.exe export -i 1.sav -o out.json`

Open the JSON in your favourite text editor and make your changes (see options.txt for options).

Import JSON appearance data to save:   
`pat.exe import -i out.json -o 1.sav`

## Thank you
- trumank for their great uesave-rs. 

## Disclaimer
- I will not be responsible for any possibility of save corruption.
- Palworld brand and name is the registered trademark of its respective owner.
- PAT has no partnership, sponsorship or endorsement with Pocket Pair.
