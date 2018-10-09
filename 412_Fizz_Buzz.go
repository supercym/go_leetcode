func fizzBuzz(n int) []string {
    s := make([]string, n)
    for i := 1; i < n+1; i++ {
        s[i-1] = strconv.Itoa(i)
    }
    for i := 3; i < n+1; i += 3 {
        s[i-1] = "Fizz"
    } 
    for i := 5; i < n+1; i += 5 {
        s[i-1] = "Buzz"
    }
    for i := 15; i < n+1; i += 15 {
        s[i-1] = "FizzBuzz"
    }
    return s
}
S