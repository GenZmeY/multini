#!/bin/bash

source "common.sh"

$Multini -o "$OutIni" --add "$InIni" NewSection

compare
