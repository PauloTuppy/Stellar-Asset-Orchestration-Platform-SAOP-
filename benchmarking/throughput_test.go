type ThroughputTest struct {
    NetworkConfig StellarConfig
    TestMode      TestMode // TPS/Stress/Latency
}

func (t *ThroughputTest) Run() BenchmarkResult {
    generator := NewTxGenerator(t.NetworkConfig)
    results := make(chan TxResult, 100000)
    
    // Parallel submission
    for i := 0; i < t.Config.Workers; i++ {
        go func() {
            for tx := range generator.Txs {
                start := time.Now()
                submitTx(tx)
                results <- TxResult{Latency: time.Since(start)}
            }
        }()
    }
    
    return aggregateResults(results)
}