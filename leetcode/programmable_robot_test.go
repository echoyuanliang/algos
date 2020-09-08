package leetcode

import "testing"

func TestRobot(t *testing.T) {
        cases := []struct{
                Command string
                Obstacles [][]int
                X int
                Y int
                Result bool
        }{
                {
                        Command:"URR",
                        Obstacles: [][]int{},
                        X: 3,
                        Y: 2,
                        Result: true,
                },
                {
                        Command:"URR",
                        Obstacles: [][]int{{2,2}},
                        X: 3,
                        Y: 2,
                        Result: false,
                },
                {
                        Command:"URR",
                        Obstacles: [][]int{{4,2}},
                        X: 3,
                        Y: 2,
                        Result: true,
                },
                {
                        Command:"RUUR",
                        Obstacles: [][]int{{10, 5}, {1, 6}, {6, 10}, {3, 0}, {0, 3}, {0, 10}, {6, 2}},
                        X: 7856,
                        Y: 9033,
                        Result: false,
                },
        }

        for idx, c := range cases{

                if c.Result != Robot(c.Command, c.Obstacles, c.X, c.Y){
                        t.Fatalf("%d case fail", idx)
                }
        }


}