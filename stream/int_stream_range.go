package stream

type RangeIntStream struct {
	begin int
	end   int
	step  int
}

func IntStreamFromZeroTo(end int) RangeIntStream {
	return RangeIntStream{0, end, 1}
}

func IntStreamFromRange(begin, end int) RangeIntStream {
	return RangeIntStream{begin, end, 1}
}

func IntStreamFromRangeAndStep(begin, end, step int) RangeIntStream {
	return RangeIntStream{begin, end, step}
}

func (s RangeIntStream) Begin() int {
	return s.begin
}

func (s RangeIntStream) End() int {
	return s.end
}

func (s RangeIntStream) Step() int {
	return s.step
}

func (s RangeIntStream) FilterInt(predicate func(int) bool) IntStream {
	return NewBasicIntStream(func(downstreamSignal <-chan struct{}) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			var cmp func(int, int) bool
			if s.step >= 0 {
				cmp = func(a, b int) bool {
					return a < b
				}
			} else {
				cmp = func(a, b int) bool {
					return a > b
				}
			}
		outer:
			for e := s.begin; cmp(e, s.end); e += s.step {
				if !predicate(e) {
					continue
				}
				select {
				case out <- e:
					// Sent, nothing to be done
				case _, ok := <-downstreamSignal:
					if !ok {
						break outer
					}
				}
			}
		}()
		return out
	})
}

func (s RangeIntStream) MapInt(mapper func(int) interface{}) Stream {
	return NewBasicStream(func(downstreamSignal <-chan struct{}) <-chan interface{} {
		out := make(chan interface{})
		go func() {
			defer close(out)
			var cmp func(int, int) bool
			if s.step >= 0 {
				cmp = func(a, b int) bool {
					return a < b
				}
			} else {
				cmp = func(a, b int) bool {
					return a > b
				}
			}
		outer:
			for e := s.begin; cmp(e, s.end); e += s.step {
				select {
				case out <- mapper(e):
					// Sent, nothing to be done
				case _, ok := <-downstreamSignal:
					if !ok {
						break outer
					}
				}
			}
		}()
		return out
	})
}

func (s RangeIntStream) ForEachInt(consumer func(int)) {
	var cmp func(int, int) bool
	if s.step >= 0 {
		cmp = func(a, b int) bool {
			return a < b
		}
	} else {
		cmp = func(a, b int) bool {
			return a > b
		}
	}
	for e := s.begin; cmp(e, s.end); e += s.step {
		consumer(e)
	}
}

// Adapter Methods

func (s RangeIntStream) Filter(predicate func(interface{}) bool) Stream {
	return s.FilterInt(func(x int) bool {
		return predicate(x)
	})
}

func (s RangeIntStream) Map(mapper func(interface{}) interface{}) Stream {
	return s.MapInt(func(x int) interface{} {
		return mapper(x)
	})
}

func (s RangeIntStream) ForEach(consumer func(interface{})) {
	s.ForEachInt(func(x int) {
		consumer(x)
	})
}
