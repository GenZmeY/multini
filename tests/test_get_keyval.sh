#!/bin/bash

source "common.sh"

$Multini --get "$InIni" MultipleKeySection Key 3 > "$OutIni"

compare
