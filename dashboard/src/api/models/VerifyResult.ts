export interface VerifyResult {
    claim_cid: ClaimCidOrProofCid;
    proof_cid: ClaimCidOrProofCid;
    signature: string;
    verified: Verified;
    verifier_id: string;
}
export interface ClaimCidOrProofCid {
    "/": string;
}
export interface Verified {
    verified: boolean;
}
