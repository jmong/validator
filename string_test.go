package validator

import(
    "testing"
)

/*
 */
func TestIsMaxLen(t *testing.T) {
    testcases := []struct{
        limit     int
        text      string
        expected  bool
    }{
        {5, "random word", false},
        {20, "random word", true},
        {18, "exactly 18 length", true},
    }

    for _, test := range testcases {
        actual := BuildStrChain().IsMaxLen(test.limit).ValidateStr(test.text)
        t.Logf("IsMaxLen(%d).ValidateStr(%s) = %v, expected = %v\n", test.limit, test.text, actual, test.expected)
        if actual != test.expected {
            t.Errorf("[FAIL]")
        } else {
            t.Log("[PASS]")
        }
    }
}

/*
 */
func TestIsInListStr(t *testing.T) {
    var haystack = []string{"get", "post", "put"}
    testcases := []struct{
        haystack  []string
        text      string
        expected  bool
    }{
        {haystack, "get", true},
        {haystack, "post", true},
        {haystack, "put", true},
        {haystack, "foo", false},
        {haystack, "GET", false},
        {haystack, "pOSt", false},
    }

    for _, test := range testcases {
        actual := BuildStrChain().IsInList(test.haystack).ValidateStr(test.text)
        t.Logf("IsIn(%v).ValidateStr(%s) = %v, expected = %v\n", test.haystack, test.text, actual, test.expected)
        if actual != test.expected {
            t.Errorf("[FAIL]")
        } else {
            t.Log("[PASS]")
        }
    }
}

/*
 */
func TestIsContains(t *testing.T) {
    testcases := []struct{
        substr    string
        text      string
        expected  bool
    }{
        {"I", "I am Spartacus!", true},
        {"I ", "I am Spartacus!", true},
        {"am", "I am Spartacus!", true},
        {" am ", "I am Spartacus!", true},
        {"Spartacus", "I am Spartacus!", true},
        {"Spartacus!", "I am Spartacus!", true},
        {"foo", "I am Spartacus!", false},
        {"spartacus", "I am Spartacus!", false},
        {"  ", "I am Spartacus!", false},
    }

    for _, test := range testcases {
        actual := BuildStrChain().IsContains(test.substr).ValidateStr(test.text)
        t.Logf("IsContains(%s).ValidateStr(%s) = %v, expected = %v\n", test.substr, test.text, actual, test.expected)
        if actual != test.expected {
            t.Errorf("[FAIL]")
        } else {
            t.Log("[PASS]")
        }
    }
}

/*
 */
func TestIsUpperFirstOnce(t *testing.T) {
    testcases := []struct{
        text      string
        expected  bool
    }{
        {"I am Spartacus!", true},
        {"i am Spartacus!", false},
        {" I am Spartacus!", true},
        {"I", true},
        {"i", false},
        {"  I", true},
        {"  i", false},
        {"@I", true},
        {"@i", false},
    }

    for _, test := range testcases {
        actual := BuildStrChain().IsUpperFirstOnce().ValidateStr(test.text)
        t.Logf("IsUpperFirstOnce().ValidateStr(%s) = %v, expected = %v\n", test.text, actual, test.expected)
        if actual != test.expected {
            t.Errorf("[FAIL]")
        } else {
            t.Log("[PASS]")
        }
    }
}

/*
 */
func TestIsLowerFirstOnce(t *testing.T) {
    testcases := []struct{
        text      string
        expected  bool
    }{
        {"I am Spartacus!", false},
        {"i am Spartacus!", true},
        {" I am Spartacus!", false},
        {"I", false},
        {"i", true},
        {"  I", false},
        {"  i", true},
        {"@I", false},
        {"@i", true},
    }

    for _, test := range testcases {
        actual := BuildStrChain().IsLowerFirstOnce().ValidateStr(test.text)
        t.Logf("IsLowerFirstOnce().ValidateStr(%s) = %v, expected = %v\n", test.text, actual, test.expected)
        if actual != test.expected {
            t.Errorf("[FAIL]")
        } else {
            t.Log("[PASS]")
        }
    }
}

/*
 */
func TestIsUpperFirstAll(t *testing.T) {
    testcases := []struct{
        text      string
        expected  bool
    }{
        {"I am Spartacus!", false},
        {"I Am Spartacus!", true},
        {"i am spartacus!", false},
        {"@I #Am ?Spartacus!", true},
        {"@i #am ?spartacus!", false},
        {"  I Am Spartacus!", true},
    }

    for _, test := range testcases {
        actual := BuildStrChain().IsUpperFirstAll().ValidateStr(test.text)
        t.Logf("IsUpperFirstAll().ValidateStr(%s) = %v, expected = %v\n", test.text, actual, test.expected)
        if actual != test.expected {
            t.Errorf("[FAIL]")
        } else {
            t.Log("[PASS]")
        }
    }
}

/*
 */
func TestIsLowerFirstAll(t *testing.T) {
    testcases := []struct{
        text      string
        expected  bool
    }{
        {"I am Spartacus!", false},
        {"I Am Spartacus!", false},
        {"i am spartacus!", true},
        {"@I #Am ?Spartacus!", false},
        {"@i #am ?spartacus!", true},
        {"  i am spartacus!", true},
    }

    for _, test := range testcases {
        actual := BuildStrChain().IsLowerFirstAll().ValidateStr(test.text)
        t.Logf("IsLowerFirstAll().ValidateStr(%s) = %v, expected = %v\n", test.text, actual, test.expected)
        if actual != test.expected {
            t.Errorf("[FAIL]")
        } else {
            t.Log("[PASS]")
        }
    }
}

/* Tests chaining multiple validators.
 */
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
        actual := BuildStrChain().IsInList(test.haystack).IsMaxLen(test.limit).ValidateStr(test.text)
        t.Logf("IsIn(%v).IsMaxLen(%d).ValidateStr(%s) = %v, expected = %v\n", test.haystack, test.limit, test.text, actual, test.expected)
        if actual != test.expected {
            t.Errorf("[FAIL]")
        } else {
            t.Log("[PASS]")
        }
    }
}
