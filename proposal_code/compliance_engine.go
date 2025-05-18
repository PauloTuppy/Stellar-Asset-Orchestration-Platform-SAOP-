type ComplianceRule interface {
    Validate(tx Transaction) ComplianceResult
    Version() semver.Version
    Jurisdiction() string
}

type ComplianceEngine struct {
    rules      []ComplianceRule
    ruleLoader RuleLoader
    auditLog   AuditLogger
}

func (e *ComplianceEngine) Process(tx Transaction) error {
    results := make(chan ComplianceResult, len(e.rules))
    
    for _, rule := range e.rules {
        go func(r ComplianceRule) {
            results <- r.Validate(tx)
        }(rule)
    }
    
    // Process results
    for range e.rules {
        result := <-results
        if !result.Valid {
            e.auditLog.LogViolation(tx, result)
            return ErrComplianceViolation
        }
    }
    return nil
}