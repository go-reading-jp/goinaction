
# この章で取り上げること

サンプルアプリケーションを作って、以下の概要を学びます。

* Reviewing a comprehensive Go program
* Declaring types, variables, functions, and methods
* Launching and synchronizing goroutines
* Writing generic code using interfaces
* Handling errors as normal program logic

[ソースコード](https://github.com/goinaction/code/tree/master/chapter2/sample)

抜粋：: William Kennedy  “Go in Action MEAP V10”。 iBooks  )

# MAIN PACKAGE
## main.go
### パッケージ
#### package
* main 関数がプログラム（build tools)のエントリーポイントとして必要
* アプリケーションとして実行するには、 package mainの中にmain関数が必要
* パッケージ名はそのソースファイルのディレクトリの基本名
* 全てのコードはパッケージに属している必要がある
* ここでは、名前空間として使えるというレベルの理解でよい。詳しくはchapter3で

#### import
* importは他の宣言文よりも前に記述する必要がある
* exportされた識別子や、標準ライブラリを利用する際に利用する
* \_をパッケージ名に利用することで、パッケージの識別子を利用しない場合でもコンパイルエラーになることを防ぐことが出来る。

### init関数
* 全てのinit関数は、main関数が呼ばれる前に処理される
* ここでは、標準loggerの出力先を、標準出力に設定している


# SEARCH PACKAGE

* ビジネロジックを定義している
* search/search.go のRunメソッドがプログラムのメインcontrol処理

## search.go

packageはseachディレクトリの下にいるのでsearchになっている。

### Declaring types, variables, functions, and methods

#### スコープ

* 変数を定義する際にはvarをつけて宣言する
* トップレベル(関数外)に宣言されている変数や関数の先頭文字が大文字か小文字かで、グローバルスペースに所属するか、パッケージスペースに所属するかを判断する
* それ以外は基本的にはブロックの中がスコープになる

#### 変数
* 基本的にはvarキーワードを使って宣言
* 変数の宣言の後に型の指定が必要。ただし、宣言と合わせて初期化を行う場合には型を推測してくれるため諸略可能。
* 省略書式（:= ）を使って、変数宣言と初期化を同時に行う事が可能。この場合はvarキーワードは不要


#### 関数

* 関数は複数の値を返すことが出来る
* 正常の結果とerrorを返すというように使われることが多い

参考
* [宣言とスコープ@golang.jp](http://golang.jp/go_spec#Declarations_and_scope)
* [go-study@t9md](https://github.com/t9md/go-study#変数のスコープ)


### Launching and synchronizing goroutines


## Writing generic code using interfaces

# 疑問

* プロジェクトの構成とか、GOPATHってどうしてます？
