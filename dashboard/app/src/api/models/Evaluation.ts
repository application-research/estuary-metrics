export interface EvaluationResult {
    Evaluation: Evaluation;
}
export interface Evaluation {
    expr: string;
    env: string;
    cont: string;
    expr_out: string;
    env_out: string;
    cont_out: string;
    status: string;
    iterations: number;
}
