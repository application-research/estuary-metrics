import http from "./utils/HttpCommons";
import {VerifyResult} from "./models/VerifyResult";
import {EvaluationResult, Evaluation} from "./models/Evaluation";

const getTotalUsers = () => {
    return http.get<String>("/users/count");
};
const prove = (data: EvaluationResult) => {
    return http.post<String>(`/prove`, data);
};
const verify = (data: String) => {
    return http.post<VerifyResult>("/verify", data);
};
const estuary = {
    getTotalUsers,
    prove,
    verify,
};
export default estuary;

// await ZapsService.evaluate(data).then(result => {
//             let jsonResult: EvaluationResult = result.data;
//             setExpressionResult(JSON.stringify(jsonResult , null, 2));
//             //setExpressionResult("[iterations: " + jsonResult.Evaluation.iterations + " ] => " + jsonResult.Evaluation.expr_out);
//         }).catch(err => {
//             setExpressionResult(err.toString());
//             //setProveButtonState(true);
//         });