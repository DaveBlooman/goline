package goline
/*
 *  Filename:    errors.go
 *  Package:     goline
 *  Author:      Bryan Matsuo <bmatsuo@soe.ucsc.edu>
 *  Created:     Tue Aug 23 19:42:49 PDT 2011
 *  Description: 
 */
import (
    "reflect"
    "os"
    "fmt"
)

type RecoverableError interface {
    os.Error
    IsRecoverable() bool
}

type ErrorNotInSet struct{ os.Error }

func (err ErrorNotInSet) IsRecoverable() bool { return true }

func (a *Answer) makeErrorNotInSet(r Responses, val interface{}) ErrorNotInSet {
    return ErrorNotInSet{
        fmt.Errorf("%s %s (%#v)", r[NotInSet], a.set.String(), val)}
}

/*
type ErrorOutOfRange struct {
    value, min, max interface{}
    err             os.Error
}

func (oor ErrorOutOfRange) String() string      { return oor.err.String() }
func (oor ErrorOutOfRange) IsRecoverable() bool { return true }

func errorOOR(val, min, max interface{}) ErrorOutOfRange {
    return ErrorOutOfRange{min, max, val,
        fmt.Errorf("Value %v out of range [%v, %v]", val, min, max),
    }
}
*/

type ErrorEmptyInput uint

func (oor ErrorEmptyInput) String() string      { return "Can not use empty value" }
func (oor ErrorEmptyInput) IsRecoverable() bool { return true }


func errorEmptyRange(min, max interface{}) os.Error {
    return fmt.Errorf("Range max is less than min (%v < %v)", min, max)
}

func errorTypeError(r Responses, expect, recv interface{}) os.Error {
    return fmt.Errorf("%s (%s != %s)",
        r[InvalidType], reflect.ValueOf(recv).Kind().String(), reflect.ValueOf(expect).Kind().String())
}
