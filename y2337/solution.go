package y2337

func canChange(start string, target string) bool {
    j :=0
    i :=0
    for ; i < len(start); i++ {
        if start[i] == 'L' {
            if j == len(target) {
                return false
            }
            for ;j<len(target);j++ {
                if target[j] == '_'   {
                    if j ==  len(target)-1 {
                        return false // target少L
                    }
                    continue
                } else if target[j] == 'L' {
                    if j <= i {
                        j++
                        break
                    } else {
                        return false
                    }
                } else {
                    return false
                }


            }
        } else if start[i] == 'R' {
            if j == len(target) {
                return false
            }
            for ;j<len(target);j++ {
                if target[j] == '_' {
                    if j ==  len(target)-1 {
                        return false  // target少R
                    }
                    continue
                } else if target[j] == 'R' {
                    if j >= i {
                        j++
                        break
                    } else {
                        return false
                    }
                } else {
                    return false
                }
            }

        }

    }

    for _,a := range target[j:] {
        if a != '_' {
            return false
        }
    }
    return  true
}