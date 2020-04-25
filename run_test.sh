#!/bin/bash

DEF='\e[0m'; BLD='\e[1m'; RED='\e[31m'; GRN='\e[32m'; WHT='\e[97m'
ScriptFullname=$(readlink -e "$0")
ScriptName=$(echo "$ScriptFullname" | awk -F '/' '{print $NF;}')
ScriptDir=$(dirname "$ScriptFullname")
TestDir="$ScriptDir/tests"
Multini=$(readlink -e "$1")

if [[ -z "$Multini" ]]; then
	Multini="$ScriptDir/multini"
fi

RetCode=0
pushd "$TestDir" > /dev/null
while read Test
do
	echo -ne "${BLD}${WHT}[----]${DEF} $Test"
	Errors=$("$TestDir/$Test" "$Multini")
	if [[ $? -ne 0 ]]; then
		echo -e "\r${BLD}${WHT}[${RED}FAIL${WHT}]${DEF} $Test"
		RetCode=1
		echo "$Errors"
	else
		echo -e "\r${BLD}${WHT}[${GRN} OK ${WHT}]${DEF} $Test"
	fi
done < <(find "$TestDir" -type f -name 'test_*.sh' -printf "%f\n")
popd > /dev/null
exit "$RetCode"
