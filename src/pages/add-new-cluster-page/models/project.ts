export interface TlsSettings {
    self_signed_certificate: boolean;
    ca_cert: string;
    client_authentication: boolean;
    server_name: string;
    client_certificate: string;
    client_key: string;
    skip_tls_verification: boolean;
}

export interface Project {
    host: string;
    authentication_enabled: boolean;
    username: string | null;
    password: string | null;
    status: string;
    tls?: TlsSettings;
}
