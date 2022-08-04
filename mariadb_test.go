package dbconnector

import "testing"

type tItem struct {
	TestString string
	Expected   string
}

var data = []tItem{
	{
		TestString: "",
		Expected:   "",
	}, {
		TestString: "m0.5s",
		Expected:   "",
	}, {
		TestString: "5m",
		Expected:   "5m",
	}, {
		TestString: "50ms",
		Expected:   "50ms",
	}, {
		TestString: "1h",
		Expected:   "1h",
	}, {
		TestString: ".6s",
		Expected:   "",
	}, {
		TestString: " 10s",
		Expected:   "",
	},
}

func Test_parseTimeout(t *testing.T) {
	for _, v := range data {
		observed := parseTimeout(v.TestString)
		if observed != v.Expected {
			t.Fatalf("parseTimeout(%v) = %v, want %v",
				v.TestString, observed, v.Expected)
		}
	}

	// fmt.Println(">>>>>>", mariadbs, *mariadbs_struct["test"])
}
