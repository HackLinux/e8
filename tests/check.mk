.PHONY: check build

check:
	leaf lex main.lf | diff - lex.result

build:
	leaf lex main.lf > lex.result

