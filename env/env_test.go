package env_test

import (
	"os"
	"testing"
	"time"

	"github.com/lestrrat/go-config/env"
	envload "github.com/lestrrat/go-envload"
	"github.com/stretchr/testify/assert"
)

type Spec struct {
	Embedded
	SimpleString          string
	SimpleInt             int
	SimpleInt8            int8
	SimpleInt16           int16
	SimpleInt32           int32
	SimpleInt64           int64
	SimpleUInt            uint
	SimpleUInt8           uint8
	SimpleUInt16          uint16
	SimpleUInt32          uint32
	SimpleUInt64          uint64
	SimpleFloat32         float32
	SimpleFloat64         float64
	ExplicitNameLowerCase string `env:"explicit_lower_case"`
	ExplicitNameUpperCase string `env:"EXPLICIT_UPPER_CASE"`
	Boolean               bool
	NestedStruct          Nested
	Pointer               *string
	PointerUninitialized  *string
	Time                  time.Time
	Duration              time.Duration
}

type Embedded struct {
	Message string
}

type Nested struct {
	Foo string
	Bar int
}

func TestDecode(t *testing.T) {
	l := envload.New()
	defer l.Restore()

	var s Spec

	now := time.Now().Truncate(time.Second).UTC()

	os.Setenv("MYAPP_EMBEDDED_MESSAGE", "Hello, Embedded!")
	os.Setenv("MYAPP_SIMPLESTRING", "foo")
	os.Setenv("MYAPP_SIMPLEINT", "100")
	os.Setenv("MYAPP_SIMPLEINT8", "100")
	os.Setenv("MYAPP_SIMPLEINT16", "100")
	os.Setenv("MYAPP_SIMPLEINT32", "100")
	os.Setenv("MYAPP_SIMPLEINT64", "100")
	os.Setenv("MYAPP_SIMPLEUINT", "100")
	os.Setenv("MYAPP_SIMPLEUINT8", "100")
	os.Setenv("MYAPP_SIMPLEUINT16", "100")
	os.Setenv("MYAPP_SIMPLEUINT32", "100")
	os.Setenv("MYAPP_SIMPLEUINT64", "100")
	os.Setenv("MYAPP_SIMPLEFLOAT32", "99.9")
	os.Setenv("MYAPP_SIMPLEFLOAT64", "99.9")
	os.Setenv("MYAPP_EXPLICIT_LOWER_CASE", "struct tag explicitly specifies lower case")
	os.Setenv("MYAPP_EXPLICIT_UPPER_CASE", "struct tag explicitly specifies upper case")
	os.Setenv("MYAPP_BOOLEAN", "true")
	os.Setenv("MYAPP_NESTEDSTRUCT_FOO", "foo")
	os.Setenv("MYAPP_NESTEDSTRUCT_BAR", "99")
	os.Setenv("MYAPP_POINTER", "pointer")
	os.Setenv("MYAPP_TIME", now.Format(time.RFC3339))

	if err := env.NewDecoder(env.System).Prefix("MYAPP").Decode(&s); !assert.NoError(t, err, "Decode should succeed") {
		t.Logf("%s", err)
		return
	}
	t.Logf("%#v", s)

	ptr := "pointer"
	var expected = Spec{
		Embedded:              Embedded{Message: "Hello, Embedded!"},
		SimpleString:          "foo",
		SimpleInt:             100,
		SimpleInt8:            int8(100),
		SimpleInt16:           int16(100),
		SimpleInt32:           int32(100),
		SimpleInt64:           int64(100),
		SimpleUInt:            uint(100),
		SimpleUInt8:           uint8(100),
		SimpleUInt16:          uint16(100),
		SimpleUInt32:          uint32(100),
		SimpleUInt64:          uint64(100),
		SimpleFloat32:         float32(99.9),
		SimpleFloat64:         float64(99.9),
		ExplicitNameLowerCase: "struct tag explicitly specifies lower case",
		ExplicitNameUpperCase: "struct tag explicitly specifies upper case",
		Boolean:               true,
		NestedStruct:          Nested{Foo: "foo", Bar: 99},
		Pointer:               &ptr,
		Time:                  now,
	}

	if !assert.Equal(t, expected, s, "result should match") {
		return
	}
}