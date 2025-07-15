package main

import "testing"

type testCases struct {
    id        string
    right       []int64
    left        []int64
    expected    int
}

var tcs = []testCases{
    {
        id:       "test_case#1",
        right:    []int64{0,1,2,3,4,5,6,7},
        left:     []int64{7,6,5,4,3,2,1,0},
        expected: 0,
    },
    {
        id:       "test_case#2",
        right:    []int64{3,9,15,21,27,33,39,45},
        left:     []int64{12,18,6,24,30,36,42,48},
        expected: 24,
    },
    {
        id:       "test_case#3",
        right:    []int64{5,10,25,40,55,70,85,100,115},
        left:     []int64{15,20,35,50,65,80,95,110,125},
        expected: 90,
    },
    {
        id:       "test_case#4",
        right:    []int64{8,16,24,32,40,48,56,64,72,80},
        left:     []int64{12,28,36,44,52,60,68,76,84,92},
        expected: 112,
    },
    {
        id:       "test_case#5",
        right:    []int64{14,28,42,56,70,84,98,112},
        left:     []int64{21,35,49,63,77,91,105,119},
        expected: 56,
    },
    {
        id:       "test_case#6",
        right:    []int64{4,8,12,16,20,24,28,32,36},
        left:     []int64{6,10,14,18,22,26,30,34,38},
        expected: 18,
    },
    {
        id:       "test_case#7",
        right:    []int64{11,22,33,44,55,66,77,88,99,110},
        left:     []int64{15,30,45,60,75,90,105,120,135,150},
        expected: 220,
    },
    {
        id:       "test_case#8",
        right:    []int64{7,21,35,49,63,77,91,105},
        left:     []int64{14,28,42,56,70,84,98,112},
        expected: 56,
    },
    {
        id:       "test_case#9",
        right:    []int64{9,18,27,36,45,54,63,72,81},
        left:     []int64{12,24,36,48,60,72,84,96,108},
        expected: 135,
    },
    {
        id:       "test_case#10",
        right:    []int64{13,26,39,52,65,78,91,104,117,130},
        left:     []int64{17,34,51,68,85,102,119,136,153,170},
        expected: 220,
    },
    {
        id:       "test_case#11",
        right:    []int64{6,12,18,24,30,36,42,48,54,60,66},
        left:     []int64{8,16,24,32,40,48,56,64,72,80,88},
        expected: 132,
    },
    {
        id:       "test_case#12",
        right:    []int64{25,50,75,100,125,150,175,200},
        left:     []int64{30,60,90,120,150,180,210,240},
        expected: 180,
    },
    {
        id:       "test_case#13",
        right:    []int64{10,20,30,40,50,60,70,80,90,100,110,120},
        left:     []int64{15,25,35,45,55,65,75,85,95,105,115,125},
        expected: 60,
    },
    {
        id:       "test_case#14",
        right:    []int64{22,44,66,88,110,132,154,176,198},
        left:     []int64{33,55,77,99,121,143,165,187,209},
        expected: 99,
    },
    {
        id:       "test_case#15",
        right:    []int64{18,36,54,72,90,108,126,144,162,180},
        left:     []int64{24,48,72,96,120,144,168,192,216,240},
        expected: 330,
    },
    {
        id:       "test_case#16",
        right:    []int64{16,32,48,64,80,96,112,128,144,160,176},
        left:     []int64{20,40,60,80,100,120,140,160,180,200,220},
        expected: 264,
    },
    {
        id:       "test_case#17",
        right:    []int64{19,38,57,76,95,114,133,152},
        left:     []int64{23,46,69,92,115,138,161,184},
        expected: 144,
    },
    {
        id:       "test_case#18",
        right:    []int64{27,54,81,108,135,162,189,216,243,270},
        left:     []int64{31,62,93,124,155,186,217,248,279,310},
        expected: 220,
    },
    {
        id:       "test_case#19",
        right:    []int64{12,24,36,48,60,72,84,96,108,120,132,144,156},
        left:     []int64{18,36,54,72,90,108,126,144,162,180,198,216,234},
        expected: 546,
    },
    {
        id:       "test_case#20",
        right:    []int64{35,70,105,140,175,210,245,280,315,350,385,420,455,490},
        left:     []int64{42,84,126,168,210,252,294,336,378,420,462,504,546,588},
        expected: 735,
    },
}

func TestSolver(t *testing.T) {
    t.Parallel()
    for _, tc := range tcs {
        tc := tc
        t.Run(tc.id, func(t *testing.T) {
            t.Parallel()
            t.Logf("Running test case: %s", tc.id)

            l := NewLists()
            l.Left = append([]int64(nil), tc.left...)
            l.Right = append([]int64(nil), tc.right...)

            got := solve(l, len(tc.left))

            if got != tc.expected {
                t.Errorf("FAILED: %s. Expected: %d, Got: %d", tc.id, tc.expected, got)
            } else {
                t.Logf("PASSED: %s", tc.id)
            }
        })
    }
}

func BenchmarkSolver(b *testing.B) {
    for _, c := range tcs {
        b.Run(c.id, func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                l := NewLists()
                l.Left = append([]int64(nil), c.left...)
                l.Right = append([]int64(nil), c.right...)

                _ = solve(l, len(c.left))
            }
        })
    }
}