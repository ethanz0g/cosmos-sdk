package depinject

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"

	"github.com/cosmos/cosmos-sdk/depinject/internal/util"
)

// ProviderDescriptor defines a special provider type that is defined by
// reflection. It should be passed as a value to the Provide function.
// Ex:
//   option.Provide(ProviderDescriptor{ ... })
type ProviderDescriptor struct {
	// Inputs defines the in parameter types to Fn.
	Inputs []ProviderInput

	// Outputs defines the out parameter types to Fn.
	Outputs []ProviderOutput

	// Fn defines the provider function.
	Fn func([]reflect.Value) ([]reflect.Value, error)

	// Location defines the source code location to be used for this provider
	// in error messages.
	Location Location

	hasError bool
}

type ProviderInput struct {
	Type     reflect.Type
	Optional bool

	// startStructType is set to the type of an In struct on the first field of that struct only
	startStructType reflect.Type

	// structFieldName is set if this sets an In struct field
	structFieldName string
}

type ProviderOutput struct {
	Type reflect.Type

	// startStructType is set to the type of an Out struct on the first field of that struct only
	startStructType reflect.Type

	// structFieldName is set if this sets an Out struct field
	structFieldName string
}

func ExtractProviderDescriptor(provider interface{}) (ProviderDescriptor, error) {
	rctr, ok := provider.(ProviderDescriptor)
	if !ok {
		var err error
		rctr, err = doExtractProviderDescriptor(provider)
		if err != nil {
			return ProviderDescriptor{}, err
		}
	}

	return expandStructArgsProvider(rctr)
}

func doExtractProviderDescriptor(ctr interface{}) (ProviderDescriptor, error) {
	val := reflect.ValueOf(ctr)
	typ := val.Type()
	if typ.Kind() != reflect.Func {
		return ProviderDescriptor{}, errors.Errorf("expected a Func type, got %v", typ)
	}

	loc := LocationFromPC(val.Pointer())

	if typ.IsVariadic() {
		return ProviderDescriptor{}, errors.Errorf("variadic function can't be used as a provider: %s", loc)
	}

	numIn := typ.NumIn()
	in := make([]ProviderInput, numIn)
	for i := 0; i < numIn; i++ {
		in[i] = ProviderInput{
			Type: typ.In(i),
		}
	}

	errIdx := -1
	numOut := typ.NumOut()
	var out []ProviderOutput
	for i := 0; i < numOut; i++ {
		t := typ.Out(i)
		if t == errType {
			if i != numOut-1 {
				return ProviderDescriptor{}, errors.Errorf("output error parameter is not last parameter in function %s", loc)
			}
			errIdx = i
		} else {
			out = append(out, ProviderOutput{Type: t})
		}
	}

	hasError := errIdx >= 0
	return ProviderDescriptor{
		Inputs:  in,
		Outputs: out,
		Fn: func(values []reflect.Value) ([]reflect.Value, error) {
			res := val.Call(values)
			if hasError {
				err := res[errIdx]
				if !err.IsZero() {
					return nil, err.Interface().(error)
				}
				return res[0:errIdx], nil
			}
			return res, nil
		},
		Location: loc,
		hasError: hasError,
	}, nil
}

var errType = reflect.TypeOf((*error)(nil)).Elem()

func (p ProviderDescriptor) codegenOutputs(ctr *container, suffix string) (varsDef string, valueExprs []expr) {
	var varRefs []varRef
	var curStructVar expr
	for _, output := range p.Outputs {
		var name string
		if output.structFieldName != "" {
			if output.startStructType == nil {
				valueExprs = append(valueExprs, fieldRef{
					e:         curStructVar,
					fieldName: output.structFieldName,
				})
				continue
			}

			name = output.startStructType.Name()
		} else {
			curStructVar = nil
			name = output.Type.Name()
		}

		v := ctr.createVar(fmt.Sprintf("%s%s", util.StringFirstLower(name), suffix))
		varRefs = append(varRefs, v)
		if output.structFieldName != "" {
			curStructVar = v
			valueExprs = append(valueExprs, fieldRef{
				e:         curStructVar,
				fieldName: output.structFieldName,
			})
		} else {
			valueExprs = append(valueExprs, v)
		}
	}

	first := true
	for _, valueVar := range varRefs {
		if !first {
			varsDef += ", "
		}
		varsDef += valueVar.emit()
		first = false
	}
	if p.hasError {
		varsDef += ", err"
	}
	varsDef += " := "
	return varsDef, valueExprs
}

func (p ProviderDescriptor) codegenErrCheck(ctr *container) {
	if p.hasError {
		ctr.codegenWriteln("if err != nil {")
		ctr.codegenWriteln("    return err")
		ctr.codegenWriteln("}")
	}
}
