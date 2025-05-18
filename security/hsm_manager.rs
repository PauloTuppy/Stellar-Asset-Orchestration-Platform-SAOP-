struct HSMKeyManager {
    hsm_client: HsmClient,
    key_version: HashMap<KeyID, Version>,
    rotation_policy: RotationPolicy,
}

impl HSMKeyManager {
    pub fn rotate_keys(&mut self) {
        for (key_id, policy) in &self.rotation_policy {
            if policy.should_rotate() {
                let new_version = self.hsm_client.generate_key(key_id);
                self.key_version.insert(key_id.clone(), new_version);
                self.archive_old_key(key_id);
            }
        }
    }
    
    pub fn sign_transaction(&self, tx: Transaction) -> SignedTransaction {
        let current_ver = self.key_version.get(tx.key_id).unwrap();
        let sig = self.hsm_client.sign_with_version(tx, current_ver);
        SignedTransaction { tx, sig }
    }
}