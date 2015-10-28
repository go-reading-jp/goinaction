
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
* フィードの種類ごとに異なるMatheが使える

### JSONとstructure

* JSONと、structは相性が良い
* structにタグとして、対応するJSONのフィールド名を定義出来る
* fileをcloseするときは、deferを使って閉じると関数の終了時にじられる
* “Decode(v interface{}) error”なので、引数は空インターフェースなのでgoの型型ならなんでも引数に渡せる。
 実際は6つのデコーダの中で型アサーションにより適切なものが使われる。

#### 参考

* [JSON and Go](http://blog.golang.org/json-and-go)
* [JSONパッケージ](http://golang.jp/pkg/json)

---

data/data.json
```json
[
{
	"site" : "npr",
	"link" : "http://www.npr.org/rss/rss.php?id=1001",
	"type" : "rss"
},
{
	"site" : "npr",
	"link" : "http://www.npr.org/rss/rss.php?id=1008",
	"type" : "rss"
}
]
```

---
search/feed.go
```go
const dataFile = "data/data.json"

// Feed contains information we need to process a feed.
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// RetrieveFeeds reads and unmarshals the feed data file.
func RetrieveFeeds() ([]*Feed, error) {
	// Open the file.
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// Schedule the file to be closed once
	// the function returns.
	defer file.Close()

	// Decode the file into a slice of pointers
	// to Feed values.
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	// We don't need to check for errors, the caller can do this.
	return feeds, err
}
```

---

## match.go/defualt.go

### Writing generic code using interfaces

* インタフェース型の宣言時に指定したメソッドリストのメソッドをすべて実装することで，インタフェースを実装することができ
* 慣習的にInterface型を定義するときはerを末尾につける    
fmt.Stringer、io.Reader
* interface{}型は，メソッドリストがないインタフェース型の型リテラル      
interface{}型の変数や引数には，どんな型の値でも代入したり，渡したりすることができる

#### interfaceの定義の仕方
```go
type <型名> interface {
    メソッド名(引数の型, ...) (返り値の型, ...)
    ・
    ・
}
```

#### 参考資料
* [Go の interface 設計](http://jxck.hatenablog.com/entry/20130325/1364251563)
* [実践Go](http://golang.jp/effective_go#interfaces)
* [How to use interfaces in Go](http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go)

---

#### 型アサーションの例

引数として受け取ったインターフェース型がstringの配列型であればそのままenvsに代入、string型であればstringの配列型に変換してenvsに代入している。

https://github.com/mattn/gom/blob/master/gomfile.go
```go
func matchOS(any interface{}) bool {
    var envs []string
    if as, ok := any.([]string); ok {
        envs = as
    } else if s, ok := any.(string); ok {
        envs = []string{s}
    } else {
        return false
    }
 
    if has(envs, runtime.GOOS) {
        return true
    }
    return false
}
```


---

### ソースコード

#### Interfaceの定義

Matcherのインターフェースを定義する。

search/match.go
```go
// Matcher defines the behavior required by types that want
// to implement a new search type.
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}
```
#### Interfaceの実装と利用

Searchメソッドを実装しているため、Matcherインターフェイスを実装していると言える。

search/defualt.go
```go
// 省略

// defaultMatcher implements the default matcher.
type defaultMatcher struct{}

// init registers the default matcher with the program.
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// Search implements the behavior for the default matcher.
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
```
---
MatchメソッドではMatcherインターフェースを引数として受け取っているため、処理を変更することができる。

```go
// Match is launched as a goroutine for each individual feed to run
// searches concurrently.
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	// Perform the search against the specified matcher.
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	// Write the results to the channel.
	for _, result := range searchResults {
		results <- result
	}
}

```

---



# 疑問

* プロジェクトの構成とか、GOPATHってどうしてます？
