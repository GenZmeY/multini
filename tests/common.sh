#!/bin/bash

ScriptFullname=$(readlink -e "$0")
ScriptName=$(echo "$ScriptFullname" | awk -F '/' '{print $NF;}')
ScriptDir=$(dirname "$ScriptFullname")
Multini="$1"
InDir="$ScriptDir/in_ini"
InIni="$InDir/$ScriptName.ini"
OutDir="$ScriptDir/out_ini"
OutIni="$OutDir/$ScriptName.ini"
ExpectedIni="$ScriptDir/expected_ini/$ScriptName.ini"

if [[ ! -d "$OutDir" ]]; then
	mkdir -p "$OutDir"
fi

rm -f "$OutIni"

function compare ()
{
	if ! cmp -s "$OutIni" "$ExpectedIni" ; then
		diff "$OutIni" "$ExpectedIni"
		return 1
	fi
	rm -f "$OutIni"
	return 0
}
