#!/bin/bash

source "common.sh"

$Multini -o "$OutIni" --add "$InIni" MultipleKeySection Key 4

compare
