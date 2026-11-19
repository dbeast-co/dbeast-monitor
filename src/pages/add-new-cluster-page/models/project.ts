export interface Project {
    host: string;
    authentication_enabled: boolean;
    username: string | null;
    password: string | null;
    skip_tls?: boolean;
    ca_cert?: string;
    status: string;
}
