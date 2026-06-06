package variables

type VariablesModel struct {
	RestVars    func()
	Convert     func(numero int) (bool, string)
	VarsNumeric func()
}

func NewClient() *VariablesModel {

	return &VariablesModel{
		RestVars:    RestoVariables,
		Convert:     ConvertToText,
		VarsNumeric: MuestroEnteros,
	}
}
