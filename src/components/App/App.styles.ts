import { css } from '@emotion/css';
import { GrafanaTheme2 } from '@grafana/data';

export const getAppStyles = (theme: GrafanaTheme2) => {
  return {
    container: css`
      padding: 1rem !important;
      max-width: 100% !important;
    `
  };
};
