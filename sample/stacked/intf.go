package stacked

//go:generate mockgen -stackable -destination mock.go -package stacked -source intf.go StackedIntf
//go:generate mockgen -stackable -destination mock_stacked/mock.go -source intf.go StackedIntf

type StackedIntf interface {
	Inc()
	Get() int
}

type StackedImpl struct {
	a int
}

func (s *StackedImpl) Inc() {
	s.a++
}

func (s *StackedImpl) Get() int {
	return s.a
}
