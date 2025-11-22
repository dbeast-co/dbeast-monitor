import { css } from '@emotion/css';
import { GrafanaTheme2 } from '@grafana/data';

export const getMonitoringClusterInjectionPanelStyles = (theme: GrafanaTheme2) => {
  return {
    monitoringClusterInjectionPanel: css`
      flex: 1 1 100%;
      
      .MuiFormControlLabel-root {
        margin: 0;
        width: 100%;
        
        .MuiFormControlLabel-label {
          color: ${theme.colors.text.primary};
          font-size: 14px;
        }
      }
      
      .MuiCheckbox-root {
        color: ${theme.colors.text.secondary};
        
        &.Mui-checked {
          color: ${theme.colors.primary.main};
        }
        
        &:hover {
          background-color: ${theme.colors.action.hover};
        }
      }
      
      .MuiSvgIcon-root {
        fill: ${theme.colors.text.secondary};
        
        &.Mui-checked {
          fill: ${theme.colors.primary.main};
        }
      }
      
      // Material-UI specific overrides for checkbox icons
      .css-vubbuv {
        fill: ${theme.colors.text.secondary} !important;
      }
    `,

    configItem: css`
      margin-bottom: 2px !important;

      .css-1n4u71h-Label{
         font-size: 16px !important;
      }
      
      &:last-child {
        margin-bottom: 0;
      }
    `
  };
};
