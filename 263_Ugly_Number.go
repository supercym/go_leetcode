func isUgly(num int) bool {
    if num <= 0 {
        return false
    } else if num == 1 {
        return true
    }
    for num % 2 == 0 {
        num /= 2
        if num == 1 {
            return true
        }
    } 
    for num % 3 == 0 {
        num /= 3
        if num == 1 {
            return true
        }
    }
    for num % 5 == 0 {
        num /= 5
        if num == 1 {
            return true
        }
    }
    return false
}

