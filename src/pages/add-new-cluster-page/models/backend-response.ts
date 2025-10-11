export interface BackendResponse {
    prod: {
        elasticsearch: {
            status: string;
            error: string;
        };
        kibana: {
            status: string;
            error: string;
        };
    };
    mon: {
        elasticsearch: {
            status: string;
            error: string;
        };

    };
}
