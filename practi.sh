#!/bin/bash

# get an argument
#   make all lowercase
#   squeeze the spaces
#   convert spaces into dashes
#   add a cpp extension
#
# perhaps think about a system for building sources
# and a system for running targets
# also need a system to tag solutions

TEMPLATE=src/template.cpp
USAGE="Usage: $0 {setup \"<problem name>\" | run \"<problem name>\"}"

if [ "$#" -lt 2 ]; then
  echo $USAGE
fi

case "$1" in
setup)
  shift
  filename=src/$(echo $@ | tr -s " " | tr " " "-").cpp
  cp $TEMPLATE $filename
  echo copied $TEMPLATE into $filename
  echo generating compile_commands.json for better intellisense
  echo "compilation will fail but you can fix that ;) (w/ tags)"
  bear -- make Q=$filename
  ;;
run)
  shift
  set -x
  make Q=src/$@
  ./a.out
  ;;
*)
  shift
  echo $USAGE
  ;;
esac
