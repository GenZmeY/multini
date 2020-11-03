#!/bin/bash

source "common.sh"

$Multini --get "$InIni" MultipleKeySection Key > "$OutIni"

compare
