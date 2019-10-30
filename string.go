package validator

import (
    "strings"
    "regexp"
)

const (
    STR_ALPHA_UPPER  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    STR_ALPHA_LOWER  = "abcdefghijklmnopqrstuvwxyz"
    STR_NUM          = "0123456789"
)

/*
 */
type StringChainer interface {
    IsMaxLen(limit int)          StringChainer
    IsInList(haystack []string)  StringChainer
    IsUpper()                    StringChainer
    IsLower()                    StringChainer
    IsUpperFirstOnce()           StringChainer
    IsLowerFirstOnce()           StringChainer
    IsUpperFirstAll()            StringChainer
    IsLowerFirstAll()            StringChainer
    IsContains(substr string)    StringChainer

    ValidateStr(text string)     bool
}

/*
 */
type stringChain struct {
    chains  []func() bool
    Text    string
}

/* BuildStrChain creates an empty validator chain to check
 * on string values.
 */
func BuildStrChain() StringChainer {
    return &stringChain{}
}

/* IsMaxLen is a validator that checks the string value
 * does not exceed limit in length.
 */
func (v *stringChain) IsMaxLen(limit int) StringChainer {
    f := func() bool {
        if len(v.Text) > limit {
            return false
        }
        return true
    }
    v.chains = append(v.chains, f)

    return v
}

/* IsInList is a validator that checks the string value
 * is one of the values contained in the haystack list.
 */
func (v *stringChain) IsInList(haystack []string) StringChainer {
    f := func() bool {
        for _, val := range haystack {
            if val == v.Text {
                return true
            }
        }
        return false
    }
    v.chains = append(v.chains, f)

    return v
}

/* IsUpper is a validator that checks that all 
 * alphabetical characters are uppercase.
 */
func (v *stringChain) IsUpper() StringChainer {
    f := func() bool {
        if strings.IndexAny(v.Text, STR_ALPHA_LOWER) == -1 {
            return true
        }
        return false
    }
    v.chains = append(v.chains, f)

    return v
}

/* IsLower is a validator that checks that all 
 * alphabetical characters are lowercase.
 */
func (v *stringChain) IsLower() StringChainer {
    f := func() bool {
        if strings.IndexAny(v.Text, STR_ALPHA_UPPER) == -1 {
            return true
        }
        return false
    }
    v.chains = append(v.chains, f)

    return v
}

/* IsUpperFirstOnce is a validator that checks that 
 * the first alpha character is uppercase.
 */
func (v *stringChain) IsUpperFirstOnce() StringChainer {
    f := func() bool {
        idx := strings.IndexAny(v.Text, STR_ALPHA_UPPER + STR_ALPHA_LOWER)
        if idx == -1 {
            return false
        }
        return strings.ContainsAny(string(v.Text[idx]), STR_ALPHA_UPPER)
    }
    v.chains = append(v.chains, f)

    return v
}

/* IsLowerFirstOnce is a validator that checks that 
 * the first alpha character is lowercase.
 */
func (v *stringChain) IsLowerFirstOnce() StringChainer {
    f := func() bool {
        idx := strings.IndexAny(v.Text, STR_ALPHA_UPPER + STR_ALPHA_LOWER)
        if idx == -1 {
            return false
        }
        return strings.ContainsAny(string(v.Text[idx]), STR_ALPHA_LOWER)
    }
    v.chains = append(v.chains, f)

    return v
}

/* IsUpperFirstAll is a validator that checks that 
 * the first character of every word is uppercase,
 * for example, "I Am Spartacus".
 * Words are separated by one or more whitespaces.
 */
func (v *stringChain) IsUpperFirstAll() StringChainer {
    f := func() bool {
        b := BuildStrChain().IsUpperFirstOnce()
        s := regexp.MustCompile(`\s+`).Split(v.Text, -1)
        for _, word := range s {
            if word != "" && b.ValidateStr(word) == false {
                return false
            }
        }
        return true
    }
    v.chains = append(v.chains, f)

    return v
}

/* IsLowerFirstAll is a validator that checks that 
 * the first character of every word is lowercase,
 * for example, "a am spartacus".
 * Words are separated by one or more whitespaces.
 */
func (v *stringChain) IsLowerFirstAll() StringChainer {
    f := func() bool {
        b := BuildStrChain().IsLowerFirstOnce()
        s := regexp.MustCompile(`\s+`).Split(v.Text, -1)
        for _, word := range s {
            if word != "" && b.ValidateStr(word) == false {
                return false
            }
        }
        return true
    }
    v.chains = append(v.chains, f)

    return v
}

/* IsContains is a validator that checks that substr is 
 * part of the string value.
 * It is a straight implementation of strings.Contains().
 */
func (v *stringChain) IsContains(substr string) StringChainer {
    f := func() bool {
        return strings.Contains(v.Text, substr)
    }
    v.chains = append(v.chains, f)

    return v
}

/* ValidateStr runs all of the validators attached to the
 * chain on text. It returns true only if text passes
 * all validators in the chain.
 */
func (v *stringChain) ValidateStr(text string) bool {
    res := true
    v.Text = text
    for _, fn := range v.chains {
        res = res && fn()
    }

    return res
}

