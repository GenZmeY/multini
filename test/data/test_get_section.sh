#!/bin/bash

source "common.sh"

$Multini --get "$InIni" '' > "$OutIni"

compare
