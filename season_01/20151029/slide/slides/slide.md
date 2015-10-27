
class: center, middle

# William Kennedy  “Go in Action MEAP V10"

## chapter2

---

# この章で取り上げること


サンプルアプリケーションを作って、以下の概要を学びます。

* Reviewing a comprehensive Go program
* Declaring types, variables, functions, and methods
* Launching and synchronizing goroutines
* Writing generic code using interfaces
* Handling errors as normal program logic

[ソースコード](https://github.com/goinaction/code/tree/master/chapter2/sample)

---
class: center, middle
# MAIN PACKAGE

## main.go

---

### パッケージ
#### package
* main 関数がプログラム（build tools)のエントリーポイントとして必要
* アプリケーションとして実行するには、 package mainの中にmain関数が必要
* パッケージ名はそのソースファイルのディレクトリの基本名
* 全てのコードはパッケージに属している必要がある
* ここでは、名前空間として使えるというレベルの理解でよい。    
詳しくはchapter3で

---
#### import
* importは他の宣言文よりも前に記述する必要がある
* exportされた識別子や、標準ライブラリを利用する際に利用する
* \_をパッケージ名に利用することで、パッケージの識別子を利用しない場合でもコンパイルエラーになることを防ぐことが出来る。

### init関数
* 全てのinit関数は、main関数が呼ばれる前に処理される
* ここでは、標準loggerの出力先を、標準出力に設定している

---
# SEARCH PACKAGE

* ビジネロジックを定義している
* search/search.go のRunメソッドがプログラムのメインcontrol処理
---

## search.go

packageはseachディレクトリの下にいるのでsearchになっている。

---

### Declaring types, variables, functions, and methods

#### スコープ

* 変数を定義する際にはvarをつけて宣言する
* トップレベル(関数外)に宣言されている変数や関数の先頭文字が大文字か小文字かで、グローバルスペースに所属するか、パッケージスペースに所属するかを判断する
* それ以外は基本的にはブロックの中がスコープになる

参考
* [宣言とスコープ@golang.jp](http://golang.jp/go_spec#Declarations_and_scope)
* [go-study@t9md](https://github.com/t9md/go-study#変数のスコープ)


---

#### 変数
* 基本的にはvarキーワードを使って宣言
* 変数の宣言の後に型の指定が必要。     
ただし、宣言と合わせて初期化を行う場合には型を推測してくれるため諸略可能。
* 省略書式（:= ）を使って、変数宣言と初期化を同時に行う事が可能。     
この場合はvarキーワードは不要


#### 関数

* 関数は複数の値を返すことが出来る
* 正常の結果とerrorを返すというように使われることが多い

---

### Launching and synchronizing goroutines

#### groutine

* 前回も出てきたので省略

#### channel

* makeで作成し、ゼロで初期化する
* Channelは参照型で定義され、go rutineで利用される
* goroutineとのデータの通信に利用される
* オペレータ <- を使って goroutine 間で値の送受信ができる
* ロックや、セマフォとして利用することで、排他制御や同期化を、簡単な構文で容易に実現できま

---

#### サンプルコード

[Go by Example: Channel Synchronization](https://gobyexample.com/channel-synchronization)
```go
package main

import "fmt"
import "time"

// This is the function we'll run in a goroutine. The
// `done` channel will be used to notify another
// goroutine that this function's work is done.
func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    // Send a value to notify that we're done.
    done <- true
}

func main() {

    // Start a worker goroutine, giving it the channel to
    // notify on.
    done := make(chan bool, 1)
    go worker(done)

    // Block until we receive a notification from the
    // worker on the channel.
    <-done
}
```
---
#### syncパッケージ

* 待ち合わせ処理として利用する
* Add()でカウンタを上げる、Done()でカウンタを下げる、Wait()ではカウンタが0になるまで待つ

#### サンプルコード

```go
import "sync"

func execLoop(list []Item) {
    var wg sync.WaitGroup
    for _, item := range list {
        wg.Add(1)
        go func(item2 Item) {
            do_something(item2)
            wg.Done()
        }(item)
    }
    wg.Wait()
}
```
---

## feed.go

* search.goのL.14に出てきたRetrieveFeedsの実装を行う
* data.jsonファイルで定義されているフィードを読み込む
* フィード毎に異なるMatherを使うかも？

### Writing generic code using interfaces






# 疑問

* プロジェクトの構成とか、GOPATHってどうしてます？
