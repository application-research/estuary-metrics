export const timeout = async (delay: number | undefined) => {
    return new Promise(res => setTimeout(res, delay));
}