package moq_test

import (
	"reflect"
	"testing"

	"moqueries.org/runtime/moq"
)

type mockT struct {
	fatalfFormat string
	fatalfArgs   []interface{}
	helperCalled int
}

func (t *mockT) Errorf(string, ...interface{}) {}

func (t *mockT) Fatalf(format string, args ...interface{}) {
	t.fatalfFormat = format
	t.fatalfArgs = args
}

func (t *mockT) Helper() {
	t.helperCalled++
}

func TestRepeat(t *testing.T) {
	tests := map[string]struct {
		prev         *moq.RepeatVal
		increment    bool
		repeaters    []moq.Repeater
		want         *moq.RepeatVal
		fatalfFormat string
		fatalfArgs   []interface{}
	}{
		"default": {
			prev:      &moq.RepeatVal{},
			increment: true,
			repeaters: nil,
			want: &moq.RepeatVal{
				MinTimes:    1,
				MaxTimes:    1,
				AnyTimes:    false,
				ResultCount: 1,
				ExplicitAny: false,
				Incremented: true,
			},
			fatalfFormat: "",
			fatalfArgs:   nil,
		},
		"default prev default": {
			prev: &moq.RepeatVal{
				MinTimes:    1,
				MaxTimes:    1,
				AnyTimes:    false,
				ResultCount: 1,
				ExplicitAny: false,
				Incremented: true,
			},
			increment: true,
			repeaters: nil,
			want: &moq.RepeatVal{
				MinTimes:    2,
				MaxTimes:    2,
				AnyTimes:    false,
				ResultCount: 2,
				ExplicitAny: false,
				Incremented: true,
			},
			fatalfFormat: "",
			fatalfArgs:   nil,
		},
		"min": {
			prev:      &moq.RepeatVal{},
			increment: true,
			repeaters: []moq.Repeater{moq.MinTimes(2)},
			want: &moq.RepeatVal{
				MinTimes:    2,
				MaxTimes:    0,
				AnyTimes:    true,
				ResultCount: 3,
				ExplicitAny: false,
				Incremented: false,
			},
			fatalfFormat: "",
			fatalfArgs:   nil,
		},
		"min prev min": {
			prev: &moq.RepeatVal{
				MinTimes:    2,
				MaxTimes:    0,
				AnyTimes:    true,
				ResultCount: 3,
				ExplicitAny: false,
			},
			increment: true,
			repeaters: []moq.Repeater{moq.MinTimes(2)},
			want: &moq.RepeatVal{
				MinTimes:    4,
				MaxTimes:    0,
				AnyTimes:    true,
				ResultCount: 5,
				ExplicitAny: false,
			},
			fatalfFormat: "",
			fatalfArgs:   nil,
		},
		"min prev explicit any": {
			prev: &moq.RepeatVal{
				MinTimes:    2,
				MaxTimes:    0,
				AnyTimes:    true,
				ResultCount: 3,
				ExplicitAny: true,
			},
			increment: true,
			repeaters: []moq.Repeater{moq.MinTimes(2)},
			want: &moq.RepeatVal{
				MinTimes:    4,
				MaxTimes:    0,
				AnyTimes:    true,
				ResultCount: 5,
				ExplicitAny: true,
			},
			fatalfFormat: "",
			fatalfArgs:   nil,
		},
		"min and any": {
			prev:      &moq.RepeatVal{},
			increment: true,
			repeaters: []moq.Repeater{moq.MinTimes(2), moq.AnyTimes()},
			want: &moq.RepeatVal{
				MinTimes:    2,
				MaxTimes:    0,
				AnyTimes:    true,
				ResultCount: 3,
				ExplicitAny: true,
			},
			fatalfFormat: "",
			fatalfArgs:   nil,
		},
		"max": {
			prev:      &moq.RepeatVal{},
			increment: true,
			repeaters: []moq.Repeater{moq.MaxTimes(4)},
			want: &moq.RepeatVal{
				MinTimes:    0,
				MaxTimes:    4,
				AnyTimes:    false,
				ResultCount: 4,
				ExplicitAny: false,
			},
			fatalfFormat: "",
			fatalfArgs:   nil,
		},
		"max prev max": {
			prev: &moq.RepeatVal{
				MinTimes:    0,
				MaxTimes:    4,
				AnyTimes:    false,
				ResultCount: 4,
				ExplicitAny: false,
			},
			increment: true,
			repeaters: []moq.Repeater{moq.MaxTimes(4)},
			want: &moq.RepeatVal{
				MinTimes:    0,
				MaxTimes:    8,
				AnyTimes:    false,
				ResultCount: 8,
				ExplicitAny: false,
			},
			fatalfFormat: "",
			fatalfArgs:   nil,
		},
		"any": {
			prev:      &moq.RepeatVal{},
			increment: true,
			repeaters: []moq.Repeater{moq.AnyTimes()},
			want: &moq.RepeatVal{
				MinTimes:    0,
				MaxTimes:    0,
				AnyTimes:    true,
				ResultCount: 1,
				ExplicitAny: true,
			},
			fatalfFormat: "",
			fatalfArgs:   nil,
		},
		"optional": {
			prev:      &moq.RepeatVal{},
			increment: true,
			repeaters: []moq.Repeater{moq.Optional()},
			want: &moq.RepeatVal{
				MinTimes:    0,
				MaxTimes:    1,
				AnyTimes:    false,
				ResultCount: 1,
				ExplicitAny: false,
			},
			fatalfFormat: "",
			fatalfArgs:   nil,
		},
		"max and optional": {
			prev:      &moq.RepeatVal{},
			increment: true,
			repeaters: []moq.Repeater{moq.MaxTimes(4), moq.Optional()},
			want: &moq.RepeatVal{
				MinTimes:    0,
				MaxTimes:    4,
				AnyTimes:    false,
				ResultCount: 4,
				ExplicitAny: false,
			},
			fatalfFormat: "",
			fatalfArgs:   nil,
		},
		"times": {
			prev:      &moq.RepeatVal{},
			increment: true,
			repeaters: []moq.Repeater{moq.Times(5)},
			want: &moq.RepeatVal{
				MinTimes:    5,
				MaxTimes:    5,
				AnyTimes:    false,
				ResultCount: 5,
				ExplicitAny: false,
			},
			fatalfFormat: "",
			fatalfArgs:   nil,
		},
		"times and optional": {
			prev:      &moq.RepeatVal{},
			increment: true,
			repeaters: []moq.Repeater{moq.Times(5), moq.Optional()},
			want: &moq.RepeatVal{
				MinTimes:    0,
				MaxTimes:    5,
				AnyTimes:    false,
				ResultCount: 5,
				ExplicitAny: false,
			},
			fatalfFormat: "",
			fatalfArgs:   nil,
		},
		"times prev times": {
			prev: &moq.RepeatVal{
				MinTimes:    2,
				MaxTimes:    2,
				AnyTimes:    false,
				ResultCount: 2,
				ExplicitAny: false,
			},
			increment: true,
			repeaters: []moq.Repeater{moq.Times(3)},
			want: &moq.RepeatVal{
				MinTimes:    5,
				MaxTimes:    5,
				AnyTimes:    false,
				ResultCount: 5,
				ExplicitAny: false,
			},
			fatalfFormat: "",
			fatalfArgs:   nil,
		},
		"min prev max": {
			prev: &moq.RepeatVal{
				MinTimes:    0,
				MaxTimes:    4,
				AnyTimes:    false,
				ResultCount: 4,
				ExplicitAny: false,
			},
			increment: true,
			repeaters: []moq.Repeater{moq.MinTimes(2)},
			want: &moq.RepeatVal{
				MinTimes:    2,
				MaxTimes:    4,
				AnyTimes:    false,
				ResultCount: 4,
				ExplicitAny: false,
			},
			fatalfFormat: "",
			fatalfArgs:   nil,
		},
		"err/conflicting mins": {
			prev:         &moq.RepeatVal{},
			increment:    true,
			repeaters:    []moq.Repeater{moq.MinTimes(6), moq.MinTimes(8)},
			want:         nil,
			fatalfFormat: "repeat min of %d conflicts with min of %d",
			fatalfArgs:   []interface{}{6, moq.MinTimer(8)},
		},
		"err/conflicting maxes": {
			prev:         &moq.RepeatVal{},
			increment:    true,
			repeaters:    []moq.Repeater{moq.MaxTimes(10), moq.MaxTimes(12)},
			want:         nil,
			fatalfFormat: "repeat max of %d conflicts with max of %d",
			fatalfArgs:   []interface{}{10, moq.MaxTimer(12)},
		},
		"err/conflicting times": {
			prev:         &moq.RepeatVal{},
			increment:    true,
			repeaters:    []moq.Repeater{moq.Times(11), moq.Times(13)},
			want:         nil,
			fatalfFormat: "repeat times of %d conflicts with times of %d",
			fatalfArgs:   []interface{}{11, moq.Timer(13)},
		},
		"err/max then any": {
			prev:         &moq.RepeatVal{},
			increment:    true,
			repeaters:    []moq.Repeater{moq.MaxTimes(14), moq.AnyTimes()},
			want:         nil,
			fatalfFormat: "max of %d conflicts with moq.AnyTimes",
			fatalfArgs:   []interface{}{14},
		},
		"err/any then max": {
			prev:         &moq.RepeatVal{},
			increment:    true,
			repeaters:    []moq.Repeater{moq.AnyTimes(), moq.MaxTimes(16)},
			want:         nil,
			fatalfFormat: "max of %d conflicts with moq.AnyTimes",
			fatalfArgs:   []interface{}{16},
		},
		"err/times then min": {
			prev:         &moq.RepeatVal{},
			increment:    true,
			repeaters:    []moq.Repeater{moq.Times(14), moq.MinTimes(15)},
			want:         nil,
			fatalfFormat: "min of %d conflicts with times %d",
			fatalfArgs:   []interface{}{15, 14},
		},
		"err/min then times": {
			prev:         &moq.RepeatVal{},
			increment:    true,
			repeaters:    []moq.Repeater{moq.MinTimes(15), moq.Times(16)},
			want:         nil,
			fatalfFormat: "min of %d conflicts with times %d",
			fatalfArgs:   []interface{}{15, 16},
		},
		"err/times then max": {
			prev:         &moq.RepeatVal{},
			increment:    true,
			repeaters:    []moq.Repeater{moq.Times(14), moq.MaxTimes(15)},
			want:         nil,
			fatalfFormat: "max of %d conflicts with times %d",
			fatalfArgs:   []interface{}{15, 14},
		},
		"err/max then times": {
			prev:         &moq.RepeatVal{},
			increment:    true,
			repeaters:    []moq.Repeater{moq.MaxTimes(15), moq.Times(16)},
			want:         nil,
			fatalfFormat: "max of %d conflicts with times %d",
			fatalfArgs:   []interface{}{15, 16},
		},
		"err/min then small max": {
			prev:         &moq.RepeatVal{},
			increment:    true,
			repeaters:    []moq.Repeater{moq.MinTimes(18), moq.MaxTimes(9)},
			want:         nil,
			fatalfFormat: "max of %d is less than min of %d",
			fatalfArgs:   []interface{}{9, 18},
		},
		"err/min prev small max": {
			prev: &moq.RepeatVal{
				MinTimes:    0,
				MaxTimes:    9,
				AnyTimes:    false,
				ResultCount: 9,
				ExplicitAny: false,
			},
			increment:    true,
			repeaters:    []moq.Repeater{moq.MinTimes(18)},
			want:         nil,
			fatalfFormat: "max of %d is less than min of %d",
			fatalfArgs:   []interface{}{9, 18},
		},
		"err/max then large min": {
			prev:         &moq.RepeatVal{},
			increment:    true,
			repeaters:    []moq.Repeater{moq.MaxTimes(10), moq.MinTimes(20)},
			want:         nil,
			fatalfFormat: "max of %d is less than min of %d",
			fatalfArgs:   []interface{}{10, 20},
		},
		"err/min then optional": {
			prev:         &moq.RepeatVal{},
			increment:    true,
			repeaters:    []moq.Repeater{moq.MinTimes(15), moq.Optional()},
			want:         nil,
			fatalfFormat: "min of %d conflicts with moq.Optional",
			fatalfArgs:   []interface{}{15},
		},
		"err/optional then min": {
			prev:         &moq.RepeatVal{},
			increment:    true,
			repeaters:    []moq.Repeater{moq.Optional(), moq.MinTimes(15)},
			want:         nil,
			fatalfFormat: "min of %d conflicts with moq.Optional",
			fatalfArgs:   []interface{}{15},
		},
		"err/forgot to increment": {
			prev:         &moq.RepeatVal{},
			increment:    false,
			repeaters:    []moq.Repeater{moq.MinTimes(15)},
			want:         nil,
			fatalfFormat: "fn Increment not called before fn Repeat",
			fatalfArgs:   nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// ASSEMBLE
			tMock := &mockT{}

			expectedHelperCalls := 0
			if test.increment {
				expectedHelperCalls++
				test.prev.Increment(tMock)
			}

			// ACT
			if test.repeaters != nil {
				expectedHelperCalls++
				test.prev.Repeat(tMock, test.repeaters)
			}

			// ASSERT
			if expectedHelperCalls != tMock.helperCalled {
				t.Errorf("got %d helper calls, want %d", tMock.helperCalled, expectedHelperCalls)
			}
			if test.fatalfFormat != "" {
				if test.fatalfFormat != tMock.fatalfFormat {
					t.Errorf("got %s fatalf format, want %s", tMock.fatalfFormat, test.fatalfFormat)
				}
				if !reflect.DeepEqual(test.fatalfArgs, tMock.fatalfArgs) {
					t.Errorf("got %#v fatalf args, want %#v", tMock.fatalfArgs, test.fatalfArgs)
				}
			}
			// if test.want is nil, we are testing an error so don't check that
			// test.prev was updated
			if test.want != nil {
				if !reflect.DeepEqual(test.prev, test.want) {
					t.Errorf("got %#v, want %#v", test.prev, test.want)
				}
			}
		})
	}
}
