export interface Proof {
    claim: Claim;
    proof: Proof1;
    reduction_count: string;
}
export interface Claim {
    Opening: Opening;
}
export interface Opening {
    input: string;
    output: string;
    status: string;
    commitment: string;
    new_commitment?: null;
}
export interface Proof1 {
    proof: Proof2;
    proof_count: number;
    chunk_frame_count: number;
}
export interface Proof2 {
    num_inputs: number;
    pi_agg: PiAgg;
    com_f?: ((number)[] | null)[] | null;
    com_w0?: ((number)[] | null)[] | null;
    com_wd?: ((number)[] | null)[] | null;
    f_eval?: ((number)[] | null)[] | null;
    f_eval_proof?: ((number)[] | null)[] | null;
}
export interface PiAgg {
    com_ab?: (ComAbEntityOrComCEntityOrEntityOrEntityEntityOrCommsAbEntityEntityEntityOrCommsCEntityEntityEntityOrZAbEntityEntityOrIpAb)[] | null;
    com_c?: (ComAbEntityOrComCEntityOrEntityOrEntityEntityOrCommsAbEntityEntityEntityOrCommsCEntityEntityEntityOrZAbEntityEntityOrIpAb)[] | null;
    ip_ab: ComAbEntityOrComCEntityOrEntityOrEntityEntityOrCommsAbEntityEntityEntityOrCommsCEntityEntityEntityOrZAbEntityEntityOrIpAb;
    agg_c?: (number)[] | null;
    tmipp: Tmipp;
}
export interface ComAbEntityOrComCEntityOrEntityOrEntityEntityOrCommsAbEntityEntityEntityOrCommsCEntityEntityEntityOrZAbEntityEntityOrIpAb {
    c0: C0OrC1;
    c1: C0OrC1;
}
export interface C0OrC1 {
    c0: C0OrC1OrC2;
    c1: C0OrC1OrC2;
    c2: C0OrC1OrC2;
}
export interface C0OrC1OrC2 {
    c0?: (number)[] | null;
    c1?: (number)[] | null;
}
export interface Tmipp {
    gipa: Gipa;
    vkey_opening?: ((number)[] | null)[] | null;
    wkey_opening?: ((number)[] | null)[] | null;
}
export interface Gipa {
    nproofs: number;
    comms_ab?: (((ComAbEntityOrComCEntityOrEntityOrEntityEntityOrCommsAbEntityEntityEntityOrCommsCEntityEntityEntityOrZAbEntityEntityOrIpAb)[] | null)[] | null)[] | null;
    comms_c?: (((ComAbEntityOrComCEntityOrEntityOrEntityEntityOrCommsAbEntityEntityEntityOrCommsCEntityEntityEntityOrZAbEntityEntityOrIpAb)[] | null)[] | null)[] | null;
    z_ab?: ((ComAbEntityOrComCEntityOrEntityOrEntityEntityOrCommsAbEntityEntityEntityOrCommsCEntityEntityEntityOrZAbEntityEntityOrIpAb)[] | null)[] | null;
    z_c?: (((number)[] | null)[] | null)[] | null;
    final_a?: (number)[] | null;
    final_b?: (number)[] | null;
    final_c?: (number)[] | null;
    final_vkey?: ((number)[] | null)[] | null;
    final_wkey?: ((number)[] | null)[] | null;
}
