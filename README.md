# go-interpreter

「Go 言語でつくるインタプリタ」を読んでインタプリタを作ってみる。

## 1. 字句解析

```go
go test ./monkey/lexer
```

<details>
<summary open>メモ1</summary>
字句解析器にソースコードを与えて初期化し、繰り返し NextToken()を呼ぶことでソースコードを読み進めていく。
トークンごとに、文字ごとに進んでいく。
</details>

<br />

<details>
<summary open>メモ2</summary>
字句解析器の仕事は、コードが意味をなすか、動作するか、エラーを含むかを判定することではないから。
字句解析器には入力をトークン列に変換することだけが求められる。
</details>

<br />

```go
go run ./monkey/main.go
```

<details>
<summary open>メモ3</summary>
REPLは「Read(読み込み)、Eval(評価)、Print(表示)、Loop(繰り返し)」の略.
REPLは入力を読み込んで、インタプリタに送って評価させ、インタプリタの結果/出力を表示して、また最初に戻る。読み込み、評価、表示、繰り返し。
</details>

## 2. 構文解析

<details>
<summary open>メモ1</summary>
ソースコードを入力として(テキストまたはトークン列として)受け取り、ソースコードを表現するようなあるデータ構造を生成する。
そのデータ構造を構築する間には、必然的に入力を解析することになる。その間、入力が期待された構造に従っているかをチェックする。
だから、構文解析という。
</details>

<br />

<details>
<summary open>メモ2</summary>

- トップダウン構文解析
  - 「再帰下降構文解析(recursive descent parsing)」
  - 「アーリー法(Earley parsing)」
  - 「予測的構文解析(predictiveparsing)」
- ボトムアップ構文解析。
</details>

## 3. 評価

## 4. インタプリタの拡張

### Copilot

```
// q: 次の関数の要約をしてください。
// a:
```
