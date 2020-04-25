#!/bin/bash

source "common.sh"

$Multini -o "$OutIni" --del "$InIni" MultipleKeySection Key

compare
