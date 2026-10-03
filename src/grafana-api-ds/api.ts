import { BackendSrvRequest, FetchResponse, getBackendSrv } from '@grafana/runtime';
import cache from 'memory-cache';
import { Observable, lastValueFrom } from 'rxjs';
import type { Pair } from './types';

/**
 * API client for Grafana internal APIs.
 * Only supports GET requests to local Grafana instance.
 */
export default class Api {
  cache: cache.CacheClass<string, unknown>;
  private pendingRequests = new Map<string, Promise<unknown>>();
  lastCacheDuration: number | undefined;

  constructor() {
    this.cache = new cache.Cache();
  }

  /**
   * Performs a GET request to a Grafana internal API endpoint.
   * @param path - API path (e.g., '/api/datasources', '/api/dashboards')
   * @param params - Optional query parameters
   */
  async get<T = unknown>(
    path: string,
    params?: Array<Pair<string, string>>
  ): Promise<T> {
    const response = this._request(path, params);
    return (await lastValueFrom(response)).data as T;
  }

  /**
   * Health check - verifies connectivity to Grafana API.
   */
  async test(): Promise<FetchResponse<{ status: number; statusText?: string }>> {
    return lastValueFrom(this._request<{ status: number; statusText?: string }>('/api/health'));
  }

  /**
   * Returns a cached API response if it exists, otherwise queries the API.
   */
  async cachedGet<T>(
    cacheDurationSeconds: number,
    path: string,
    params: Array<Pair<string, string>>
  ): Promise<T> {
    if (!cacheDurationSeconds) {
      return this.get(path, params);
    }

    let cacheKey = path;

    if (params && params.length > 0) {
      cacheKey =
        cacheKey +
        (cacheKey.search(/\?/) >= 0 ? '&' : '?') +
        params.map(([k, v]) => `${encodeURIComponent(k)}=${encodeURIComponent(v)}`).join('&');
    }

    // If cache duration changed, clear all cached entries to prevent stale data
    if (this.lastCacheDuration !== cacheDurationSeconds) {
      this.cache = new cache.Cache();
    }
    this.lastCacheDuration = cacheDurationSeconds;

    const cachedItem = this.cache.get(cacheKey);
    if (cachedItem !== null) {
      return cachedItem as T;
    }

    const pendingRequest = this.pendingRequests.get(cacheKey);
    if (pendingRequest) {
      return pendingRequest as Promise<T>;
    }

    const request = this.get<T>(path, params);
    this.pendingRequests.set(cacheKey, request);

    try {
      const result = await request;
      this.cache.put(cacheKey, result, cacheDurationSeconds * 1000);
      return result;
    } finally {
      this.pendingRequests.delete(cacheKey);
    }
  }

  /**
   * Makes a GET request to Grafana internal API.
   */
  private _request<T = unknown>(
    path: string,
    params?: Array<Pair<string, string>>
  ): Observable<FetchResponse<T>> {
    let url = path;

    // Deduplicate forward slashes
    url = url.replace(/[\/]+/g, '/');

    // Add query parameters
    if (params && params.length > 0) {
      url =
        url +
        (url.search(/\?/) >= 0 ? '&' : '?') +
        params
          .filter(([key]) => key)
          .map(([k, v]) => `${encodeURIComponent(k)}=${encodeURIComponent(v)}`)
          .join('&');
    }

    // Validate URL safety
    if (!isSafeURL(url)) {
      throw new Error('URL path contains unsafe characters');
    }

    const req: BackendSrvRequest = {
      url,
      method: 'GET',
    };

    return getBackendSrv().fetch<T>(req);
  }
}

function isSafeURL(origUrl: string) {
  // browsers interpret backslash as slash
  const url = decodeURIComponent(origUrl.replace(/\\/g, '/'));
  if (url.endsWith('/..')) {
    return false;
  }

  if (url.includes('/../')) {
    return false;
  }

  if (url.includes('/..?')) {
    return false;
  }

  if (url.includes('\t')) {
    return false;
  }

  return true;
}
