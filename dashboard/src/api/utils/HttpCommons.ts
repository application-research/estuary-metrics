import axios from "axios";

export default axios.create({
    baseURL: "https://localhost:3030",
    headers: {
        "Content-type": "application/json",
        Authorization: `Bearer ${process.env.ESTUARY_AUTH_KEY}`,
    }
});