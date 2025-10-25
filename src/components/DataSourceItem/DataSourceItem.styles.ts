import { css } from '@emotion/css';
import { GrafanaTheme2 } from '@grafana/data';

export const getDataSourceItemStyles = (theme: GrafanaTheme2) => {
  return {
    positionRelative: css`
      position: relative;
    `,

    formGroup: css`
      border: 1px solid ${theme.colors.border.strong};
      display: flex;
      flex-direction: column;
      max-width: 500px;
      padding: 20px;
      position: relative;
      min-height: 300px;
      min-width: 500px;
      border-radius: ${theme.shape.radius.default};
      background: ${theme.colors.background.primary};

      header {
        display: flex;
        justify-content: space-between;
        align-items: flex-start;
        width: 100%;
        margin-bottom: ${theme.spacing(2)};
      }

      footer {
        margin-top: ${theme.spacing(2)};
      }

      h3 {
        font-size: ${theme.typography.h3.fontSize};
        font-weight: ${theme.typography.h3.fontWeight};
        color: ${theme.colors.text.primary};
        margin: 0;
      }

      p {
        color: ${theme.colors.text.secondary};
        margin: ${theme.spacing(0.5, 0, 0, 0)};
        font-size: ${theme.typography.bodySmall.fontSize};
      }
    `,

    divider: css`
      height: 1px;
      background: ${theme.colors.border.medium};
      margin: ${theme.spacing(2, 0)};
    `,

    headerCluster: css`
      max-width: 80%;
      white-space: nowrap;
      overflow: auto;
    `,

    statusBadge: css`
      span {
        padding: ${theme.spacing(1, 2)};
        border-radius: ${theme.shape.radius.default};
        font-size: ${theme.typography.bodySmall.fontSize};
        font-weight: ${theme.typography.fontWeightMedium};
        display: inline-block;
        text-align: center;

        &.UNTESTED {
          background: ${theme.colors.secondary.main};
          color: ${theme.colors.secondary.contrastText};
        }

        &.YELLOW {
          background: ${theme.colors.warning.main};
          color: ${theme.colors.warning.contrastText};
        }

        &.RED,
        &.ERROR {
          background: ${theme.colors.error.main};
          color: ${theme.colors.error.contrastText};
        }

        &.GREEN {
          background: ${theme.colors.success.main};
          color: ${theme.colors.success.contrastText};
        }
      }
    `,

    statsGrid: css`
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(120px, 1fr));
      gap: ${theme.spacing(2)};
      width: 100%;
    `,

    statItem: css`
      display: flex;
      flex-direction: column;
      gap: ${theme.spacing(0.5)};
    `,

    statLabel: css`
      font-weight: ${theme.typography.fontWeightMedium};
      color: ${theme.colors.text.secondary};
      font-size: ${theme.typography.bodySmall.fontSize};
      padding: ${theme.spacing(0.5, 1)};
      border: 1px solid ${theme.colors.border.medium};
      border-radius: ${theme.shape.radius.default};
      background: transparent;
      display: inline-block;
    `,

    statValue: css`
      color: ${theme.colors.text.primary};
      font-size: ${theme.typography.body.fontSize};
      padding-left: ${theme.spacing(1)};
    `,

    selectWrapper: css`
      min-width: 250px;
    `,

    spinnerOverlay: css`
      position: fixed;
      inset: 0;
      background: ${theme.colors.background.canvas}88;
      z-index: ${theme.zIndex.modal};
      display: flex;
      justify-content: center;
      align-items: center;
      width: 100%;
      height: 100%;
    `,
  };
};
