import { css } from '@emotion/css';
import { GrafanaTheme2 } from '@grafana/data';

export const getStyles = (theme: GrafanaTheme2) => {
  return {
    connectionsAndConfig: css`
      display: flex;
      flex-direction: column;
      justify-content: center;
      gap: 0;
      max-width: 1200px;
      margin: 0 auto;
    `,

    header: css`
      display: flex;
      justify-content: center;
      width: 100%;
      margin-block-end: 2rem;

      h1 {
        color: ${theme.colors.text.primary};
        font-size: ${theme.typography.h1.fontSize};
        font-weight: ${theme.typography.h1.fontWeight};
        margin: 0;
      }
    `,

    title: css`
      font-size: 18px;  // Back to original SCSS size
      text-align: center;
      margin-block: 0.4rem;
      color: ${theme.colors.text.primary};
    `,

    sourcePanel: css`
      border: 1px solid ${theme.colors.border.weak};
      display: flex;
      flex-direction: column;
      max-width: 350px !important;  // Back to original SCSS with !important
      padding: 10px 20px;
      position: relative;
      min-height: 300px;  // Back to original SCSS height
      justify-content: space-between;
      border-radius: 4px 0 0 4px;
      background: ${theme.colors.background.primary};
      border-right: 1px solid ${theme.colors.border.medium};

      section {
        min-height: calc(100% - 62px);
        display: grid;
      }

      label {
        color: ${theme.colors.text.primary};
        font-size: 16px;  // Added larger font size
      }

      input {
        color: ${theme.colors.text.primary};
        border-bottom: 2px solid ${theme.colors.border.medium} !important;
        font-size: 16px;  // Added larger font size
      }

      // Material-UI form control width
      .MuiFormControl-root {
        width: 100%;
      }

      .css-1vv4lmi::before {
        border-bottom: 2px solid ${theme.colors.border.medium} !important;
      }

      .css-1vv4lmi::after {
        border-bottom: 2px solid ${theme.colors.text.primary} !important;
      }

      .css-vubbuv {
        fill: ${theme.colors.text.secondary};
      }

      .css-mqt2s5.Mui-disabled {
        color: ${theme.colors.text.primary};
        opacity: 0.3;
      }

      .css-mnn31.Mui-disabled {
        background: transparent;
      }

      .auth_wrapper {
        width: 100%;
        display: flex;
        justify-content: space-between;
        column-gap: 10px;
        margin-block-end: 7px;
      }
    `,

    hostWrapper: css`
      margin-block-end: 1.5rem;
      position: relative;

      .invalid {
        position: absolute;
        color: ${theme.colors.error.text};
        top: 100%;
        display: flex;
      }
    `,

    status: css`
      color: ${theme.colors.text.primary};
      position: absolute;
      top: -18%;
      left: 80%;  // Back to original SCSS position
      display: inline-flex;

      .UNTESTED {
        background: ${theme.colors.secondary.main};
        padding: 2px 10px;
        font-size: 10px;
        border-radius: 7px;
        color: ${theme.colors.secondary.contrastText};
      }

      .YELLOW {
        background: ${theme.colors.warning.main};
        padding: 2px 10px;
        font-size: 10px;
        border-radius: 7px;
        color: ${theme.colors.warning.contrastText};
      }

      .ERROR, .RED {
        background: ${theme.colors.error.main};
        padding: 2px 10px;
        font-size: 10px;
        border-radius: 7px;
        color: ${theme.colors.error.contrastText};
      }

      .GREEN {
        background: ${theme.colors.success.main};
        padding: 2px 10px;
        font-size: 10px;
        border-radius: 7px;
        color: ${theme.colors.success.contrastText};
      }
    `,

    config: css`
      padding: 10px 20px;
      border-top-right-radius: 4px;
      border-bottom-right-radius: 4px;
      border-left: 0 !important;
      max-width: 350px;  // Back to original SCSS
      width: 100%;
      border: 1px solid ${theme.colors.border.weak};
      background: ${theme.colors.background.primary};
      min-height: 475px;  // Keep this for content
      font-size: 16px;

      .wrapper {
        display: grid;
        grid-template-columns: 1fr;
        align-items: center;

        .hide {
          display: none;
        }
      }

      .MuiDivider-root {
        margin: 7px 0 !important;  // Back to original SCSS
        flex-shrink: 0 !important;
        border-width: 1px 3px thin !important;
        border-style: solid !important;
        border-color: ${theme.colors.border.weak} !important;
        opacity: 0.2 !important;
      }

      .MuiFormControlLabel-label {
        font-size: 16px !important;
      }

      &.not-rounded {
        border-top-right-radius: 0;
        border-bottom-right-radius: 0;
        border-right: 1px solid ${theme.colors.border.medium};
      }
    `,

    actions: css`
      display: flex;
      column-gap: 10px;
      row-gap: 10px;
      flex-wrap: wrap;
      margin-block-start: 10px;
      align-self: flex-end;
      justify-content: flex-end;

      button,
      span {
        min-width: 75px;
        padding: 8px 16px;
        border-radius: 8px;
        border: none;
        box-shadow: none;
        color: ${theme.colors.primary.contrastText};
        background: ${theme.colors.primary.main};
        cursor: pointer;
        font-size: 16px;

        &:hover {
          background: ${theme.colors.primary.shade};
        }

        &:disabled {
          background: ${theme.colors.action.disabledBackground};
          color: ${theme.colors.action.disabledText};
          cursor: not-allowed;
        }
      }

      .deploy-btn {
        background: ${theme.colors.success.main};
        color: ${theme.colors.success.contrastText};
        position: relative;
        bottom: 1rem;

        &:hover {
          background: ${theme.colors.success.shade};
        }
      }
    `,

    spinnerOverlay: css`
      position: fixed;
      inset: 0;
      background: transparent;
      margin: 0 auto;
      z-index: 10000;

      .MuiCircularProgress-root {
        position: absolute;
        top: 35%;
        left: 45%;
        transform: translate(-35%, -45%);

        .MuiCircularProgress-svg {
          .css-14891ef {
            stroke: ${theme.colors.primary.main} !important;
          }
        }
      }
    `,

    logstashList: css`
      padding: 10px 20px;
      border-top-right-radius: 4px;
      border-bottom-right-radius: 4px;
      border-left: 0;
      max-width: 350px;
      width: 100%;
      border: 1px solid ${theme.colors.border.weak};
      background: ${theme.colors.background.primary};
      min-height: 475px;

      .cards-wrapper {
        overflow: auto;
        max-height: 500px;
        display: grid;
        grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
        gap: 10px;
      }

      .logstash-card {
        border: 1px solid ${theme.colors.border.weak};
        border-radius: 5px;
        padding: 7px;
        cursor: pointer;
        background: ${theme.colors.background.secondary};

        &:hover {
          border-color: ${theme.colors.border.strong};
          background: ${theme.colors.background.canvas};
        }

        .form-group {
          margin-block-end: 0.75rem;
          border-bottom: 1px solid ${theme.colors.border.weak};
          padding-block: 0.15rem;

          .logstash-label {
            opacity: 0.6;
            color: ${theme.colors.text.secondary};
          }

          .value {
            font-weight: 600;
            color: ${theme.colors.text.primary};

            &:last-child {
              border-bottom: none;
            }
          }
        }
      }
    `,

    divider: css`
      margin: 10px 0 !important;  // Back to original SCSS
      flex-shrink: 0 !important;
      border-width: 2px 3px thin !important;
      border-style: solid !important;
      border-color: ${theme.colors.border.weak} !important;
      opacity: 1 !important;  // Back to original SCSS (no opacity)

      &.add-new-cluster-divider {
        margin-block: 5.3rem !important;
        border-width: 1px !important;
        opacity: 1 !important;
        position: relative;
        //bottom: 2rem;
      }

      &.add-new-cluster-divider-2 {
        border-width: 1px !important;
        opacity: 0.2 !important;
      }
    `,

    sourceConnectionGroup: css`
      margin-block-end: 4.5rem;
    `,

    monitoringClusterGroup: css`
      margin-block-end: 2.1rem;
    `,

    monitoringClusterTitle: css`
      position: relative;
      bottom: 20px;
    `,

    formWrapper: css`
      margin-block-start: 3rem;
    `,

    sectionHeader: css`
      display: flex;
      align-items: center;
      gap: 8px;
      justify-content: center;  // Back to original SCSS center alignment
      margin-bottom: 1rem;

      .title {
        margin: 0;
        font-size: 18px;
        font-weight: 600;
        color: ${theme.colors.text.primary};
      }
    `,

    uploadFileContainer: css`
      display: flex;
      align-items: center;
      justify-content: space-evenly;
    `,

    sslWrapper: css`
      display: flex;
      align-items: center;
      justify-content: space-between;

      .upload_file {
        input {
          display: none;
        }

        span {
          min-width: 70px;
          padding: 5px 10px;
          display: inline-block;
          background: ${theme.colors.primary.main};
          color: ${theme.colors.primary.contrastText};
          cursor: pointer;
          border-radius: 8px;
        }

        input[disabled] + span {
          background: ${theme.colors.action.disabledBackground};
          color: ${theme.colors.action.disabledText};
          cursor: not-allowed;
        }
      }
    `,

    toastContainer: css`
      .Toastify__toast-icon {
        width: 11% !important;

        svg {
          width: 30px;
          height: 30px;
        }
      }

      .Toastify__toast {
        width: 55% !important;
        margin: 0 auto !important;
        bottom: -30px !important;
      }

      .Toastify__toast-container--bottom-right {
        bottom: 6em !important;
        right: 1em !important;
      }
    `,

    dialog: css`
      #alert-dialog-title {
        color: ${theme.colors.text.primary};
      }

      .MuiDialog-paper {
        border: 1px solid ${theme.colors.border.weak} !important;
        display: flex !important;
        flex-direction: column !important;
        max-width: 600px !important;
        padding: 10px 20px !important;
        position: relative !important;
        min-height: 300px !important;
        justify-content: space-between !important;
        border-radius: 4px 0 0 4px !important;
        background: ${theme.colors.background.primary} !important;
        width: 30% !important;
      }

      svg.MuiSvgIcon-root[data-testid='CheckBoxOutlineBlankIcon'],
      svg.MuiSvgIcon-root[data-testid='CheckBoxIcon'] {
        fill: ${theme.colors.text.primary} !important;
      }

      .header-actions {
        padding-inline-start: 1.7rem;
      }

      .MuiDialogActions-root {
        padding-inline-end: 1.7rem !important;
      }
    `,

    saveBtn: css`
      min-width: 75px;
      padding: 8px 16px;
      border-radius: 8px;
      border: none;
      box-shadow: none;
      color: ${theme.colors.primary.contrastText} !important;
      background: ${theme.colors.primary.main} !important;
      max-width: fit-content !important;
      cursor: pointer;

      &:hover {
        background: ${theme.colors.primary.shade} !important;
      }
    `,

    panelContainer: css`
      background-color: transparent !important;
    `,

    spinner: css`
      [data-testid='Spinner'] {
        background: transparent !important;
        inset: 0;
        z-index: 10000;
        display: flex;
        justify-content: center;
        align-items: center;
        position: fixed;

        .fa-spinner {
          font-size: 4rem;
        }
      }
    `
  };
};
