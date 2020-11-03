#!/bin/bash

source "common.sh"

$Multini -o "$OutIni" --del "$InIni" SimpleSection Key1

compare
