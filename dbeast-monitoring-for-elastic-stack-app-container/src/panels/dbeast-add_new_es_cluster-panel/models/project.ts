export interface Project {
    host: string;
    authentication_enabled: boolean;
    username: string | null;
    password: string | null;
    status: string;
}
