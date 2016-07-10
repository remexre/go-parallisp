#!/bin/bash

(
	echo 'package interpreter' &&
	echo &&
	echo '// Prelude contains the prelude -- the standard library functions implemented' &&
	echo '// in-language.' &&
	echo 'const Prelude = `; Begin prelude' &&
	echo &&
	sed 's/`/` + "`" + `/g' prelude.lisp &&
	echo '`';
) > prelude.go;
