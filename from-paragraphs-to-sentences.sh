cat input.txt | go run from-paragraphs-to-sentences.go > myoutput.txt;
diff myoutput.txt output.txt;
cat input2.txt | go run from-paragraphs-to-sentences.go > myoutput2.txt;
diff myoutput2.txt output2.txt;
