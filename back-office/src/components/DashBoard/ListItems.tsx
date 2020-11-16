import * as React from "react";

import { Link,useRouteMatch } from "react-router-dom";

import ListItem from "@material-ui/core/ListItem";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import ListItemText from "@material-ui/core/ListItemText";
import ListSubheader from "@material-ui/core/ListSubheader";
import AccountBox from "@material-ui/icons/AccountBox";
import PeopleIcon from "@material-ui/icons/People";
import EventNote from "@material-ui/icons/EventNote";
import LayersIcon from "@material-ui/icons/Layers";
import AssignmentIcon from "@material-ui/icons/Assignment";

export const MainListItems = (props) => {
  let {url} = useRouteMatch();
  console.log("likn props ", history);
  return (
    <div>
      <ListItem button>
        <ListItemIcon>
          <AccountBox />
        </ListItemIcon>
        <Link to={`${url}/account`}>
        <ListItemText primary="Account" />
        </Link>
      </ListItem>
      <ListItem button>
        <ListItemIcon>
          <PeopleIcon />
        </ListItemIcon>
        <Link to={`${url}/members`}>
          <ListItemText primary="Members" />
        </Link>
      </ListItem>
      {/* <ListItem button>
        <ListItemIcon>
          <EventNote />
        </ListItemIcon>
        <ListItemText primary="Plans" />
      </ListItem>
      <ListItem button>
        <ListItemIcon>
          <LayersIcon />
        </ListItemIcon>
        <ListItemText primary="Programs" />
      </ListItem> */}
    </div>
  );
};
export const secondaryListItems = (
  <div>
    <ListSubheader inset>Saved reports</ListSubheader>
    <ListItem button>
      <ListItemIcon>
        <AssignmentIcon />
      </ListItemIcon>
      <ListItemText primary="Current month" />
    </ListItem>
    <ListItem button>
      <ListItemIcon>
        <AssignmentIcon />
      </ListItemIcon>
      <ListItemText primary="Last quarter" />
    </ListItem>
    <ListItem button>
      <ListItemIcon>
        <AssignmentIcon />
      </ListItemIcon>
      <ListItemText primary="Year-end sale" />
    </ListItem>
  </div>
);
