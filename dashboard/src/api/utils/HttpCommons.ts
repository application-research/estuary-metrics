import axios from "axios";

export default axios.create({
    baseURL: "https://metrics-api.onrender.com/",
    headers: {
        "Content-type": "application/json",
        Authorization: `Bearer ${process.env.ESTUARY_AUTH_KEY}`,
    }
});