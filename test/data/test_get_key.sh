#!/bin/bash

source "common.sh"

$Multini --get "$InIni" '' DefKey2 > "$OutIni"

compare
