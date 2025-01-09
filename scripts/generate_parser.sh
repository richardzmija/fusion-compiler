#!/bin/zsh

antlr -Dlanguage=Go -visitor -Xexact-output-dir -o ../internal/parser/ ../grammar/C.g4
