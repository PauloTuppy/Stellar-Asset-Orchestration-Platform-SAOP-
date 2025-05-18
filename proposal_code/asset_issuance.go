type AssetIssuanceRequest struct {
    Code          string `json:"code"`
    Issuer        string `json:"issuer"`
    Compliance    ComplianceProfile `json:"compliance"`
    Distribution  DistributionRules `json:"distribution"`
}