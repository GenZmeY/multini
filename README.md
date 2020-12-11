# Multini

[![build](https://github.com/GenZmeY/multini/workflows/build/badge.svg)](https://github.com/GenZmeY/multini/actions?query=workflow%3Abuild)
[![tests](https://github.com/GenZmeY/multini/workflows/tests/badge.svg)](https://github.com/GenZmeY/multini/actions?query=workflow%3Atests)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/GenZmeY/multini)](https://golang.org)
[![GitHub](https://img.shields.io/github/license/genzmey/multini)](https://github.com/GenZmeY/multini/blob/master/LICENSE)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/GenZmeY/multini)](https://github.com/GenZmeY/multini/releases)

*Command line utility for manipulating ini files with duplicate key names.*

A compiled version of multini is available on the [release page](https://github.com/GenZmeY/multini/releases).

***

# Description
Some programs use ini file format with duplicate key names.  
For example, these are games based on the [unreal engine](https://en.wikipedia.org/wiki/Unreal_Engine).  
It might look like this (part of the Killing Floor 2 config):  
```
[OnlineSubsystemSteamworks.KFWorkshopSteamworks]
ServerSubscribedWorkshopItems=2267561023
ServerSubscribedWorkshopItems=2085786712
ServerSubscribedWorkshopItems=2222630586
ServerSubscribedWorkshopItems=2146677560
```
Most implementations only support having one property with a given name in a section. If there are several of them, only the first (or last) key will be processed, which is not enough in this case. multini solves this problem.

**note:**  
- multini is case sensitive;
- quotes around the value are not processed (they are part of the value for multini);  
- multi-line values are not supported.  
(but this may change in the future)  

# Build & Install (Manual)
1. Install [golang](https://golang.org), [git](https://git-scm.com/), [make](https://www.gnu.org/software/make/);
2. Clone this repo: `git clone https://github.com/GenZmeY/multini`
3. Go to the source directory: `cd multini`
4. Build: `make`
5. Install: `make install`

# Usage
```
Usage: multini [OPTION]... ACTION ini_file [section] [param] [value]
Actions:
  -g, --get          Get values for a given combination of parameters.
  -s, --set          Set values for a given combination of parameters.
  -a, --add          Add values for a given combination of parameters.
  -d, --del          Delete the given combination of parameters.
  -c, --chk          Display parsing errors for the specified file.

Options:
  -e, --existing     For --set and --del, fail if item is missing.
  -r, --reverse      For --add, adds an item to the top of the section
  -i, --inplace      Lock and write files in place.
                     This is not atomic but has less restrictions
                     than the default replacement method.
  -o, --output FILE  Write output to FILE instead. '-' means stdout
  -u, --unix         Use LF as end of line
  -w, --windows      Use CRLF as end of line
  -q, --quiet        Suppress all normal output
  -h, --help         Write this help to stdout
      --version      Write version to stdout
```

# Examples
**output a global value not in a section:**  
`multini --get ini_file '' param`

**output section:**  
`multini --get ini_file section`

**output list of existing sections:**  
`multini --get ini_file`

**output value:**  
`multini --get ini_file section param`  
- if there are several parameters, a list of all values of this parameter will be displayed

**add/update a single parameter:**  
`multini --set ini_file section parameter value`  
- if there is no parameter, it will be added  
- if the parameter exists, the value will be updated  
- if the parameter exists and has several values, the parameter with the specified value will be set, the rest of the values will be deleted

**add a parameter with specified value:**  
`multini --add ini_file section parameter value`  
- if there is no parameter, it will be added  
- if the parameter exists and does not have the specified value, the new value will be added  
- if the specified value repeats the existing one, no changes will be made

**delete all parameters with specified name:**  
`multini --del ini_file section parameter`

**delete a parameter with specified name and value:**  
`multini --del ini_file section parameter value`

**delete a section:**  
`multini --del ini_file section`

**short options can be combined:**  
`multini -gq ini_file section parameter value`  
- check the existence of a parameter with a given value using the return code

# License
Copyright © 2020 GenZmeY

The content of this repository is licensed under [MIT License](https://github.com/GenZmeY/multini/blob/master/LICENSE).

