package core

func Choose(condition bool, a interface{}, b interface{}) interface{} {
    if condition {
        return a
    } else {
        return b
    }
}
