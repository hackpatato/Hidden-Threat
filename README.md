Hidden-Threat

  [![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com)
  [![LANG Go](https://img.shields.io/badge/Lang-Go-blue.svg)](http://makeapullrequest.com)
  [![version](https://img.shields.io/badge/version:-v0.1.beta-white.svg)](http://makeapullrequest.com)


------
##  Overview

**Hidden Threat Description: A lightweight, minimal Proof-of-Concept (PoC) dropper developed for Red Teaming scenarios, payload delivery testing, and malware analysis research.**
------
<div align="center">

  <img src="logo.png" alt="Hidden-Threat Logo" width="360">


Compilation
""

# Clone the repository
git clone https://github.com/hackpatato/Hidden-Threat.git

# Navigate to the project directory
cd Hidden-Threat

# Compile the project
# FOR Ubuntu/Debian :
sudo apt install zig
# FOR Arch Linux :
sudo pacman -S zig
# AND 
CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC="zig cc -target x86_64-windows-gnu" CXX="zig c++ -target x86_64-windows-gnu" go build -o TEST.exe .
""
