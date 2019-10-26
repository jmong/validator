package validator

import(
    "testing"
)

/*
 */
func TestIsRange(t *testing.T) {
    testcases := []struct{
        min       int
        max       int
        num       int
        expected  bool
    }{
        {1, 100, 20, true},
        {1, 100, 1, true},
        {1, 100, 100, true},
        {100, 1, 20, false},
    }

    for _, test := range testcases {
        actual := BuildIntChain().IsInRange(test.min, test.max).ValidateInt(test.num)
        t.Logf("IsInRange(%d, %d).ValidateInt(%d) = %v, expected = %v\n", test.min, test.max, test.num, actual, test.expected)
        if actual != test.expected {
            t.Errorf("[FAIL]")
        } else {
            t.Log("[PASS]")
        }
    }
}

/*
 */
func TestIsInListInt(t *testing.T) {
    var haystack = []int{20, 10, 30, 40, 50, 60, 0, 90, 80, 70, -10, -20, -50}
    testcases := []struct{
        haystack  []int
        num       int
        expected  bool
    }{
        {haystack, 20, true},
        {haystack, -50, true},
        {haystack, 90, true},
        {haystack, 0, true},
        {haystack, 43, false},
    }

    for _, test := range testcases {
        actual := BuildIntChain().IsInList(test.haystack).ValidateInt(test.num)
        t.Logf("IsInList(%v).ValidateInt(%d) = %v, expected = %v\n", test.haystack, test.num, actual, test.expected)
        if actual != test.expected {
            t.Errorf("[FAIL]")
        } else {
            t.Log("[PASS]")
        }
    }
}

/*
 */
func TestIsGreater(t *testing.T) {
    testcases := []struct{
        comp      int
        num       int
        expected  bool
    }{
        {1, 100, true},
        {-100, 10, true},
        {-100, -10, true},
        {200, 2, false},
        {-200, -100, true},
        {100, 100, false},
    }

    for _, test := range testcases {
        actual := BuildIntChain().IsGreater(test.comp).ValidateInt(test.num)
        t.Logf("IsGreater(%d).ValidateInt(%d) = %v, expected = %v\n", test.comp, test.num, actual, test.expected)
        if actual != test.expected {
            t.Errorf("[FAIL]")
        } else {
            t.Log("[PASS]")
        }
    }
}

/*
 */
func TestIsLess(t *testing.T) {
    testcases := []struct{
        comp      int
        num       int
        expected  bool
    }{
        {200, 2, true},
        {-200, -100, false},
        {1, 100, false},
        {-100, 10, false},
        {-100, -10, false},
        {100, 100, false},
    }

    for _, test := range testcases {
        actual := BuildIntChain().IsLess(test.comp).ValidateInt(test.num)
        t.Logf("IsLess(%d).ValidateInt(%d) = %v, expected = %v\n", test.comp, test.num, actual, test.expected)
        if actual != test.expected {
            t.Errorf("[FAIL]")
        } else {
            t.Log("[PASS]")
        }
    }
}


/* Tests chaining multiple validations.
 */
/*
func TestInListAndMaxLen(t *testing.T) {
    var haystack = []string{"get", "post", "put"}
    testcases := []struct{
        haystack  []string
        limit     int
        text      string
        expected  bool
    }{
        {haystack, 10, "get", true},
        {haystack, 4, "post", true},
        {haystack, 10, "put", true},
        {haystack, 2, "get", false},
        {haystack, 10, "foo", false},
        {haystack, 10, "GET", false},
        {haystack, 10, "pOSt", false},
    }

    for _, test := range testcases {
        actual := BuildStringChain().IsInList(test.haystack).IsMaxLen(test.limit).ValidateStr(test.text)
        t.Logf("IsIn(%v).IsMaxLen(%d).ValidateStr(%s) = %v, expected = %v\n", test.haystack, test.limit, test.text, actual, test.expected)
        if actual != test.expected {
            t.Errorf("[FAIL]")
        } else {
            t.Log("[PASS]")
        }
    }
}
*/ 
