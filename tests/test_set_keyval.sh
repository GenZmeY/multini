#!/bin/bash

source "common.sh"

$Multini -o "$OutIni" --set "$InIni" MultipleKeySection Key onlyone

compare
