#!/bin/zsh

source ~/.zshrc

echo 'generating antlr parser ...'
antlr -Dlanguage=Go -o parser Calc.g4
echo 'generation done'