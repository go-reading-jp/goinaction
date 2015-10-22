func log(msg string) {
    // ロギング処理
}

func main(){
    err := doSomething()
    if err != nil {
        // 別のゴルーチンで実行される
        go log("エラー発生")
    }
    
    // 処理の続き
    ...
}
