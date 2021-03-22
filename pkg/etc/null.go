package etc

import (
	"database/sql"

	"github.com/golang/protobuf/ptypes/wrappers"
)

//NullString returns value if it is valid
func NullString(s *wrappers.StringValue) (ns sql.NullString) {
	if s != nil {
		ns.String = s.Value
		ns.Valid = true
	}
	return ns
}

//NullFloat64 returns value if it is valid
func NullDouble(s *wrappers.DoubleValue) (ns sql.NullFloat64) {
	if s != nil {
		ns.Float64 = s.Value
		ns.Valid = true
	}
	return ns
}

//StringValue ...
func StringValue(ns sql.NullString) *wrappers.StringValue {
	if ns.Valid {
		s := wrappers.StringValue{Value: ns.String}
		return &s
	}
	return nil
}

//DoubleValue ...
func DoubleValue(ns sql.NullFloat64) *wrappers.DoubleValue {
	if ns.Valid {
		s := wrappers.DoubleValue{Value: ns.Float64}
		return &s
	}
	return nil
}
