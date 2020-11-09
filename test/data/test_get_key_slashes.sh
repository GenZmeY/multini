#!/bin/bash

source "common.sh"

$Multini --get "$InIni" 'Slashes/Test' './Dir1/File' > "$OutIni"

compare
