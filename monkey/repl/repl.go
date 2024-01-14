package repl

import (
	"bufio"
	"fmt"
	"go-interpreter/monkey/lexer"
	"go-interpreter/monkey/token"
	"io"
)

const PROMPT = ">> "

// 標準入力から読み込んだ文字列を、字句解析して、トークンを出力する。
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		// ユーザーからの新しい行の入力を待ちます。何も入力されなければループを抜けます。
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		// ループの最後で tok = l.NextToken() を使用して次のトークンを取得し、次のイテレーションに進みます。
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
