package validator

import (
)

/*
 */
type IntChainer interface {
    IsInRange(min, max int)   IntChainer
    IsInList(haystack []int)  IntChainer
    IsGreater(comp int)       IntChainer
    IsLess(comp int)          IntChainer
    IsNegative()              IntChainer
    IsPositive()              IntChainer
    IsNonNegative()           IntChainer
    IsNonPositive()           IntChainer

    ValidateInt(num int)      bool
}

/*
 */
type intChain struct {
    chains  []func() bool
    Num     int
}

/* BuildIntChain creates an empty validator chain to check
 * on int values.
 */
func BuildIntChain() IntChainer {
    return &intChain{}
}

/* IsInRange is a validator that checks the number
 * is between or equal to min and max.
 */
func (v *intChain) IsInRange(min, max int) IntChainer {
    f := func() bool {
        if min > max || min == max {
            return false
        }
        if v.Num >= min && v.Num <= max {
            return true
        }
        
        return false
    }
    v.chains = append(v.chains, f)

    return v
}

/* IsInList is a validator that checks the number
 * is one of the numbers contained in the haystack list.
 */
func (v *intChain) IsInList(haystack []int) IntChainer {
    f := func() bool {
        for _, val := range haystack {
            if val == v.Num {
                return true
            }
        }
        return false
    }
    v.chains = append(v.chains, f)

    return v
}

/* IsGreater is a validator that checks that the number
 * is greater than comp.
 */
func (v *intChain) IsGreater(comp int) IntChainer {
    f := func() bool {        
        return v.Num > comp
    }
    v.chains = append(v.chains, f)

    return v
}

/* IsLess is a validator that checks that the number
 * is less than comp.
 */
func (v *intChain) IsLess(comp int) IntChainer {
    f := func() bool {        
        return v.Num < comp
    }
    v.chains = append(v.chains, f)

    return v
}

/* IsNegative is a validator that checks that the number
 * is a negative number, that is, less than zero.
 */
func (v *intChain) IsNegative() IntChainer {
    f := func() bool {        
        return v.Num < 0
    }
    v.chains = append(v.chains, f)

    return v
}

/* IsPositive is a validator that checks that the number
 * is a postive number, that is, greater than zero.
 */
func (v *intChain) IsPositive() IntChainer {
    f := func() bool {        
        return v.Num > 0
    }
    v.chains = append(v.chains, f)

    return v
}

/* IsNonNegative is a validator that checks that the number
 * is not a negative number, that is, zero or greater.
 */
func (v *intChain) IsNonNegative() IntChainer {
    f := func() bool {        
        return v.Num >= 0
    }
    v.chains = append(v.chains, f)

    return v
}

/* IsNonPositive is a validator that checks that the number
 * is not a positive number, that is, zero or less.
 */
func (v *intChain) IsNonPositive() IntChainer {
    f := func() bool {        
        return v.Num <= 0
    }
    v.chains = append(v.chains, f)

    return v
}

/* ValidateInt runs all of the validators attached to the
 * chain on num. It returns true only if num passes
 * all validators in the chain.
 */
func (v *intChain) ValidateInt(num int) bool {
    res := true
    v.Num = num
    for _, fn := range v.chains {
        res = res && fn()
    }

    return res
}

