import React, {FunctionComponent, useEffect} from 'react';
import {Dialog, DialogActions, DialogContent, DialogTitle} from '@mui/material';
import TextField from '@mui/material/TextField';
import {Button, Spinner} from '@grafana/ui';
import "./password-dialog.scss";
import {Project} from '../../panels/dbeast-add_new_es_cluster-panel/models/project';


interface OwnProps {
    project: Project;
    handleUpgrade: (project: Project) => void;
    handleClose: () => void;
    handleSkip: () => void;
    isShowSpinner: boolean;
    isShowError: boolean;
}

type Props = OwnProps;

const PasswordDialogComponent: FunctionComponent<Props> = (props) => {

    const [password, setPassword] = React.useState("");

    const [username, setUsername] = React.useState("");

    useEffect(() => {
        console.log("PasswordDialogComponent props", props);
    }, [props]);

    const handleUpgrade = () => {

        const project = {
            ...props.project,
            password: password,
            username: username
        }

       props.handleUpgrade(project);
    }
    const handleClose = () => {
        props.handleClose();
    }
    const handleSkip = () => {
        props.handleSkip();
    }

  return (
      <Dialog open={true}>

        <DialogTitle>Upgrade project
            <Button
                onClick={handleClose}
                className="close-button"
                style={{ position: 'absolute', right: '16px', top: '16px' }}
            >
                X
            </Button>
        </DialogTitle>
        <DialogContent>
            {props.project.authentication_enabled && !props.isShowError && (
                <>
                    <label className="label">Host url:</label>
                    <span className="value">{props.project.host}</span>
                    <div className="form-control">
                        <label className="label">Username:</label>
                       <TextField
                            label="Username"
                            value={username}
                            fullWidth
                            margin="dense"
                            variant="standard"
                            onChange={(e) => setUsername(e.target.value)}
                        />
                    </div>
                    <div className="form-control">
                        <label className="label">Password:</label>
                        <TextField
                            label="Password"
                            type="password"
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                            fullWidth
                            margin="dense"
                            variant="standard"
                        />
                    </div>
                </>
            )}
            <div className="actions">
                <Button variant={'primary'} onClick={handleUpgrade}>Upgrade</Button>
                <Button variant={'primary'} onClick={handleSkip}>Skip</Button>
            </div>

        </DialogContent>

          {props.isShowError && (
                <div className="error-message">Cannot update cluster with url: {props.project.host}</div>
          )}

          {props.isShowSpinner && (
            <Spinner></Spinner>
          )}
        <DialogActions>

        </DialogActions>
      </Dialog>
  );
};

export default PasswordDialogComponent;
