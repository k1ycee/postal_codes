package main

/* Each state in Nigeria has districts with postal codes
   how this works is that A state is divided into Local government areas
   Each local government has a district and each district has postal code
*/

// type PostalCodeData struct {
// 	districtName        string
// 	localGovernmentArea string
// 	state               string
// 	postalCode          string
// }

type State struct {
	name     string
	district District
}

type District struct {
	name                string
	localGovernmentArea string
	postalCode          string
}
