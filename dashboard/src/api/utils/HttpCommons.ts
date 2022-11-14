import axios from "axios";

export default axios.create({
    baseURL: "https://metrics-api.onrender.com/",
    headers: {
        "Content-type": "application/json",s
        Authorization: `Bearer ${process.env.ESTUARY_AUTH_KEY}`,
    }
});