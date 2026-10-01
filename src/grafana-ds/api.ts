import { BackendSrvRequest, getBackendSrv } from '@grafana/runtime';
import cache from 'memory-cache';
import { Observable, lastValueFrom } from 'rxjs';
import type { Pair } from './types';

/**
 * API client for Grafana internal APIs.
 * Only supports GET requests to local Grafana instance.
 */
export default class Api {
  cache: any;
  lastCacheDuration: number | undefined;

  constructor() {
    this.cache = new cache.Cache();
  }

  /**
   * Performs a GET request to a Grafana internal API endpoint.
   * @param path - API path (e.g., '/api/datasources', '/api/dashboards')
   * @param params - Optional query parameters
   */
  async get(
    path: string,
    params?: Array<Pair<string, string>>
  ): Promise<any> {
    const response = this._request(path, params);
    return (await lastValueFrom(response)).data;
  }

  /**
   * Health check - verifies connectivity to Grafana API.
   */
  async test(): Promise<any> {
    return lastValueFrom(this._request('/api/health'));
  }

  /**
   * Returns a cached API response if it exists, otherwise queries the API.
   */
  async cachedGet(
    cacheDurationSeconds: number,
    path: string,
    params: Array<Pair<string, string>>
  ): Promise<any> {
    if (!cacheDurationSeconds) {
      return await this.get(path, params);
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
    if (cachedItem) {
      return Promise.resolve(cachedItem);
    }

    const result = await this.get(path, params);
    this.cache.put(cacheKey, result, cacheDurationSeconds * 1000);

    return result;
  }

  /**
   * Makes a GET request to Grafana internal API.
   */
  private _request(
    path: string,
    params?: Array<Pair<string, string>>
  ): Observable<any> {
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

    return getBackendSrv().fetch(req);
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
