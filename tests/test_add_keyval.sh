#!/bin/bash

source "common.sh"

$Multini -o "$OutIni" --add "$InIni" SimpleSection Key3 3

compare
