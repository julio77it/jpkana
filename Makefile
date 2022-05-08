all: html

html: hiragana booklike

hiragana:
	tools/htmltable.py resources/hiragana.json resources/filters/hiragana.pure.json resources/filters/hiragana.n.json resources/filters/hiragana.dakuten.json resources/filters/hiragana.andakuten.json resources/filters/hiragana.y.json resources/filters/hiragana.double.json > html/hiragana.html

booklike:
	tools/htmltable.py resources/hiragana.json resources/filters/hiragana.pg1.json > html/hiragana.book.pg1.html
	tools/htmltable.py resources/hiragana.json resources/filters/hiragana.pg2.json > html/hiragana.book.pg2.html
