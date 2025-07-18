import "math"

func reverse(x int) int {
    o := len(strconv.Itoa(x)) - 1
    
    if x < 0 {
        o--
    }
    
    var r int

    for i:=o; i>=0; i-- {
        m := pow10Int(i)
        n := pow10Int(o-i)

        a := x / m

        if o==9 && i==0 {
            if a > 2 || a < -2{
                return 0
            }

            if a>0 && (math.MaxInt32 -r) < a * n {
                return 0
            }

            if a<0 && (math.MinInt32 -r) > a * n {
                return 0
            }
        }

        r = r + a * n
        
        x = x - a * m
    }
    
    return r
}

func pow10Int(x int) int {
    return int(math.Pow10(x))
}
