package main

//标识符：
//是用户或系统定义的有意义单词组合,或单词与数字组合(具体意义有定义者决定)
//标识符以字母下划线开头，大小写敏感，比如：boy,  Boy,  _boy,  _（匿名变量，用来忽略结果）
//标识符命名规范：在习惯上，Go语言程序员推荐使用驼峰式命名，当名字有几个单词组成的时优先使用大小写分隔，而不是优先用下划线分隔。
//因此，在标准库有QuoteRuneToASCII和parseRequestLine这样的函数命名，但是一般不会用quote_rune_to_ASCII和parse_request_line这样的命名。
//而像ASCII和HTML这样的缩略词则避免使用大小写混合的写法，它们可能被称为htmlEscape、HTMLEscape或escapeHTML，但不会是escapeHtml。

//关键字：
//是 Go 语言提供的有特殊含义的符号，也叫做“保留字”
//系统保留关键字：
/*
break	default	func	interface	select
case	defer	go	map	struct
chan	else	goto	package	switch
const	fallthough	if	range	type
continue	for	import	return	var
*/
