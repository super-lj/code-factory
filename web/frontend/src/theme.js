import { createMuiTheme } from '@material-ui/core/styles';
import teal from '@material-ui/core/colors/teal';
import amber from '@material-ui/core/colors/amber';

const appTheme = createMuiTheme({
  palette: {
    primary: teal,
    secondary: amber,
  },
});

export default appTheme;
