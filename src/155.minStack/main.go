package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/min-stack/
func main() {
	obj := Constructor()
	obj.Push(395)
	min := obj.GetMin()
	top := obj.Top()
	min = obj.GetMin()
	obj.Push(276)
	obj.Push(29)
	min = obj.GetMin()
	obj.Push(-482)
	min = obj.GetMin()
	obj.Pop()
	obj.Push(-108)
	obj.Push(-251)
	min = obj.GetMin()
	obj.Push(-439)
	top = obj.Top()
	obj.Push(370)
	obj.Pop()
	obj.Pop()
	min = obj.GetMin()
	obj.Pop()
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Pop()
	min = obj.GetMin()
	obj.Push(-158)
	obj.Push(82)
	obj.Pop()
	obj.Push(-176)
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Push(-90)
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Push(411)
	min = obj.GetMin()
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Push(-494)
	min = obj.GetMin()
	obj.Pop()
	top = obj.Top()
	obj.Push(155)
	top = obj.Top()
	obj.Push(-370)
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Push(273)
	obj.Pop()
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Pop()
	obj.Pop()
	min = obj.GetMin()
	obj.Push(173)
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Push(0)
	min = obj.GetMin()
	top = obj.Top()
	min = obj.GetMin()
	obj.Push(-266)
	min = obj.GetMin()
	obj.Push(-359)
	top = obj.Top()
	min = obj.GetMin()
	obj.Push(189)
	min = obj.GetMin()
	top = obj.Top()
	min = obj.GetMin()
	obj.Push(375)
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Push(-188)
	obj.Pop()
	obj.Push(-275)
	obj.Push(-223)
	min = obj.GetMin()
	obj.Push(294)
	obj.Push(380)
	obj.Push(-42)
	top = obj.Top()
	min = obj.GetMin()
	obj.Push(33)
	min = obj.GetMin()
	obj.Push(457)
	obj.Push(422)
	obj.Push(-352)
	min = obj.GetMin()
	obj.Push(326)
	obj.Push(-378)
	obj.Push(231)
	obj.Pop()
	obj.Push(416)
	min = obj.GetMin()
	top = obj.Top()
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Push(277)
	top = obj.Top()
	obj.Push(472)
	obj.Push(205)
	top = obj.Top()
	obj.Push(-443)
	min = obj.GetMin()
	obj.Push(-5)
	top = obj.Top()
	min = obj.GetMin()
	min = obj.GetMin()
	min = obj.GetMin()
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Push(-312)
	min = obj.GetMin()
	obj.Push(-249)
	obj.Push(-209)
	min = obj.GetMin()
	min = obj.GetMin()
	min = obj.GetMin()
	top = obj.Top()
	min = obj.GetMin()
	obj.Push(196)
	obj.Pop()
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Push(-347)
	min = obj.GetMin()
	min = obj.GetMin()
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Push(400)
	top = obj.Top()
	top = obj.Top()
	obj.Push(269)
	obj.Push(434)
	obj.Push(-213)
	top = obj.Top()
	top = obj.Top()
	obj.Push(-58)
	min = obj.GetMin()
	obj.Push(-353)
	obj.Push(-426)
	obj.Push(-335)
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Push(-188)
	obj.Push(-388)
	obj.Push(369)
	obj.Push(319)
	min = obj.GetMin()
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Push(121)
	top = obj.Top()
	obj.Pop()
	min = obj.GetMin()
	obj.Push(155)
	top = obj.Top()
	obj.Pop()
	obj.Push(155)
	min = obj.GetMin()
	obj.Pop()
	min = obj.GetMin()
	obj.Pop()
	min = obj.GetMin()
	obj.Push(495)
	top = obj.Top()
	obj.Push(-53)
	min = obj.GetMin()
	min = obj.GetMin()
	top = obj.Top()
	obj.Pop()
	top = obj.Top()
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Push(465)
	obj.Push(162)
	obj.Push(-334)
	obj.Pop()
	obj.Push(282)
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Push(432)
	obj.Push(-418)
	obj.Push(195)
	top = obj.Top()
	min = obj.GetMin()
	top = obj.Top()
	min = obj.GetMin()
	min = obj.GetMin()
	top = obj.Top()
	top = obj.Top()
	obj.Pop()
	obj.Pop()
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Push(374)
	min = obj.GetMin()
	obj.Push(202)
	min = obj.GetMin()
	obj.Push(181)
	obj.Push(173)
	obj.Push(-323)
	min = obj.GetMin()
	obj.Pop()
	obj.Pop()
	obj.Push(-430)
	obj.Pop()
	top = obj.Top()
	obj.Push(346)
	top = obj.Top()
	top = obj.Top()
	obj.Pop()
	top = obj.Top()
	obj.Push(244)
	obj.Push(-467)
	top = obj.Top()
	top = obj.Top()
	min = obj.GetMin()
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Push(279)
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Push(-30)
	min = obj.GetMin()
	obj.Pop()
	top = obj.Top()
	obj.Push(264)
	obj.Push(-217)
	obj.Push(-418)
	obj.Push(12)
	obj.Push(-439)
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Push(7)
	min = obj.GetMin()
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Push(-189)
	min = obj.GetMin()
	min = obj.GetMin()
	min = obj.GetMin()
	top = obj.Top()
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Push(303)
	top = obj.Top()
	min = obj.GetMin()
	obj.Push(-297)
	min = obj.GetMin()
	obj.Push(-21)
	min = obj.GetMin()
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Push(461)
	obj.Pop()
	obj.Push(-303)
	obj.Pop()
	obj.Push(-338)
	top = obj.Top()
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Push(-42)
	min = obj.GetMin()
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Push(85)
	obj.Push(132)
	obj.Push(57)
	min = obj.GetMin()
	obj.Push(-17)
	top = obj.Top()
	obj.Push(-10)
	min = obj.GetMin()
	obj.Push(-500)
	top = obj.Top()
	min = obj.GetMin()
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Pop()
	min = obj.GetMin()
	top = obj.Top()
	min = obj.GetMin()
	obj.Push(-7)
	min = obj.GetMin()
	min = obj.GetMin()
	min = obj.GetMin()
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Pop()
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Push(346)
	min = obj.GetMin()
	obj.Pop()
	obj.Push(16)
	top = obj.Top()
	obj.Push(472)
	obj.Push(-211)
	min = obj.GetMin()
	obj.Pop()
	obj.Pop()
	obj.Push(-381)
	obj.Pop()
	min = obj.GetMin()
	obj.Push(214)
	obj.Push(169)
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Pop()
	obj.Pop()
	obj.Pop()
	obj.Push(33)
	obj.Pop()
	obj.Push(115)
	obj.Push(-346)
	obj.Push(-309)
	obj.Push(130)
	min = obj.GetMin()
	top = obj.Top()
	min = obj.GetMin()
	min = obj.GetMin()
	min = obj.GetMin()
	top = obj.Top()
	obj.Push(335)
	min = obj.GetMin()
	obj.Push(-92)
	obj.Push(-16)
	min = obj.GetMin()
	obj.Pop()
	min = obj.GetMin()
	obj.Push(-470)
	obj.Pop()
	obj.Pop()
	obj.Push(250)
	obj.Push(327)
	obj.Push(144)
	obj.Pop()
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Pop()
	obj.Push(359)
	obj.Push(138)
	min = obj.GetMin()
	top = obj.Top()
	min = obj.GetMin()
	obj.Pop()
	obj.Push(272)
	obj.Push(-241)
	obj.Push(-362)
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Push(-83)
	obj.Push(195)
	obj.Push(-208)
	min = obj.GetMin()
	obj.Pop()
	min = obj.GetMin()
	obj.Push(-500)
	obj.Push(5)
	min = obj.GetMin()
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Push(284)
	min = obj.GetMin()
	min = obj.GetMin()
	min = obj.GetMin()
	min = obj.GetMin()
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Push(477)
	min = obj.GetMin()
	obj.Pop()
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Push(30)
	obj.Pop()
	obj.Pop()
	obj.Pop()
	obj.Push(-269)
	top = obj.Top()
	obj.Push(-413)
	obj.Push(-323)
	min = obj.GetMin()
	min = obj.GetMin()
	min = obj.GetMin()
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Push(217)
	obj.Push(-408)
	top = obj.Top()
	obj.Push(-353)
	top = obj.Top()
	obj.Push(-142)
	top = obj.Top()
	obj.Pop()
	obj.Push(-321)
	min = obj.GetMin()
	obj.Push(-33)
	obj.Push(382)
	min = obj.GetMin()
	min = obj.GetMin()
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Pop()
	min = obj.GetMin()
	obj.Push(298)
	top = obj.Top()
	obj.Pop()
	obj.Push(495)
	min = obj.GetMin()
	obj.Push(195)
	obj.Push(-147)
	obj.Push(-85)
	obj.Push(195)
	obj.Pop()
	min = obj.GetMin()
	obj.Push(154)
	obj.Push(-311)
	top = obj.Top()
	min = obj.GetMin()
	min = obj.GetMin()
	top = obj.Top()
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Push(-390)
	min = obj.GetMin()
	obj.Push(323)
	min = obj.GetMin()
	top = obj.Top()
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Push(338)
	obj.Push(263)
	min = obj.GetMin()
	obj.Push(160)
	obj.Push(148)
	obj.Push(142)
	obj.Push(427)
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Pop()
	obj.Push(-153)
	top = obj.Top()
	obj.Push(-384)
	obj.Pop()
	min = obj.GetMin()
	obj.Push(159)
	obj.Push(419)
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Push(-385)
	min = obj.GetMin()
	obj.Push(461)
	obj.Push(-346)
	min = obj.GetMin()
	min = obj.GetMin()
	min = obj.GetMin()
	top = obj.Top()
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Push(343)
	top = obj.Top()
	obj.Push(-269)
	top = obj.Top()
	obj.Pop()
	obj.Push(-401)
	obj.Push(81)
	min = obj.GetMin()
	obj.Push(130)
	obj.Pop()
	obj.Pop()
	obj.Push(-428)
	min = obj.GetMin()
	obj.Push(454)
	min = obj.GetMin()
	min = obj.GetMin()
	min = obj.GetMin()
	top = obj.Top()
	top = obj.Top()
	obj.Pop()
	obj.Pop()
	obj.Pop()
	min = obj.GetMin()
	obj.Push(59)
	top = obj.Top()
	obj.Push(143)
	min = obj.GetMin()
	min = obj.GetMin()
	min = obj.GetMin()
	obj.Push(-154)
	min = obj.GetMin()
	top = obj.Top()
	min = obj.GetMin()
	obj.Push(114)
	obj.Push(-346)
	min = obj.GetMin()
	obj.Pop()

	fmt.Println("min:", min, "top:", top)
}

type MinStack struct {
	stack    []int
	hasMin   bool
	min      int
	minCount int
	sorted   []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(x int) {
	s.stack = append(s.stack, x)
	if !s.hasMin {
		s.min = x
		s.minCount = 1
		s.hasMin = true
	} else if s.min > x {
		s.min = x
		s.minCount = 1
	} else if s.min == x {
		s.minCount++
	}
}

func (s *MinStack) Pop() {
	e := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	if e == s.min {
		s.minCount--
		if s.minCount == 0 {
			if len(s.stack) == 0 {
				s.min = 0
				s.hasMin = false
			} else {
				// resort
				s.sorted = make([]int, len(s.stack))
				copy(s.sorted, s.stack)
				sort.Ints(s.sorted)
				s.min = s.sorted[0]
				s.minCount = 1
				for i := 1; i < len(s.sorted); i++ {
					if s.min == s.sorted[i] {
						s.minCount++
					} else {
						break
					}
				}
			}
		}
	}
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	return s.min
}
