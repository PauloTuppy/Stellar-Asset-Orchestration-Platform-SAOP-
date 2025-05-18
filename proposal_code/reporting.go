type ReportGenerator struct {
    templates   map[Regime]ReportTemplate
    dataSources []DataSource
    crypto      CryptoEngine
}

func (g *ReportGenerator) Generate(reportType Regime, period time.Time) (Report, error) {
    template := g.templates[reportType]
    data := g.collectData(period, template.Requirements)
    
    report := applyTemplate(template, data)
    signed := g.crypto.SignReport(report)
    
    store.Archive(signed)
    return submitToAuthority(reportType, signed)
}