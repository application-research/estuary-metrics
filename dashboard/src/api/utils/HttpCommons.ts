import axios from "axios";

export default axios.create({
    baseURL: "https://localhost:3030",
    headers: {
        "Content-type": "application/json",
        Authorization: `Bearer EST8010567a-7ae5-4679-8854-abfa8ade0bbfARY`
    }
});