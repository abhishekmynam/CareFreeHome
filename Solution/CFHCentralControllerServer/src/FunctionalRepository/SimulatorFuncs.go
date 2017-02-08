package FunctionalRepository

import(
	CR "ConfigurationRepository"
)

func GetOutsideVals() CR.CurrCond{

	var conditions  CR.CurrCond

	conditions.ZipCode = "19333"
	conditions.Country = "USA"
	conditions.Working = false
	conditions.DateTime = "10Jan201612"
	conditions.Outvals.TempOut = 55
	conditions.Outvals.LightOut = 40
	conditions.Outvals.PplIn = 100
	conditions.Outvals.CondOut = "snow"

	return conditions
}
