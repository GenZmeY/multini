#!/bin/bash

source "common.sh"

$Multini -o "$OutIni" --del "$InIni" SectionWithIndent

compare
