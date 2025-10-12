import { css } from '@emotion/css';
import { GrafanaTheme2 } from '@grafana/data';

export const getLogstashStyles = (theme: GrafanaTheme2) => {
  return {
    sourcePanel: css`
      display: flex;
      flex-direction: column;
      row-gap: 1rem;
      padding: 1rem;
      min-width: 400px;
      min-height: 300px;
    `,

    dialogContent: css`
      .MuiDialogContent-root {
        padding: 20px !important;
        background: ${theme.colors.background.primary} !important;
        min-width: 400px;
        min-height: 300px;
      }
    `,

    textField: css`
      .MuiTextField-root {
        width: 100%;
        margin-bottom: 1rem;

        input {
          color: ${theme.colors.text.primary} !important;
          border-bottom: 2px solid ${theme.colors.border.medium} !important;
        }
      }

      .MuiFormLabel-root {
        color: ${theme.colors.text.primary} !important;
      }

      .css-1vv4lmi::before {
        border-bottom: 2px solid ${theme.colors.border.medium} !important;
      }

      .css-1vv4lmi::after {
        border-bottom: 2px solid ${theme.colors.text.primary} !important;
      }
    `,

    actions: css`
      display: flex;
      column-gap: 0.75rem;
      margin-top: 1.5rem;
      justify-content: flex-end; // This ensures buttons are aligned to the right

      button {
        min-width: 75px;
        padding: 8px 16px;
        border-radius: 8px;
        border: none;
        box-shadow: none;
        color: ${theme.colors.primary.contrastText};
        background: ${theme.colors.primary.main};
        cursor: pointer;

        &:hover {
          background: ${theme.colors.primary.shade};
        }

        &:disabled {
          background: ${theme.colors.action.disabledBackground};
          color: ${theme.colors.action.disabledText};
          cursor: not-allowed;
          opacity: 0.3;
        }

        &.cancel-btn {
          background: ${theme.colors.secondary.main};
          color: ${theme.colors.secondary.contrastText};

          &:hover {
            background: ${theme.colors.secondary.shade};
          }
        }
      }
    `,

    dialog: css`
      .MuiDialog-paper {
        background: ${theme.colors.background.primary} !important;
        border: 1px solid ${theme.colors.border.weak} !important;
        border-radius: 4px !important;
        min-width: 500px !important;
        min-height: 400px !important;
        max-width: 600px !important;
        padding: 20px !important;
      }

      .MuiDialogTitle-root {
        color: ${theme.colors.text.primary} !important;
        background: ${theme.colors.background.primary} !important;
        padding: 16px 20px !important;
      }

      .MuiDialogContent-root {
        background: ${theme.colors.background.primary} !important;
        padding: 20px !important;
        color: ${theme.colors.text.primary} !important;
      }

      .MuiDialogActions-root {
        background: ${theme.colors.background.primary} !important;
        padding: 16px 20px !important;
      }

      // Material-UI specific overrides for proper theming
      .MuiPaper-root {
        background-color: ${theme.colors.background.primary} !important;
      }
    `
  };
};
