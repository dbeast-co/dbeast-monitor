import { css } from '@emotion/css';
import { GrafanaTheme2 } from '@grafana/data';

export const getDataSourceItemStyles = (theme: GrafanaTheme2) => {
  return {
    positionRelative: css`
      position: relative;
    `,

    formGroup: css`
      border: 1px solid ${theme.colors.border.weak};
      display: flex;
      flex-direction: column;
      max-width: 500px;
      padding: 10px 20px;
      position: relative;
      min-height: 300px;
      min-width: 500px;
      border-radius: 10px;
      box-shadow: 5px 5px 10px ${theme.colors.action.hover};
      background: ${theme.colors.background.primary};

      header {
        display: grid;
        width: 100%;
        grid-template-columns: 1fr 0.2fr;
      }

      footer {
        margin-block-start: 1rem;
      }

      .actions {
        display: unset !important;
        margin-top: 0 !important;

        .col {
          flex: unset !important;

          &.header-cluster {
            max-width: 95%;
            width: 100%;
            overflow: auto;
          }
        }
      }

      h3 {
        font-size: 1.5rem;
        font-weight: 700;
        color: ${theme.colors.text.primary};
      }

      p {
        color: ${theme.colors.text.secondary};
      }
    `,

    divider: css`
      .MuiDivider-light,
      .MuiDivider-root {
        border-color: ${theme.colors.border.medium} !important;
      }
    `,

    actions: css`
      button,
      span {
        min-width: 75px;
        padding: 8px 16px;
        border-radius: 8px;
        border: none;
        box-shadow: none;
        color: ${theme.colors.primary.contrastText};
        background: ${theme.colors.primary.main};
      }

      button[disabled] {
        background: ${theme.colors.action.disabledBackground};
        color: ${theme.colors.action.disabledText};
        cursor: not-allowed;
      }

      span {
        display: flex;
        align-items: center;
        justify-content: center;
        color: ${theme.colors.text.primary};
        border: 1px solid ${theme.colors.border.medium};

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

    listItem: css`
    
      .MuiListItem-root {
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(100px, 1fr));
        grid-row-gap: 20px;
        grid-column-gap: 9px;

        div {
          .label {
            font-weight: 600;
            color: ${theme.colors.text.primary};
            border: 1px solid ${theme.colors.border.medium};
            border-radius: 6px;
            background: transparent;
            font-size: 14px;
            padding: 2px 6px;
            display: inline-block;
            margin-bottom: 4px;
          }

          .MuiListItemText-primary {
            color: ${theme.colors.text.primary};
          }
        }
      }
    `,

    buttons: css`
    
      display: flex;
      justify-content: flex-end;
      gap: 10px;
    
      button.btn {
        border: 1px solid ${theme.colors.border.medium};
        border-radius: 5px;
        background: ${theme.colors.primary.main};
        color: ${theme.colors.primary.contrastText};
        display: flex;
        justify-content: center;
        align-items: center;
        text-decoration: none;
        height: revert;
        max-height: 40px;
        min-width: 70px;
        min-height: 40px;

        &:hover {
          background: ${theme.colors.primary.shade};
        }
      }
    `,

    stack: css`
      .MuiStack-root {
        display: flex;
        justify-content: space-between;
      }
    `,

    select: css`
      .MuiInputBase-root {
        max-width: 100%;
      }
      :global(.MuiList-root) {
        background-color: ${theme.colors.background.secondary} !important;
      }

      :global(.MuiMenu-list) {
        background: red !important;
        // Add other overrides here as needed
      }

      .MuiSelect-select {
        padding: 10px 14px !important;
        color: ${theme.colors.text.primary} !important;
      }

      #demo-simple-select-label,
      .MuiSelect-icon,
      #demo-simple-select {
        color: ${theme.colors.primary.main} !important;
      }

      .MuiInputBase-formControl {
        margin-block-start: 21px;
      }

      .MuiInputLabel-formControl {
        transform: translate(10px, 30px) !important;
      }

      .MuiInputLabel-shrink {
        transform: translate(13px, 13px) scale(0.75) !important;
      }

      .monitor-type {
        font-size: 14px;
        color: ${theme.colors.primary.main};
      }

      .MuiFormControl-fullWidth {
        position: relative;
        top: -21px;
      }

      .MuiOutlinedInput-notchedOutline {
        border-color: ${theme.colors.primary.main};
      }

    `,

    dialog: css`
      .MuiPaper-root {
        background-color: ${theme.colors.background.primary} !important;
        border: 1.5px solid ${theme.colors.border.medium};
        border-radius: 7px;

        .MuiList-root {
          color: ${theme.colors.text.primary};
          background-color: ${theme.colors.background.secondary};
        }

        .MuiDialogTitle-root {
          color: ${theme.colors.text.primary} !important;
        }

        .MuiDialogActions-root {
          button {
            background-color: ${theme.colors.primary.main};
            color: ${theme.colors.primary.contrastText} !important;

            &:hover {
              background-color: ${theme.colors.primary.shade};
            }
          }

          button.btn-error {
            background-color: ${theme.colors.error.main} !important;
            color: ${theme.colors.error.contrastText} !important;

            &:hover {
              background-color: ${theme.colors.error.shade} !important;
            }
          }
        }
      }
    `,

    spinnerOverlay: css`
      position: fixed;
      inset: 0;
      background: ${theme.colors.background.canvas}88;
      margin: 0 auto;
      z-index: 1000;
      display: flex;
      justify-content: center;
      align-items: center;
      width: 100%;
      height: 100%;

      .MuiCircularProgress-root {
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);

        .MuiCircularProgress-svg {
          .css-14891ef {
            stroke: ${theme.colors.primary.main} !important;
          }
        }
      }
    `,

    headerCluster: css`
      max-width: 80%;
      white-space: nowrap;
      overflow: auto;
    `,

    menuItem: css`
      .MuiMenuItem-root {
        color: ${theme.colors.text.primary} !important;
        background-color: ${theme.colors.background.secondary} !important;
       
        // &:hover {
        //   background-color: ${theme.colors.action.hover} !important;
        //   color: ${theme.colors.text.primary} !important;
        // }

        a {
          color: inherit;
          text-decoration: none;
        }
      }
    `
  };
};
