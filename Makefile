all: html

html: hiragana booklike

hiragana:
	tools/htmltable.py resources/hiragana.json resources/filters/pure.json resources/filters/n.json resources/filters/dakuten.json resources/filters/andakuten.json resources/filters/y.json resources/filters/double.json > html/hiragana.html
	tools/htmltable.py resources/katakana.json resources/filters/pure.json resources/filters/n.json resources/filters/dakuten.json resources/filters/andakuten.json resources/filters/y.json resources/filters/double.json > html/katakana.html

booklike:
	tools/htmltable.py resources/hiragana.json resources/filters/book.pg1.json > html/hiragana.book.pg1.html
	tools/htmltable.py resources/hiragana.json resources/filters/book.pg2.json > html/hiragana.book.pg2.html
	tools/htmltable.py resources/katakana.json resources/filters/book.pg1.json > html/katakana.book.pg1.html
	tools/htmltable.py resources/katakana.json resources/filters/book.pg2.json > html/katakana.book.pg2.html
	tools/htmltable.py resources/katakana.json resources/filters/book.pg3.json > html/katakana.book.pg3.html
