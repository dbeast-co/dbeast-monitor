/** @jest-environment jsdom */
import React from 'react';
import { fireEvent, render, screen, waitFor } from '@testing-library/react';
import { AppConfig, AppConfigProps } from './AppConfig';

const get = jest.fn();
const post = jest.fn();
jest.mock('@grafana/runtime', () => ({ getBackendSrv: () => ({ get, post }) }));
jest.mock('@grafana/ui', () => ({
  useStyles2: () => ({}),
  Alert: ({ title, children }: { title: string; children: React.ReactNode }) => <div role="alert">{title}{children}</div>,
  Legend: ({ children }: { children: React.ReactNode }) => <div>{children}</div>,
  Button: ({ children, onClick }: { children: React.ReactNode; onClick: () => void }) => <button onClick={onClick}>{children}</button>,
}));

const uid = 'dbeast-grafanaapi-datasource';
const definition = { uid, type: uid, name: 'Grafana internal API' };
const lookup = `/api/datasources/uid/${uid}`;
const resource = '/api/plugins/dbeast-dbeastmonitor-app/resources/grafana_api_datasource';
const props = {
  plugin: { meta: { id: 'dbeast-dbeastmonitor-app', enabled: true, jsonData: {} } },
  query: {},
} as AppConfigProps;

beforeEach(() => {
  get.mockReset();
  post.mockReset();
});

it('does not access resources while the plugin is disabled', () => {
  render(<AppConfig {...props} plugin={{ ...props.plugin, meta: { ...props.plugin.meta, enabled: false } } as typeof props.plugin} />);
  expect(get).not.toHaveBeenCalled();
});

it('does not change an existing datasource', async () => {
  get.mockResolvedValue(definition);
  render(<AppConfig {...props} />);
  await waitFor(() => expect(get).toHaveBeenCalledWith(lookup, undefined, undefined, { showErrorAlert: false }));
  expect(get).toHaveBeenCalledTimes(1);
  expect(post).not.toHaveBeenCalled();
});

it('fetches the backend definition and creates a missing datasource', async () => {
  get.mockRejectedValueOnce({ status: 404 }).mockResolvedValueOnce(definition);
  post.mockResolvedValue({});
  render(<AppConfig {...props} />);
  await waitFor(() => expect(post).toHaveBeenCalledWith('/api/datasources', definition, {
    showErrorAlert: false, showSuccessAlert: false,
  }));
  expect(get).toHaveBeenCalledWith(resource, undefined, undefined, { showErrorAlert: false });
});

it('reports lookup errors without creating a datasource and permits retry', async () => {
  get.mockRejectedValueOnce({ status: 403, message: 'Access denied' }).mockResolvedValueOnce(definition);
  render(<AppConfig {...props} />);
  expect(await screen.findByText('Access denied')).toBeInTheDocument();
  expect(post).not.toHaveBeenCalled();
  fireEvent.click(screen.getByRole('button', { name: 'Retry' }));
  await waitFor(() => expect(get).toHaveBeenCalledTimes(2));
  await waitFor(() => expect(screen.queryByRole('alert')).not.toBeInTheDocument());
});

it('reports an existing datasource with the wrong type', async () => {
  get.mockResolvedValue({ ...definition, type: 'other' });
  render(<AppConfig {...props} />);
  expect(await screen.findByText(/already used by another datasource type/)).toBeInTheDocument();
  expect(post).not.toHaveBeenCalled();
});

it('accepts a concurrent creation only after verifying the existing datasource', async () => {
  get.mockRejectedValueOnce({ status: 404 }).mockResolvedValueOnce(definition).mockResolvedValueOnce(definition);
  post.mockRejectedValue({ status: 409 });
  render(<AppConfig {...props} />);
  await waitFor(() => expect(get).toHaveBeenCalledTimes(3));
  expect(screen.queryByRole('alert')).not.toBeInTheDocument();
});

it('reports backend definition and creation errors', async () => {
  get.mockRejectedValueOnce({ status: 404 }).mockResolvedValueOnce(definition);
  post.mockRejectedValue({ status: 403, message: 'Cannot create datasource' });
  render(<AppConfig {...props} />);
  expect(await screen.findByText('Cannot create datasource')).toBeInTheDocument();
});

it('reports a backend definition request failure', async () => {
  get.mockRejectedValueOnce({ status: 404 }).mockRejectedValueOnce({ status: 500, message: 'Template unavailable' });
  render(<AppConfig {...props} />);
  expect(await screen.findByText('Template unavailable')).toBeInTheDocument();
  expect(post).not.toHaveBeenCalled();
});

it('reports a creation conflict when the expected UID is still missing', async () => {
  get.mockRejectedValueOnce({ status: 404 }).mockResolvedValueOnce(definition)
    .mockRejectedValueOnce({ status: 404, message: 'Datasource not found' });
  post.mockRejectedValue({ status: 409 });
  render(<AppConfig {...props} />);
  expect(await screen.findByText('Datasource not found')).toBeInTheDocument();
});

it('rejects an invalid backend definition', async () => {
  get.mockRejectedValueOnce({ status: 404 }).mockResolvedValueOnce({ ...definition, uid: 'wrong' });
  render(<AppConfig {...props} />);
  expect(await screen.findByText('The Grafana API datasource definition is invalid.')).toBeInTheDocument();
  expect(post).not.toHaveBeenCalled();
});
