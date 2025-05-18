impl CrossChainSwap {
    pub fn execute(&self) -> Result<(), SwapError> {
        let stellar_tx = build_stellar_lock_tx();
        let evm_tx = build_evm_unlock_tx();
        submit_atomic_batch(vec![stellar_tx, evm_tx]);
    }
}