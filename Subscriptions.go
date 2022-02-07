package msgraph

// ListSecureScores returns a slice of SecureScore objects. Each SecureScore represents
// Supports optional OData query parameters https://docs.microsoft.com/en-us/graph/query-parameters
// a tenant's security score for a particular day.
func (g *GraphClient) ListSubscribedSKUs(opts ...ListQueryOption) ([]SubscribedSKU, error) {
	resource := "/subscribedSkus"
	var marsh struct {
		SKUs []SubscribedSKU `json:"value"`
	}
	err := g.makeGETAPICall(resource, compileListQueryOptions(opts), &marsh)
	return marsh.SKUs, err
}

type SubscribedSKU struct {
	CapabilityStatus string `json:"capabilityStatus"`
	ConsumedUnits    int    `json:"consumedUnits"`
	ID               string `json:"id"`
	PrepaidUnits     struct {
		Enabled   int `json:"enabled"`
		Suspended int `json:"suspended"`
		Warning   int `json:"warning"`
	} `json:"prepaidUnits"`
	ServicePlans []struct {
		ServicePlanID      string `json:"servicePlanId"`
		ServicePlanName    string `json:"servicePlanName"`
		ProvisioningStatus string `json:"provisioningStatus"`
		AppliesTo          string `json:"appliesTo"`
	} `json:"servicePlans"`
	SkuID         string `json:"skuId"`
	SkuPartNumber string `json:"skuPartNumber"`
	AppliesTo     string `json:"appliesTo"`
}
