export const timeout = async (delay) => {
    return new Promise(res => setTimeout(res, delay));
}