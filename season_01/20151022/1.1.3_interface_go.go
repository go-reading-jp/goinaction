type Animal interface {
    Speak()
}

type Cat struct {} 
func (_ *Cat) Speak(){
    // にゃー
}
