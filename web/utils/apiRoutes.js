export default function useApiRoutes() {
    const config = useRuntimeConfig();
    const apiBase = config.public.apiBase;

    return {
        LOGIN: `${apiBase}/user/login`,
        PROFILE: `${apiBase}/api/profile`,
        SIGNUP: `${apiBase}/basic/new`,
        CAPTCHA: `${apiBase}/captcha/`,
    }
}