type ComplianceRuleSet struct {
    versions     []RuleVersion
    active       *RuleVersion
    draft        *RuleVersion
    auditTrail   []RuleChange
}

type RuleVersion struct {
    id        string
    rules     map[string]ComplianceRule
    effective time.Time
    hash      string
}

func (c *ComplianceRuleSet) UpdateRules(update RuleUpdate) error {
    newVersion := c.draft.ApplyUpdate(update)
    
    // Legal approval workflow
    if err := legalReview(newVersion); err != nil {
        return err
    }
    
    // Technical validation
    if err := testRules(newVersion); err != nil {
        return err
    }
    
    c.versions = append(c.versions, newVersion)
    c.active = &newVersion
    c.logChange(update.Author)
    return nil
}