.PHONY:all listen

antlr4=java -Xmx500M -cp "/usr/local/lib/antlr-4.7.2-complete.jar:$CLASSPATH" org.antlr.v4.Tool

all:listen

listen:
	$(antlr4) -Dlanguage=Go -no-listener -visitor -o parser PlayScript.g4
