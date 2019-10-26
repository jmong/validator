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


/* Tests chaining multiple validators.
 */
func TestInListAndIsInRange(t *testing.T) {
    var haystack = []int{0, 2, 4, 6, 8, 10, 12}
    testcases := []struct{
        num       int
        haystack  []int
        min, max  int
        expected  bool
    }{
        {0, haystack, 0, 10, true},
        {10, haystack, 0, 10, true},
        {4, haystack, 0, 10, true},
        {5, haystack, 0, 10, false},
        {12, haystack, 0, 10, false},
    }

    for _, test := range testcases {
        actual := BuildIntChain().IsInList(test.haystack).IsInRange(test.min, test.max).ValidateInt(test.num)
        t.Logf("IsInList(%v).IsInRange(%d, %d).ValidateInt(%d) = %v, expected = %v\n", test.haystack, test.min, test.max, test.num, actual, test.expected)
        if actual != test.expected {
            t.Errorf("[FAIL]")
        } else {
            t.Log("[PASS]")
        }
    }
}
