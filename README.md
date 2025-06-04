# GOGENPASS

### GOGENPASS is a terminal-based interactive password generator written in Go. It allows users to create secure, randomized passwords by combining numbers, letters, and symbols through an intuitive command-line interface. The tool uses the pterm library for a visually appealing and interactive user experience.

### Features
- Interactive CLI: User-friendly prompts to customize password generation.
- Customizable Passwords: Choose whether to include numbers, letters, and symbols, and specify their quantities.
- Randomization: Generates random passwords by combining user-specified components in a randomized order.
- Progress Bar: Visual feedback during password generation for an engaging experience.
- Colorful UI: Leverages pterm for styled terminal output with colors and big text.

### Prerequisites

To run GOGENPASS, you need:
- Go (version 1.16 or higher)
- A terminal or command-line interface
- The following Go modules:
- github.com/pterm/pterm for interactive terminal UI
- Standard Go libraries (fmt, math/rand, strconv, time)
