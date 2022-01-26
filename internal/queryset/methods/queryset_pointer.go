package methods

import "fmt"

// PointerEqFilterMethod filters with pointer EQ condition
type PointerEqFilterMethod struct {
	fieldName string
	chainedQuerySetMethod
	onFieldMethod
	oneArgMethod
}

func (m PointerEqFilterMethod) GetBody() string {
	tmpl := `if %s == nil {
		return qs.%sIsNull()
	} else {
		return qs.%sEq(*%s)
	}`
	return fmt.Sprintf(tmpl, m.getArgName(), m.fieldName,m.fieldName,m.getArgName())
}

func NewPointerEqFilterMethodImpl(ctx QsFieldContext) PointerEqFilterMethod {
	ctx = ctx.WithOperationName("EqP")
	return PointerEqFilterMethod{
		fieldName: ctx.fieldName(),
		onFieldMethod:         ctx.onFieldMethod(),
		oneArgMethod:          newOneArgMethod(fieldNameToArgName(ctx.fieldName()), ctx.fieldTypeName()),
		chainedQuerySetMethod: ctx.chainedQuerySetMethod(),
	}
}
