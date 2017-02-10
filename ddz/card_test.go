package ddz

import (
	"testing"
)

func TestCardSliceFromString(t *testing.T) {
	var tests = []string{
		"♣3 ♣4 ♣5 ♣6 ♣7 ♣8 ♣9 ♣T ♣J ♣Q ♣K ♣A ♣2 ♦3 ♦4 ♦5 ♦6 ♦7 ♦8 ♦9 ♦T ♦J ♦Q ♦K ♦A ♦2 ♥3 ♥4 ♥5 ♥6 ♥7 ♥8 ♥9 ♥T ♥J ♥Q ♥K ♥A ♥2 ♠3 ♠4 ♠5 ♠6 ♠7 ♠8 ♠9 ♠T ♠J ♠Q ♠K ♠A ♠2 ♣r ♦R",
		"♣3,♣4,♣5,♣6,♣7,♣8,♣9",
		"♣K♣A哈♣2の♦3z♦4♦5♦6♦7♦8",
	}

	var values = [][]uint8{
		{0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2A, 0x2B, 0x2C, 0x2D, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3A, 0x3B, 0x3C, 0x3D, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4A, 0x4B, 0x4C, 0x4D, 0x1E, 0x2F},
		{0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17},
		{0x1B, 0x1C, 0x1D, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26},
	}

	for i := 0; i < len(tests); i++ {
		str := tests[i]
		val := values[i]
		cs := CardSliceFromString(str)
		if len(cs) != len(val) {
			t.Errorf("Test vector %d failed.", i)
			break
		}

		for j := 0; j < len(val); j++ {
			if cs[j] != val[j] {
				t.Errorf("Test vector %d failed at position %d.", i, j)
				break
			}
		}
	}
}

func TestCardSlice_Clone(t *testing.T) {
	a := CardSlice{0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17}
	b := a.Clone()
	if len(a) != len(b) {
		t.Error("Clone failed, length not equal.")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			t.Errorf("Clone failed, position %d not euqal.", i)
			break
		}
	}
}

func TestCardSlice_Concat(t *testing.T) {
	raw := CardSlice{0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17}
	a := CardSlice{0x11, 0x12, 0x13, 0x14}
	b := CardSlice{0x15, 0x16, 0x17}
	c := CardSlice{}
	d := a.Concat(b).Concat(c)
	if len(d) != len(a)+len(b)+len(c) {
		t.Error("Concat failed, length not sum up.")
	}

	for i := 0; i < len(d); i++ {
		if d[i] != raw[i] {
			t.Errorf("Concat failed, position %d not correct.", i)
		}
	}
}

//func test_whatever() {
//    f1 := ddz.CardSliceFromString("♣T ♦9 ♠8 ♥8 ♠7 ♣7 ♦6 ♣6 ♠5 ♣5 ♣4")
//
//    fmt.Println(f1.ToString())
//    fmt.Println("Standard analyze:")
//
//    standardList := f1.StandardAnalyze()
//    for i := 0; i < len(standardList); i++ {
//        fmt.Println(standardList[i].ToString())
//    }
//
//    fmt.Println("-----------------")
//    fmt.Println("Advance analyze:")
//
//    advanceList := f1.AdvanceAnalyze()
//    for i := 0; i < len(advanceList); i++ {
//        fmt.Println(advanceList[i].ToString())
//    }
//}
