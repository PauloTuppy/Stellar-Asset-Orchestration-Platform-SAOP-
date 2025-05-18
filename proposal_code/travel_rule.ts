import Ajv from 'ajv';

type ValidationResult = boolean;

function loadSchema(schemaName: string): object {
    if (schemaName === 'IVMS101') {
        return {
            type: 'object',
            properties: {
                originator: {
                    type: 'object',
                    properties: {
                        name: { type: 'string' },
                        account: { type: 'string' },
                        jurisdiction: { type: 'string' },
                        id_type: { enum: ['LEI', 'DID'] }
                    },
                    required: ['name', 'account', 'jurisdiction', 'id_type']
                },
                beneficiary: {
                    type: 'object',
                    properties: {
                        name: { type: 'string' },
                        account: { type: 'string' },
                        jurisdiction: { type: 'string' }
                    },
                    required: ['name', 'account', 'jurisdiction']
                },
                transaction: {
                    type: 'object',
                    properties: {
                        amount: { type: 'string' },
                        currency: { type: 'string' },
                        tx_hash: { type: 'string' }
                    },
                    required: ['amount', 'currency', 'tx_hash']
                }
            },
            required: ['originator', 'beneficiary', 'transaction']
        };
    }
    throw new Error(`Unknown schema: ${schemaName}`);
}

interface TravelRuleData {
    originator: {
        name: string;
        account: string;
        jurisdiction: string;
        id_type: 'LEI' | 'DID';
    };
    beneficiary: {
        name: string;
        account: string;
        jurisdiction: string;
    };
    transaction: {
        amount: string;
        currency: string;
        tx_hash: string;
    };
}

function validateIVMS101(data: TravelRuleData): ValidationResult {
    const schema = loadSchema('IVMS101');
    const ajv = new Ajv({ strict: true });
    return ajv.validate(schema, data);
}
