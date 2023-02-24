package stacked

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestStackWithRealImpl(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	layer1 := &StackedImpl{}
	layer2 := NewMockStackedIntf(ctl, WithBackingStackedIntf(layer1))

	layer2.Inc() // not expected, but real implementation handles it
	if layer1.Get() != 1 {
		t.Errorf("Expected call to layer1.Inc() but it was never called")
	}
	if layer2.Get() != 1 {
		t.Errorf("Expected call to layer1.Get() but it was never called")
	}
}

func TestStackWithMockImpl(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	layer1 := NewMockStackedIntf(ctl)
	layer2 := NewMockStackedIntf(ctl, WithBackingStackedIntf(layer1))
	layer3 := NewMockStackedIntf(ctl, WithBackingStackedIntf(layer2))

	if layer3.frame.Level != 0 || layer2.frame.Level != 1 || layer1.frame.Level != 2 {
		t.Errorf("Unexpected layer level(s): %v, %v, %v",
			layer3.frame.Level, layer2.frame.Level, layer1.frame.Level)
	}

	layer1.EXPECT().Inc()
	layer2.EXPECT().Get()
	layer3.EXPECT().Inc()

	layer3.Inc() // from layer3
	layer3.Inc() // from layer1
	layer3.Get() // from layer2
}
