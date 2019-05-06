package any

import "strconv"

type Value string

func (f Value) String() string { return string(f) }

func (f Value) ToInt() (int, error) { return strconv.Atoi(f.String()) }

func (f Value) ToBool() (bool, error) { return strconv.ParseBool(f.String()) }

func (f Value) ToFloat64() (float64, error) { return strconv.ParseFloat(f.String(), 64) }
